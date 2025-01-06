package zkapi

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/evmapi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type BlockResp struct {
	time uint64
	hash string
}

func (b *BlockResp) Time() uint64 {
	return b.time
}

func (b *BlockResp) Hash() string {
	return b.hash
}

type Client struct {
	evmapi.BaseClient
	BaseURL          string
	BlockMap         map[uint64]*BlockResp
	BlockMapMtx      sync.Mutex
	chainID          uint64
	MinGasPrice      string
	PaymasterFeeZero bool
	PaymasterAddress string
	PaymasterToken   string
}

func (c *Client) getZkClient() (clients.Client, error) {
	zkClient, err := clients.Dial(c.BaseURL)
	if err != nil {
		return nil, err
	}
	return zkClient, nil
}

func (c *Client) ChainID() uint64 {
	return c.chainID
}

func (c *Client) Address() (string, string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
	prk := hex.EncodeToString(key.D.Bytes())
	if len(prk) != 64 {
		return "", "", errors.New("bad private key len")
	}
	return addr, prk, nil
}

func (c *Client) BlockByNumber(blockNumber int64) (*zktypes.Block, error) {
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *Client) GetLastBlock() (int64, error) {
	client, err := c.getZkClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	lastBlock, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	lastNumber := lastBlock.Number.Int64()
	return lastNumber, nil
}

func (c *Client) getBlockTime(n uint64) (uint64, error) {
	var ts uint64
	chainId, err := c.GetChainID()
	if err != nil {
		return ts, err
	}
	switch chainId {
	// nos regtest
	case 42070:
		{
			ts = 1686907878 + (n-1)*2
		}
		// nos mainnet
	case 42213:
		{
			ts = 1687812293 + (n-1)*2
		}
	case 8453:
		{
			ts = 1686789347 + n*2
		}
		// nos mainnet
	case 84531:
		{
			ts = 1688240016 + n*2
		}
	default:
		{
			block, err := c.getBlock(n)
			if err != nil {
				return ts, err
			}
			ts = block.Time()
		}
	}
	return ts, nil
}

func (c *Client) getBlock(n uint64) (*BlockResp, error) {
	if c.BlockMap == nil {
		c.BlockMap = map[uint64]*BlockResp{}
	}
	c.BlockMapMtx.Lock()
	blockResp, ok := c.BlockMap[n]
	c.BlockMapMtx.Unlock()
	if !ok {
		var blockInfoResp struct {
			Result *struct {
				Timestamp string `json:"timestamp"`
				Hash      string `json:"hash"`
			} `json:"result"`
		}
		err := c.postJSON(
			c.BaseURL,
			map[string]string{},
			map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      1,
				"method":  "eth_getBlockByNumber",
				"params": []interface{}{
					fmt.Sprintf("0x%s", big.NewInt(int64(n)).Text(16)),
					false,
				},
			},
			&blockInfoResp,
		)
		if err != nil {
			return nil, err
		}
		var timeBN *big.Int
		if strings.HasPrefix(blockInfoResp.Result.Timestamp, "0x") {
			timeBN, ok = big.NewInt(0).SetString(blockInfoResp.Result.Timestamp[2:], 16)
			if !ok {
				return nil, errors.New("wrong time")
			}
		} else {
			timeBN, ok = big.NewInt(0).SetString(blockInfoResp.Result.Timestamp, 10)
			if !ok {
				return nil, errors.New("wrong time")
			}
		}
		c.BlockMapMtx.Lock()
		//
		c.BlockMap[n] = &BlockResp{
			time: timeBN.Uint64(),
			hash: blockInfoResp.Result.Hash,
		}
		blockResp = c.BlockMap[n]
		//
		for i := n - 2000; i < n-1000; i++ {
			_, ok := c.BlockMap[i]
			if ok {
				delete(c.BlockMap, i)
			}
		}
		c.BlockMapMtx.Unlock()
	}
	return blockResp, nil
}

func (c *Client) doWithAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) postJSON(apiURL string, headers map[string]string, jsonObject interface{}, result interface{}) error {
	bodyBytes, _ := json.Marshal(jsonObject)
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.doWithAuth(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.Unmarshal(bodyBytes, result)
	}
	return nil
}

func (c *Client) WaitMined(hash string) error {
	client, err := c.getZkClient()
	if err != nil {
		return err
	}
	defer client.Close()
	r, err := client.WaitMined(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return errors.New("transaction is not Successful")
	}
	return nil
}

