package zkclient

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/eth"
	"io"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type Client struct {
	BaseURL          string
	chainID          uint64
	PaymasterFeeZero bool
	PaymasterAddress string
	PaymasterToken   string
}

var maxGasTipCap = big.NewInt(20000)
var maxGasFeeCap = big.NewInt(20000)

var ethClientMap map[string]*ethclient.Client // key by url
var zkClientMap map[string]*clients.Client    // key by url
var zkClientLock = &sync.Mutex{}
var sendTransactionLock = &sync.Mutex{}

func init() {
	zkClientMap = make(map[string]*clients.Client)
	ethClientMap = make(map[string]*ethclient.Client)
}

func NewZkClient(rpc string, paymentFeeZero bool, paymentMasterAddress string, paymentMasterToken string) *Client {
	aiZkClient := &Client{
		BaseURL:          rpc,
		PaymasterFeeZero: paymentFeeZero,
		PaymasterAddress: paymentMasterAddress,
		PaymasterToken:   paymentMasterToken,
	}

	if zkClientMap[rpc] == nil {
		zkClientLock.Lock()
		defer zkClientLock.Unlock()
		if zkClientMap[rpc] == nil {
			zkClient, err := aiZkClient.makeZkClient(rpc)
			if err == nil {
				zkClientMap[rpc] = &zkClient
			}
			ethClient, err := aiZkClient.makeETHClient(rpc)
			if err == nil {
				ethClientMap[rpc] = ethClient
			}
		}
	}

	return aiZkClient
}

func (c *Client) makeZkClient(url string) (clients.Client, error) {
	zkClient, err := clients.Dial(url)
	if err != nil {
		return nil, err
	}

	return zkClient, nil
}

func (c *Client) makeETHClient(url string) (*ethclient.Client, error) {
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return ethClient, nil
}

func (c *Client) GetZkClient() (clients.Client, error) {
	if zkClientMap[c.BaseURL] == nil {
		zkClientLock.Lock()
		defer zkClientLock.Unlock()

		if zkClientMap[c.BaseURL] != nil {
			return *zkClientMap[c.BaseURL], nil
		}

		zkClient, err := c.makeZkClient(c.BaseURL)
		if err != nil {
			return nil, err
		}

		zkClientMap[c.BaseURL] = &zkClient
		return *zkClientMap[c.BaseURL], nil
	}

	return *zkClientMap[c.BaseURL], nil
}

func (c *Client) GetETHClient() (*ethclient.Client, error) {
	if ethClientMap[c.BaseURL] == nil {
		zkClientLock.Lock()
		defer zkClientLock.Unlock()

		if ethClientMap[c.BaseURL] != nil {
			return ethClientMap[c.BaseURL], nil
		}

		ethClient, err := c.makeETHClient(c.BaseURL)
		if err != nil {
			return nil, err
		}

		ethClientMap[c.BaseURL] = ethClient
		return ethClientMap[c.BaseURL], nil
	}

	return ethClientMap[c.BaseURL], nil
}

func (c *Client) GetChainID() (uint64, error) {
	if c.chainID > 0 {
		return c.chainID, nil
	}
	client, err := c.GetZkClient()
	if err != nil {
		return 0, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	c.chainID = chainID.Uint64()
	return c.chainID, nil
}

func (c *Client) GetGasPrice() (*big.Int, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func (c *Client) PopulateTransaction(ctx context.Context, address common.Address, tx accounts.Transaction) (*zktypes.Transaction712, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return nil, err
	}
	if tx.ChainID == nil {
		tx.ChainID = big.NewInt(int64(chainID))
	}
	tx.Nonce = new(big.Int).SetUint64(nonce)
	tx.GasFeeCap = gasPrice
	if tx.GasTipCap == nil {
		tx.GasTipCap = big.NewInt(0)
	}
	if tx.Meta == nil {
		tx.Meta = &zktypes.Eip712Meta{GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64())}
	} else if tx.Meta.GasPerPubdata == nil {
		tx.Meta.GasPerPubdata = utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64())
	}
	if tx.Gas == 0 {
		gas, err := client.EstimateGasL2(context.Background(), tx.ToCallMsg(address))
		if err != nil {
			return nil, fmt.Errorf("failed to EstimateGasL2: %w", err)
		}
		tx.Gas = gas
	}
	if tx.Data == nil {
		tx.Data = hexutil.Bytes{}
	}
	if tx.Meta.PaymasterParams != nil {
		paymasterParams, err := c.GetPaymasterParamsWithFee(big.NewInt(0).Mul(tx.GasFeeCap, big.NewInt(int64(tx.Gas))))
		if err != nil {
			panic(err)
		}
		tx.Meta.PaymasterParams = paymasterParams
	}
	return tx.ToTransaction712(address), nil
}

