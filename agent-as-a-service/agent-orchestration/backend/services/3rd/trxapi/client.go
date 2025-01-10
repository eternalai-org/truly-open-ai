package trxapi

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"sync"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/mr-tron/base58"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func CreateTRONAddress() (string, string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
	prk := hex.EncodeToString(key.D.Bytes())
	if len(prk) != 64 {
		return "", "", errors.New("bad private key len")
	}
	addr = AddrEvmToTron(addr)
	return addr, prk, nil
}

func sha256Hex(msg string) (string, error) {
	bytes, err := hex.DecodeString(msg)
	if err != nil {
		return "", fmt.Errorf("invalid hex input: %w", err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:]), nil
}

func AddrTronToEvm(address string) string {
	decoded, err := base58.Decode(address)
	if err != nil {
		panic(fmt.Errorf("invalid hex input: %w", err))
	}
	if len(decoded) < 8 {
		panic(errors.New("invalid TRON address"))
	}
	return "0x" + hex.EncodeToString(decoded[1:len(decoded)-4])
}

func AddrEvmToTron(address string) string {
	if len(address) < 2 || address[:2] != "0x" {
		panic(errors.New("invalid EVM address"))
	}
	addr := "41" + address[2:]
	doubleSha1, err := sha256Hex(addr)
	if err != nil {
		panic(fmt.Errorf("error in first SHA-256 hash: %w", err))
	}
	doubleSha2, err := sha256Hex(doubleSha1)
	if err != nil {
		panic(fmt.Errorf("error in second SHA-256 hash: %w", err))
	}
	checkSum := doubleSha2[:8]
	fullAddr := addr + checkSum
	decoded, err := hex.DecodeString(fullAddr)
	if err != nil {
		panic(fmt.Errorf("failed to decode address: %w", err))
	}
	return base58.Encode(decoded)
}

type Client struct {
	ApiURL  string
	RpcURL  string
	GrpcURL string
	APIKey  string
	conn    *client.GrpcClient
	mtx     sync.Mutex
	client  *ethclient.Client
}

func (c *Client) Conn() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.conn == nil {
		var err error
		c.client, err = ethclient.Dial(c.RpcURL)
		if err != nil {
			panic(err)
		}
		conn := client.NewGrpcClient(c.GrpcURL)
		conn.SetAPIKey(c.APIKey)
		err = conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		c.conn = conn
	}
}

func (c *Client) buildUrl(resourcePath string) string {
	if resourcePath != "" {
		return c.ApiURL + "/" + resourcePath
	}
	return c.ApiURL
}

func (c *Client) buildRpcUrl(resourcePath string) string {
	if resourcePath != "" {
		return c.RpcURL + "/" + resourcePath
	}
	return c.RpcURL
}

func (c *Client) doWithoutAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) methodJSON(method string, apiURL string, jsonObject interface{}, result interface{}) error {
	var buffer io.Reader
	if jsonObject != nil {
		bodyBytes, _ := json.Marshal(jsonObject)
		buffer = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(method, apiURL, buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.doWithoutAuth(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *Client) SignTx(prk *ecdsa.PrivateKey, tx *core.Transaction) (*core.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	signature, err := crypto.Sign(hash, prk)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	return tx, nil
}

type TRC20TransfersResp struct {
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Num0  string         `json:"0"`
			Num1  string         `json:"1"`
			Num2  string         `json:"2"`
			From  string         `json:"from"`
			To    string         `json:"to"`
			Value numeric.BigInt `json:"value"`
		} `json:"result"`
		ResultType struct {
			From  string `json:"from"`
			To    string `json:"to"`
			Value string `json:"value"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionID string `json:"transaction_id"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

func (c *Client) GetTRC20Transfers(addr string) (*TRC20TransfersResp, error) {
	resp := TRC20TransfersResp{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl(fmt.Sprintf("/v1/contracts/%s/events?only_unconfirmed=false&only_confirmed=true&limit=200&event_name=Transfer", addr)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	for _, v := range resp.Data {
		v.Result.From = AddrEvmToTron(v.Result.From)
		v.Result.To = AddrEvmToTron(v.Result.To)
	}
	return &resp, nil
}

func (c *Client) Balance(addr string) (*big.Int, error) {
	c.Conn()
	acc, err := c.conn.GetAccount(addr)
	if err != nil {
		if err.Error() == "account not found" {
			return big.NewInt(0), nil
		}
		return nil, err
	}
	return big.NewInt(acc.GetBalance()), nil
}

func (c *Client) CheckBalance(addr string, fee *big.Int) error {
	c.Conn()
	acc, err := c.conn.GetAccount(addr)
	if err != nil {
		return err
	}
	if big.NewInt(acc.GetBalance()).Cmp(fee) <= 0 {
		return errors.New("not balance")
	}
	return nil
}