func (c *Client) parsePrkAuth(prkHex string) (common.Address, *ecdsa.PrivateKey, error) {
	prk, err := crypto.HexToECDSA(prkHex)
	if err != nil {
		return common.Address{}, nil, err
	}
	pbk, ok := prk.Public().(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, errors.New("error casting public key to ECDSA")
	}
	pbkHex := crypto.PubkeyToAddress(*pbk)
	return pbkHex, prk, nil
}

func (c *Client) GetChainID() (uint64, error) {
	if c.chainID > 0 {
		return c.chainID, nil
	}
	client, err := c.getZkClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	c.chainID = chainID.Uint64()
	return c.chainID, nil
}

func (c *Client) getGasPrice() (*big.Int, error) {
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	chainID, err := c.GetChainID()
	if err != nil {
		return nil, err
	}
	var gasPrice *big.Int
	switch chainID {
	case 22215, 22213:
		{
			gasPrice = big.NewInt(10000000000) // 10 gwei
		}
	case 42070, 42069, 42213:
		{
			gasPrice = big.NewInt(2000000000) // 2 gwei
		}
	case 42225:
		{
			gasPrice = big.NewInt(100000) // 0.0001 gwei
		}
	case 43337:
		{
			gasPrice = big.NewInt(100000) // 1 gwei
		}
	case 45454:
		{
			gasPrice = big.NewInt(358000000000000) // 358000 gwei
		}
	default:
		{
			if c.MinGasPrice != "" {
				minGasPriceNum, err := strconv.ParseInt(c.MinGasPrice, 10, 0)
				if err == nil {
					gasPrice = big.NewInt(minGasPriceNum)
				}
			}
			if gasPrice.Cmp(big.NewInt(0)) <= 0 {
				gasPrice, err = client.SuggestGasPrice(context.Background())
				if err != nil {
					return nil, err
				}
				switch chainID {
				case 42161:
					{
						gasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))
					}
				default:
					{
						gasPrice = gasPrice.Mul(gasPrice, big.NewInt(12))
					}
				}
				gasPrice = gasPrice.Div(gasPrice, big.NewInt(10))
			}
		}
	}
	return gasPrice, nil
}

func (c *Client) getGasTipCap() (*big.Int, error) {
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	chainID, err := c.GetChainID()
	if err != nil {
		return nil, err
	}
	var gasPrice *big.Int
	switch chainID {
	case 22215, 22213:
		{
			gasPrice = big.NewInt(10000000000) // 10 gwei
		}
	case 42070, 42069, 42213:
		{
			gasPrice = big.NewInt(2000000000) // 2 gwei
		}
	case 42225:
		{
			gasPrice = big.NewInt(100000) // 0.0001 gwei
		}
	case 43337:
		{
			gasPrice = big.NewInt(100000) // 1 gwei
		}
	case 43338:
		{
			gasPrice = big.NewInt(1100000000) // 1 gwei
		}
	case 45454:
		{
			gasPrice = big.NewInt(358000000000000) // 358000 gwei
		}
	default:
		{
			gasPrice, err = client.SuggestGasTipCap(context.Background())
			if err != nil {
				return nil, err
			}
			gasPrice = gasPrice.Mul(gasPrice, big.NewInt(12))
			gasPrice = gasPrice.Div(gasPrice, big.NewInt(10))
		}
	}
	return gasPrice, nil
}

func (c *Client) GetCachedGasPriceAndTipCap() (*big.Int, *big.Int, error) {
	cachedGasPrice, err := c.getGasPrice()
	if err != nil {
		return nil, nil, err
	}
	cachedGasTipCap, err := c.getGasTipCap()
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(
		cachedGasPrice.Text(10),
		cachedGasTipCap.Text(10),
	)
	return cachedGasPrice, cachedGasTipCap, nil
}

func (c *Client) PopulateTransaction(ctx context.Context, address common.Address, tx accounts.Transaction) (*zktypes.Transaction712, error) {
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	gasPrice, err := c.getGasPrice()
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

func (c *Client) ValidateMessageSignature(msg string, signatureHex string, signer string) error {
	msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgBytes := []byte(msg)
	msgHash := crypto.Keccak256Hash(
		msgBytes,
	)
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}
	if signature[crypto.RecoveryIDOffset] > 1 {
		signature[crypto.RecoveryIDOffset] -= 27
	}
	sigPublicKey, err := crypto.SigToPub(msgHash.Bytes(), signature)
	if err != nil {
		return err
	}
	pbkHex := crypto.PubkeyToAddress(*sigPublicKey)
	if !strings.EqualFold(pbkHex.Hex(), signer) {
		return errors.New("not valid signer")
	}
	return nil
}