func (c *Client) SignTransaction(signer accounts.Signer, tx *zktypes.Transaction712) ([]byte, error) {
	var gas uint64 = 0
	if tx.Gas != nil {
		gas = tx.Gas.Uint64()
	}
	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		signer.Address(),
		accounts.Transaction{
			To:         tx.To,
			Data:       tx.Data,
			Value:      tx.Value,
			Nonce:      tx.Nonce,
			GasTipCap:  tx.GasTipCap,
			GasFeeCap:  tx.GasFeeCap,
			Gas:        gas,
			AccessList: tx.AccessList,
			ChainID:    tx.ChainID,
			Meta:       tx.Meta,
		},
	)
	if err != nil {
		return nil, err
	}
	signature, err := signer.SignTypedData(signer.Domain(), preparedTx)
	if err != nil {
		return nil, err
	}
	return preparedTx.RLPValues(signature)
}

func (c *Client) GetPaymasterParams() (*zktypes.PaymasterParams, error) {
	if c.PaymasterAddress == "" || c.PaymasterToken == "" {
		return nil, nil
	}
	fee := big.NewInt(1000000000000)
	if c.PaymasterFeeZero {
		fee = big.NewInt(0)
	}
	paymasterParams, err := utils.GetPaymasterParams(
		common.HexToAddress(c.PaymasterAddress),
		&zktypes.ApprovalBasedPaymasterInput{
			Token:            common.HexToAddress(c.PaymasterToken),
			MinimalAllowance: fee,
			InnerInput:       []byte{},
		})
	if err != nil {
		return nil, err
	}
	return paymasterParams, nil
}

func (c *Client) GetPaymasterParamsWithFee(fee *big.Int) (*zktypes.PaymasterParams, error) {
	if c.PaymasterAddress == "" || c.PaymasterToken == "" {
		return nil, nil
	}
	if c.PaymasterFeeZero {
		fee = big.NewInt(0)
	}
	paymasterParams, err := utils.GetPaymasterParams(
		common.HexToAddress(c.PaymasterAddress),
		&zktypes.ApprovalBasedPaymasterInput{
			Token:            common.HexToAddress(c.PaymasterToken),
			MinimalAllowance: fee,
			InnerInput:       []byte{},
		})
	if err != nil {
		return nil, err
	}
	return paymasterParams, nil
}

func (c *Client) SetPaymasterParams(tx *accounts.Transaction) error {
	paymasterParams, err := c.GetPaymasterParams()
	if err != nil {
		return err
	}
	tx.Meta.PaymasterParams = paymasterParams
	return nil
}

func (c *Client) PaymasterParams() *zktypes.PaymasterParams {
	paymasterParams, err := c.GetPaymasterParams()
	if err != nil {
		panic(err)
	}
	return paymasterParams
}

