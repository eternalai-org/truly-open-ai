package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/logger"
	"decentralized-inference/internal/models"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"go.uber.org/zap"
	"io"
	"math/big"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	RPC              string
	chainID          *big.Int
	chainType        models.ChainType
	PaymasterFeeZero bool
	PaymasterAddress string
	PaymasterToken   string
	Client           *clients.Client
	ETHClient        *ethclient.Client
}
type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  *ResultResponse `json:"result"`
	ID      int             `json:"id"`
}

type ResultResponse struct {
	Hash          string  `json:"hash"`
	BlockHash     *string `json:"blockHash"`
	BlockNumber   *string `json:"blockNumber"`
	Number        *string `json:"number"`
	ChainID       string  `json:"chainId"`
	L1BlockNumber *string `json:"l1BlockNumber"`
}

var mapClient = map[string]*Client{}
var lock = sync.RWMutex{}

func NewClient(rpc string, chainType models.ChainType, paymentFeeZero bool, paymentMasterAddress string, paymentMasterToken string) (*Client, error) {

	lock.Lock()
	defer lock.Unlock()
	client, _ := mapClient[rpc]
	if client != nil {
		return client, nil
	}
	c, err := clients.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", rpc, err)
	}
	id, err := c.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID: %v ,rpc :%v", err, rpc)
	}

	client = &Client{
		RPC:              rpc,
		chainID:          id,
		chainType:        chainType,
		PaymasterFeeZero: paymentFeeZero,
		PaymasterAddress: paymentMasterAddress,
		PaymasterToken:   paymentMasterToken,
		Client:           c,
		ETHClient:        ethclient.NewClient(c.Client()),
	}
	mapClient[rpc] = client
	return client, nil
}

func GetAccountInfo(privKey string) (*ecdsa.PrivateKey, *common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	publicKeyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, &publicKeyAddress, nil
}

func (c *Client) GetGasPrice() (*big.Int, error) {
	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w ,rpc", err, c.RPC)
	}
	return gasPrice, nil
}

func (c *Client) PopulateTransaction(ctx context.Context, address common.Address, tx accounts.Transaction) (*zktypes.Transaction, error) {
	client := c.Client
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, err
	}
	if tx.ChainID == nil {
		tx.ChainID = c.chainID
	}
	tx.Nonce = new(big.Int).SetUint64(nonce)
	tx.GasFeeCap = gasPrice
	if tx.GasTipCap == nil {
		tx.GasTipCap = big.NewInt(0)
	}
	if tx.GasPerPubdata == nil {
		tx.GasPerPubdata = big.NewInt(utils.DefaultGasPerPubdataLimit.Int64())
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
	if tx.PaymasterParams != nil {
		paymasterParams, err := c.GetPaymasterParamsWithFee(big.NewInt(0).Mul(tx.GasFeeCap, big.NewInt(int64(tx.Gas))))
		if err != nil {
			panic(err)
		}
		tx.PaymasterParams = paymasterParams
	}
	return tx.ToTransaction(address), nil
}