func (c *Client) DeployContract(prkHex string, dataBin string, contructorHex string) (string, string, error) {
	client, err := c.getZkClient()
	if err != nil {
		return "", "", err
	}
	defer client.Close()
	tx := accounts.CreateTransaction{
		Bytecode: common.FromHex(dataBin),
		Calldata: common.FromHex(contructorHex),
	}
	preparedTx, err := tx.ToTransaction(accounts.DeployContract, nil)
	if err != nil {
		return "", "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", "", err
	}
	prkBytes, err := hex.DecodeString(prkHex)
	if err != nil {
		return "", "", err
	}
	baseSigner, err := accounts.NewBaseSignerFromRawPrivateKey(prkBytes, int64(chainID))
	if err != nil {
		return "", "", err
	}
	signer := accounts.Signer(baseSigner)
	rawTx, err := c.SignTransaction(signer, preparedTx.ToTransaction712(signer.Address()))
	if err != nil {
		return "", "", err
	}
	hash, err := client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", "", err
	}
	receipt, err := client.WaitMined(context.Background(), hash)
	if err != nil {
		return "", "", err
	}
	return hash.Hex(), receipt.ContractAddress.Hex(), nil
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
		helpers.HexToAddress(c.PaymasterAddress),
		&zktypes.ApprovalBasedPaymasterInput{
			Token:            helpers.HexToAddress(c.PaymasterToken),
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
		helpers.HexToAddress(c.PaymasterAddress),
		&zktypes.ApprovalBasedPaymasterInput{
			Token:            helpers.HexToAddress(c.PaymasterToken),
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

func (c *Client) validateAddress(address string) {
	if !(strings.EqualFold(address, helpers.HexToAddress(address).Hex())) {
		panic("wrong address")
	}
}

func (c *Client) TransactionConfirmed(hash string) error {
	err := c.WaitMined(hash)
	if err != nil {
		return errs.NewError(err)
	}
	return err
}

func (c *Client) Transfer(prkHex string, toAddr string, amount string, includeFee bool, wait bool) (string, error) {
	c.validateAddress(toAddr)
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
	}
	value, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		return "", errors.New("amount is invalid")
	}
	pbkHex, _, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	client, err := c.getZkClient()
	if err != nil {
		return "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	toAddress := helpers.HexToAddress(toAddr)
	gasNumber, err := client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From:  pbkHex,
				To:    &toAddress,
				Data:  []byte{},
				Value: value,
			},
			Meta: nil,
		},
	)
	if err != nil {
		return "", err
	}
	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		pbkHex,
		accounts.Transaction{
			GasFeeCap: gasPrice,
			GasTipCap: gasPrice,
			Gas:       gasNumber,
			To:        &toAddress,
			Value:     value,
			Data:      []byte{},
			Meta:      nil,
		},
	)
	if err != nil {
		return "", err
	}
	prkBytes, err := hex.DecodeString(prkHex)
	if err != nil {
		return "", err
	}
	baseSigner, err := accounts.NewBaseSignerFromRawPrivateKey(prkBytes, chainID.Int64())
	if err != nil {
		return "", err
	}
	signer := accounts.Signer(baseSigner)
	rawTx, err := c.SignTransaction(signer, preparedTx)
	if err != nil {
		return "", err
	}
	hash, err := client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}
	if wait {
		c.WaitMined(hash.Hex())
	}
	return hash.Hex(), nil
}

func (c *Client) InscribeTxs(txHashs []string) (string, error) {
	return "", nil
}

func (c *Client) Transact(contractAddr string, prkHex string, dataBytes []byte, value *big.Int) (string, error) {
	if value == nil {
		value = common.Big0
	}
	pbkHex, _, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	client, err := c.getZkClient()
	if err != nil {
		return "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	gasNumber, err := client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From:  pbkHex,
				To:    &contractAddress,
				Data:  []byte{},
				Value: value,
			},
			Meta: nil,
		},
	)
	if err != nil {
		return "", err
	}
	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		pbkHex,
		accounts.Transaction{
			GasFeeCap: gasPrice,
			GasTipCap: gasPrice,
			Gas:       gasNumber,
			To:        &contractAddress,
			Data:      []byte{},
			Meta:      nil,
			Value:     value,
		},
	)
	if err != nil {
		return "", err
	}
	prkBytes, err := hex.DecodeString(prkHex)
	if err != nil {
		return "", err
	}
	baseSigner, err := accounts.NewBaseSignerFromRawPrivateKey(prkBytes, chainID.Int64())
	if err != nil {
		return "", err
	}
	signer := accounts.Signer(baseSigner)
	rawTx, err := c.SignTransaction(signer, preparedTx)
	if err != nil {
		return "", err
	}
	hash, err := client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}
	return hash.Hex(), nil
}