func (c *Client) Transact(prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (*zktypes.Receipt, error) {
	var err error
	maxRetry := 5
	chainId, err := c.GetChainID()
	if err != nil {
		return nil, err
	}
	if chainId == configs.SubtensorEVMChainIDInt {
		maxRetry = 1
	}
	for i := 0; i < maxRetry; i++ {
		paymaster, err1 := c.GetPaymasterParams()
		if err1 != nil {
			return nil, err1
		}
		chainId, _ := c.GetChainID()
		if paymaster != nil || chainId == configs.ZkSyncChainIDInt {
			transact, err1 := c.createZKTransact(from, to, value, input)
			if err1 != nil {
				time.Sleep(1 * time.Second)
				err = err1
				continue
			}
			// get nonce again
			tx, err1 := c.signAndSendTx(prkHex, from, transact)
			if err1 != nil {
				err = err1
				logger.Info("retry signAndSendTx ", "", zap.Any("from", from),
					zap.Any("to", to), zap.Any("i", i), zap.Any("rpc", c.BaseURL))
				continue
			}
			if tx.Receipt.Status == types.ReceiptStatusFailed {
				return tx, fmt.Errorf("transact failed with status %d", tx.Receipt.Status)
			}
			return tx, nil
		} else {
			txReceipt, err1 := c.ExecuteETHTransact(prkHex, from, to, value, input)
			if err1 != nil {
				time.Sleep(1 * time.Second)
				err = err1
				logger.Info("retry signAndSendTx ", "", zap.Any("from", from),
					zap.Any("to", to), zap.Any("i", i), zap.Any("rpc", c.BaseURL), zap.Any("err", err))
				continue
			}
			zkReceipt := zktypes.Receipt{
				Receipt: *txReceipt,
			}
			if txReceipt.Status == types.ReceiptStatusFailed {
				return &zkReceipt, fmt.Errorf("transact failed with status %d", txReceipt.Status)
			}
			return &zkReceipt, nil
		}
	}
	return nil, fmt.Errorf("failed to transact from :%v ,to: %v after 5 retries at rpc :%v , err:%v", from, to, c.BaseURL, err)
}

func (c *Client) createZKTransact(from common.Address, to common.Address, value *big.Int, input []byte) (*accounts.Transaction, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}

	gasNumber, err := client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From:  from,
				To:    &to,
				Value: value,
				Data:  input,
			},
			Meta: &zktypes.Eip712Meta{
				GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
				PaymasterParams: c.PaymasterParams(),
			},
		},
	)

	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, err
	}

	return &accounts.Transaction{
		GasFeeCap: gasPrice,
		GasTipCap: gasPrice,
		Gas:       gasNumber,
		To:        &to,
		Value:     value,
		Data:      input,
		Meta: &zktypes.Eip712Meta{
			GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			PaymasterParams: c.PaymasterParams(),
		},
	}, err
}

/*
Copy from the function `(c *BoundContract) createDynamicTx` in the `eth` package.
*/
func (c *Client) SendETHTransact(client *ethclient.Client, prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (common.Hash, error) {
	fromAcc, _, err := eth.GetAccountInfo(prkHex)
	if err != nil {
		return common.Hash{}, errors.Join(err, errors.New("error while getting account info"))
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("NetworkID :%v", err.Error())
	}
	tx, err := eth.CreateEthTransaction(client, from, to, value, input)
	if err != nil {
		return common.Hash{}, fmt.Errorf("CreateEthTransaction :%v", err.Error())
	}
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), fromAcc)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SignTx:%v , err:%v", c.BaseURL, err.Error())
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SendTransaction:%v , err:%v", c.BaseURL, err.Error())
	}
	return signedTx.Hash(), err
}

func (c *Client) SendETHTransactWithLock(client *ethclient.Client, prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (common.Hash, error) {
	logger.Info("start SendETHTransactWithLock", "")
	defer logger.Info("end SendETHTransactWithLock", "")
	sendTransactionLock.Lock()
	defer sendTransactionLock.Unlock()
	fromAcc, _, err := eth.GetAccountInfo(prkHex)
	if err != nil {
		return common.Hash{}, errors.Join(err, errors.New("error while getting account info"))
	}
	logger.Info("GetAccountInfo SendETHTransactWithLock", "")
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("NetworkID :%v", err.Error())
	}
	logger.Info("NetworkID SendETHTransactWithLock", "")
	tx, err := eth.CreateEthTransaction(client, from, to, value, input)
	if err != nil {
		return common.Hash{}, fmt.Errorf("CreateEthTransaction :%v", err.Error())
	}
	logger.Info("CreateEthTransaction SendETHTransactWithLock", "")
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), fromAcc)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SignTx:%v , err:%v", c.BaseURL, err.Error())
	}
	logger.Info("SignTx SendETHTransactWithLock", "")
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SendTransaction:%v , err:%v", c.BaseURL, err.Error())
	}
	logger.Info("SendTransaction SendETHTransactWithLock", "")
	return signedTx.Hash(), err
}