func (c *Client) SignTransaction(signer *accounts.ECDSASigner, tx *zktypes.Transaction) ([]byte, error) {
	var gas uint64 = 0
	if tx.Gas != nil {
		gas = tx.Gas.Uint64()
	}
	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		signer.Address(),
		accounts.Transaction{
			To:              tx.To,
			Data:            tx.Data,
			Value:           tx.Value,
			Nonce:           tx.Nonce,
			GasTipCap:       tx.GasTipCap,
			GasFeeCap:       tx.GasFeeCap,
			Gas:             gas,
			ChainID:         tx.ChainID,
			GasPerPubdata:   tx.GasPerPubdata,
			CustomSignature: tx.CustomSignature,
			FactoryDeps:     tx.FactoryDeps,
			PaymasterParams: tx.PaymasterParams,
		},
	)
	if err != nil {
		return nil, err
	}
	typedData, err := preparedTx.TypedData()
	if err != nil {
		return nil, err
	}
	signature, err := signer.SignTypedData(context.Background(), typedData)
	if err != nil {
		return nil, err
	}
	return preparedTx.Encode(signature)
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
	tx.PaymasterParams = paymasterParams
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

	paymaster, err := c.GetPaymasterParams()
	if err != nil {
		return nil, err
	}
	if paymaster != nil || c.chainType == models.ChainTypeZkSync {
		transact, err := c.createZKTransact(from, to, value, input)
		if err != nil {
			return nil, fmt.Errorf("failed to create transact: %w", err)
		}
		// get nonce again
		tx, err := c.signAndSendTx(prkHex, from, transact)
		if err != nil {
			return nil, fmt.Errorf("failed to sign and transact: %w", err)
		}
		if tx.Receipt.Status == types.ReceiptStatusFailed {
			return tx, fmt.Errorf("transact failed with status %d", tx.Receipt.Status)
		}
		return tx, nil
	} else {
		txReceipt, err := c.ExecuteETHTransact(prkHex, from, to, value, input)
		if err != nil {
			return nil, err
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

func (c *Client) createZKTransact(from common.Address, to common.Address, value *big.Int, input []byte) (*accounts.Transaction, error) {
	gasNumber, err := c.Client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			From:            from,
			To:              &to,
			Value:           value,
			Data:            input,
			GasPerPubdata:   big.NewInt(utils.DefaultGasPerPubdataLimit.Int64()),
			PaymasterParams: c.PaymasterParams(),
		},
	)

	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, err
	}

	return &accounts.Transaction{
		GasFeeCap:       gasPrice,
		GasTipCap:       gasPrice,
		Gas:             gasNumber,
		To:              &to,
		Value:           value,
		Data:            input,
		GasPerPubdata:   big.NewInt(utils.DefaultGasPerPubdataLimit.Int64()),
		PaymasterParams: c.PaymasterParams(),
	}, err
}

func (c *Client) SendETHTransact(prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (common.Hash, error) {
	fromAcc, _, err := GetAccountInfo(prkHex)
	if err != nil {
		return common.Hash{}, errors.Join(err, errors.New("error while getting account info"))
	}
	tx, err := c.CreateEthTransaction(from, to, value, input)
	if err != nil {
		return common.Hash{}, fmt.Errorf("CreateEthTransaction :%v", err.Error())
	}
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(c.chainID), fromAcc)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SignTx:%v , err:%v", c.RPC, err.Error())
	}
	err = c.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SendTransaction:%v , err:%v", c.RPC, err.Error())
	}
	return signedTx.Hash(), err
}

func (c *Client) ExecuteETHTransact(prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (*types.Receipt, error) {
	signedTx, err := c.SendETHTransact(prkHex, from, to, value, input)
	if err != nil {
		return nil, fmt.Errorf("sendETHTransact:%v , err:%v", c.RPC, err.Error())
	}
	txReceipt, err := c.WaitMinedEthTx(signedTx)
	if err != nil {
		return nil, fmt.Errorf("WaitForTxReceipt:%v , err:%v", c.RPC, err.Error())
	}
	return txReceipt, err
}

func (c *Client) signAndSendTx(prkHex string, pbkHex common.Address, transact *accounts.Transaction) (*zktypes.Receipt, error) {
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
	baseSigner, err := accounts.NewECDSASignerFromRawPrivateKey(prkBytes, c.chainID)
	if err != nil {
		return nil, err
	}

	rawTx, err := c.SignTransaction(baseSigner, preparedTx)
	if err != nil {
		return nil, err
	}
	hash, err := c.Client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return nil, err
	}
	tx, err := c.WaitMined(hash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Client) WaitMined(txHash common.Hash) (*zktypes.Receipt, error) {
	queryTicker := time.NewTicker(2 * time.Second)
	defer queryTicker.Stop()
	retry := 0
	maxRetry := 1000
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}

		receipt, err := c.Client.TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil && receipt.BlockNumber != nil {
			return receipt, nil
		}
		retry++
		if retry%3 == 0 {
			_, _, err := c.CheckTxPendingByHash(txHash)
			if errors.Is(err, ethereum.NotFound) {
				return nil, err
			}
		}
		if retry%10 == 0 {
			logger.GetLoggerInstanceFromContext(context.Background()).Info("waiting get tx receipt", zap.Any("hash", txHash),
				zap.Any("rpc", c.RPC))
		}
		if retry > maxRetry {
			return nil, err
		}
	}
}

func (c *Client) WaitMinedEthTx(txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(2 * time.Second)
	defer queryTicker.Stop()
	retry := 0
	maxRetry := 1000
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}

		receipt, err := c.ETHClient.TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil && receipt.BlockNumber != nil {
			return receipt, nil
		}
		retry++
		if retry%3 == 0 {
			_, _, err := c.CheckTxPendingByHash(txHash)
			if errors.Is(err, ethereum.NotFound) {
				return nil, err
			}
		}
		if retry%10 == 0 {
			logger.GetLoggerInstanceFromContext(context.Background()).Info("waiting get tx receipt", zap.Any("hash", txHash),
				zap.Any("rpc", c.RPC))
		}
		if retry > maxRetry {
			return nil, err
		}
	}
}
func (c *Client) BlockByNumber(ctx context.Context, blockNumber uint64) (*zktypes.Block, error) {
	block, err := c.Client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *Client) CheckTxPendingByHash(txHash common.Hash) (*RPCResponse, bool, error) {
	jsonData := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\",\"params\":[\"" + txHash.Hex() + "\"],\"id\":1}"
	resp, err := http.Post(c.RPC, "application/json", bytes.NewBuffer([]byte(jsonData)))
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
	var rpcResponse RPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return nil, false, err
	}
	if rpcResponse.Result == nil {
		return nil, false, ethereum.NotFound
	}
	return &rpcResponse, rpcResponse.Result.BlockHash == nil, nil
}

func (c *Client) CreateEthTransaction(from common.Address, to common.Address, value *big.Int, input []byte) (*types.Transaction, error) {
	head, err := c.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest head: %w , rpc :%v", err, c.RPC)
	}

	gasLimit, err := c.Client.EstimateGas(
		context.Background(),
		ethereum.CallMsg{
			From:  from,
			To:    &to,
			Value: value,
			Data:  input,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("err when estimate gas at rpc :%v ,from:%v ,to %v , value:%v, input:%v, err:%v", c.RPC, from, &to, value.String(), common.Bytes2Hex(input), err.Error())
	}

	if c.chainID.String() != config.SubtensorEVMChainID {
		if gasLimit < head.GasLimit/2 {
			gasLimit = gasLimit * 2
		}
	} else {
		gasLimit = config.GasLimitDefault
	}

	nonce, err := c.Client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, fmt.Errorf("err when PendingNonceAt:%v , err:%v", c.RPC, err.Error())
	}
	var tx *types.Transaction
	if head.BaseFee != nil {
		// Estimate TipCap
		gasTipCap := big.NewInt(0)
		gasFeeCap := big.NewInt(10000000000)
		if c.chainID.String() != config.SubtensorEVMChainID {
			gasTipCap, err = c.ETHClient.SuggestGasTipCap(context.Background())
			if err != nil {
				return nil, fmt.Errorf("err when SuggestGasTipCap:%v , err:%v", c.RPC, err.Error())
			}
			gasFeeCap = new(big.Int).Add(
				gasTipCap,
				new(big.Int).Mul(head.BaseFee, big.NewInt(config.BaseFeeWiggleMultiplier)),
			)
		} else {
			gasFeeCap, err = c.ETHClient.SuggestGasPrice(context.Background())
			if err != nil {
				return nil, fmt.Errorf("err when SuggestGasPrice:%v , err:%v", c.RPC, err.Error())
			}
		}
		tx = types.NewTx(&types.DynamicFeeTx{
			To:        &to,
			Nonce:     nonce,
			GasFeeCap: gasFeeCap,
			GasTipCap: gasTipCap,
			Gas:       gasLimit,
			Value:     value,
			Data:      input,
		})
	} else {
		gasPrice, err := c.ETHClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("err when SuggestGasPrice:%v , err:%v", c.RPC, err.Error())
		}
		tx = types.NewTx(&types.LegacyTx{
			To:       &to,
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			Value:    value,
			Data:     input,
		})
	}
	return tx, nil
}