func (c *Client) ExecuteETHTransact(prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (*types.Receipt, error) {
	client, err := c.GetETHClient()
	if err != nil {
		return nil, fmt.Errorf("GetETHClient:%v , err:%v", c.BaseURL, err.Error())
	}
	signedTx, err := c.SendETHTransact(client, prkHex, from, to, value, input)
	if err != nil {
		return nil, fmt.Errorf("sendETHTransact:%v , err:%v", c.BaseURL, err.Error())
	}
	start := time.Now()
	txReceipt, err := eth.WaitForTxReceipt(client, signedTx)
	if execTime := time.Since(start); execTime > 2*time.Second || err != nil {
		logger.Info("wait mint a tx too long or err ", "", zap.Any("hash", signedTx.Hex()),
			zap.Any("BaseURL", c.BaseURL), zap.Any("execTime", execTime), zap.Any("err", err))
	}
	if err != nil {
		return nil, fmt.Errorf("WaitForTxReceipt:%v , err:%v", c.BaseURL, err.Error())
	}
	return txReceipt, err
}

func (c *Client) signAndSendTx(prkHex string, pbkHex common.Address, transact *accounts.Transaction) (*zktypes.Receipt, error) {
	chainID, err := c.GetChainID()
	if err != nil {
		return nil, err
	}

	zkClient, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}

	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		pbkHex,
		*transact,
	)
	if err != nil {
		return nil, err
	}

	prkBytes, err := hex.DecodeString(prkHex)
	if err != nil {
		return nil, err
	}
	baseSigner, err := accounts.NewBaseSignerFromRawPrivateKey(prkBytes, int64(chainID))
	if err != nil {
		return nil, err
	}
	signer := accounts.Signer(baseSigner)
	rawTx, err := c.SignTransaction(signer, preparedTx)
	if err != nil {
		return nil, err
	}
	hash, err := zkClient.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return nil, err
	}
	start := time.Now()
	tx, err := c.WaitMined(context.Background(), hash)
	if execTime := time.Since(start); execTime > 2*time.Second || err != nil {
		logger.Info("wait mint a tx too long or err ", "", zap.Any("hash", hash),
			zap.Any("BaseURL", c.BaseURL), zap.Any("execTime", execTime), zap.Any("err", err))
	}
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Client) WaitMined(ctx context.Context, txHash common.Hash) (*zktypes.Receipt, error) {
	zkClient, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}
	queryTicker := time.NewTicker(time.Second)
	maxRetry := 10
	checkExit := 3
	chainId, _ := c.GetChainID()
	if chainId == 45762 {
		queryTicker = time.NewTicker(500 * time.Millisecond)
		maxRetry = 20
		checkExit = 4
	} else if chainId == configs.BaseChainIDInt {
		queryTicker = time.NewTicker(2 * time.Millisecond)
		maxRetry = 20
		checkExit = 4
	}
	defer queryTicker.Stop()
	retry := 0
	for {
		retry++
		receipt, err := zkClient.TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil && receipt.BlockNumber != nil {
			return receipt, nil
		}
		if retry == checkExit {
			rpcResponse, pending, err := c.CheckTxPendingByHash(txHash)
			logger.Info("wait mint a tx not found ", "", zap.Any("hash", txHash),
				zap.Any("BaseURL", c.BaseURL),
				zap.Any("is not found", errors.Is(err, ethereum.NotFound)),
				zap.Any("pending", pending),
				zap.Any("rpcResponse", rpcResponse),
				zap.Any("err", err))
			if errors.Is(err, ethereum.NotFound) {
				return nil, err
			}
		}
		if retry > maxRetry {
			return nil, err
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}
func (c *Client) BlockByNumber(ctx context.Context, blockNumber uint64) (*zktypes.Block, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {

		ethClient, err := c.GetETHClient()
		if err != nil {
			return nil, err
		}

		blockEth, err := ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
		if err != nil {
			return nil, err
		}

		return &zktypes.Block{
			Header:          blockEth.Header(),
			Uncles:          blockEth.Uncles(),
			Withdrawals:     blockEth.Withdrawals(),
			Hash:            blockEth.Hash(),
			TotalDifficulty: blockEth.Difficulty(),
			ReceivedAt:      blockEth.ReceivedAt,
			ReceivedFrom:    blockEth.ReceivedFrom,
		}, nil
	}

	return block, nil
}

func (c *Client) CheckTxPendingByHash(txHash common.Hash) (*configs.RPCResponse, bool, error) {
	jsonData := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\",\"params\":[\"" + txHash.Hex() + "\"],\"id\":1}"
	resp, err := http.Post(c.BaseURL, "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("status code :%v = 200", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	var rpcResponse configs.RPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return nil, false, err
	}
	if rpcResponse.Result == nil {
		return nil, false, ethereum.NotFound
	}
	return &rpcResponse, rpcResponse.Result.BlockHash == nil, nil
}
