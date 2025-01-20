package btcapi

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/pkg/errors"
)

func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

const MinSat = 546

type ExecRequest struct {
	Args []string `json:"args"`
}

type InscriptionT1 struct {
	InscID      string
	BlockHeight uint64
}

type InscriptionT2 struct {
	Offset int64  `json:"offset"` //BigNumber
	ID     string `json:"id"`
}

type CreateSendTxRequest struct {
	// PrivateKey         string                   `json:"-"`             //buffer
	PrivateKeyStr      string                     `json:"privateString"` //buffer
	SenderAddress      string                     `json:"senderAddress"`
	Utxos              []Utxo                     `json:"utxos"`
	Inscriptions       map[string][]InscriptionT2 `json:"inscriptions"`
	SendInscriptionID  string                     `json:"sendInscriptionID"`
	ReceiverInsAddress string                     `json:"receiverInsAddress"`
	SendAmount         string                     `json:"sendAmount"`
	FeeRatePerByte     int                        `json:"feeRatePerByte"`
	Network            int                        `json:"network"`
}

type CreateSendTxMultiRequest struct {
	// PrivateKey         string                   `json:"-"`             //buffer
	PrivateKeyStr  string                     `json:"privateString"` //buffer
	SenderAddress  string                     `json:"senderAddress"`
	Utxos          []Utxo                     `json:"utxos"`
	Inscriptions   map[string][]InscriptionT2 `json:"inscriptions"`
	PaymentInfos   []PaymentInfo              `json:"paymentInfos"`
	FeeRatePerByte int                        `json:"feeRatePerByte"`
	Network        int                        `json:"network"`
}

type CreateSendTxMultiInscRequest struct {
	PrivateKeyStr    string                     `json:"privateString"` // wif string
	SenderAddress    string                     `json:"senderAddress"`
	Utxos            []Utxo                     `json:"utxos"`
	Inscriptions     map[string][]InscriptionT2 `json:"inscriptions"`
	InscPaymentInfos []InscPaymentInfo          `json:"inscPaymentInfos"`
	PaymentInfos     []PaymentInfo              `json:"paymentInfos"`
	FeeRatePerByte   int                        `json:"feeRatePerByte"`
	Network          int                        `json:"network"`
}

type CreateRawTxTransferSRC20Request struct {
	PublicKey       string                     `json:"publicKey"` // wif string
	SenderAddress   string                     `json:"senderAddress"`
	Utxos           []Utxo                     `json:"utxos"`
	Inscriptions    map[string][]InscriptionT2 `json:"inscriptions"`
	Data            string                     `json:"data"`
	ReceiverAddress string                     `json:"receiverAddress"`
	PaymentInfos    []PaymentInfo              `json:"paymentInfos"`
	FeeRatePerByte  int                        `json:"feeRatePerByte"`
	Network         int                        `json:"network"`
}

type CreateTransferSRC20ScriptRequest struct {
	Data      string `json:"data"`
	SecretKey string `json:"secretKey"`
}
type PaymentInfo struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type InscPaymentInfo struct {
	Address string `json:"address"`
	InscID  string `json:"inscID"`
}

type Utxo struct {
	TxHash    string `json:"tx_hash"`
	TxOutputN int64  `json:"tx_output_n"`
	Value     string `json:"value"`
}

type CreateSendTxResponse struct {
	ErrorCode string `json:"errorCode"`
	Error     string `json:"error"`
	Data      *struct {
		TxID          string `json:"txID"`
		TxHex         string `json:"txHex"`
		Fee           string `json:"fee"`
		ChangeAmount  string `json:"changeAmount"`
		SelectedUTXOs []Utxo `json:"selectedUTXOs"`
	} `json:"data"`
}

// type Inscription struct {
// 	InscID      string
// 	BlockHeight uint64
// }

type InscribeTxRequest struct {
	// PrivateKey         string                   `json:"-"`             //buffer
	PrivateKeyStr  string                     `json:"privateString"` //buffer
	SenderAddress  string                     `json:"senderAddress"`
	Utxos          []Utxo                     `json:"utxos"`
	Inscriptions   map[string][]InscriptionT2 `json:"inscriptions"`
	Data           string                     `json:"data"`
	FeeRatePerByte int                        `json:"feeRatePerByte"`
	Network        int                        `json:"network"`
}

type InscribeTxResponse struct {
	ErrorCode string `json:"errorCode"`
	Error     string `json:"error"`
	Data      *struct {
		CommitTxID    string `json:"commitTxID"`
		CommitTxHex   string `json:"commitTxHex"`
		RevealTxID    string `json:"revealTxID"`
		RevealTxHex   string `json:"revealTxHex"`
		Fee           string `json:"totalFee"`
		SelectedUTXOs []Utxo `json:"selectedUTXOs"`
		NewUTXOs      []Utxo `json:"newUTXOs"`
	} `json:"data"`
}

type CreateRawTxResponse struct {
	ErrorCode string `json:"errorCode"`
	Error     string `json:"error"`
	Data      *struct {
		Base64Psbt    string `json:"base64Psbt"`
		Fee           string `json:"fee"`
		SelectedUTXOs []Utxo `json:"selectedUTXOs"`
		ChangeAmount  string `json:"changeAmount"`

		IndicesToSign []int `json:"indicesToSign"`
	} `json:"data"`
}

type CreateTransferSRC20ScriptResponse struct {
	ErrorCode string   `json:"errorCode"`
	Error     string   `json:"error"`
	Data      []string `json:"data"`
}

func (c *Client) CreateSendTxSendMultiInsc(req CreateSendTxMultiInscRequest) (CreateSendTxResponse, error) {
	var rs CreateSendTxResponse
	err := c.postJSON(
		fmt.Sprintf("%s/create-tx-send-insc-multi", c.SdkUrl),
		map[string]string{},
		&req,
		&rs,
	)
	if err != nil {
		return rs, errs.NewError(err)
	}
	return rs, nil
}

type CreateOrdInscImgResp struct {
	ErrorCode string `json:"errorCode"`
	Error     string `json:"error"`
	Data      *struct {
		CommitTxHex string `json:"commitTxHex"`
		CommitTxID  string `json:"commitTxID"`
		RevealTxHex string `json:"revealTxHex"`
		RevealTxID  string `json:"revealTxID"`
	} `json:"data"`
}

func (c *Client) CreateOrdInscImgParam(req CreateOrdInscImgDtoRequest) (CreateOrdInscImgResp, error) {
	var rs CreateOrdInscImgResp
	err := c.postJSON(
		fmt.Sprintf("%s/create-ord-insc-img", c.SdkUrl),
		map[string]string{},
		&req,
		&rs,
	)
	if err != nil {
		return rs, errs.NewError(err)
	}
	return rs, nil
}

func (u *Client) getBlockHeightFromBlockStream() (uint64, error) {
	url := u.BlockstreamUrl + "/api/blocks/tip/height"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode == 429 {
		return 0, errors.New("429 Too Many Requests")
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, "RPC error") {
		return 0, errors.New(bodyStr)
	}
	blockHeight, err := strconv.ParseUint(bodyStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return blockHeight, nil
}

func (u *Client) GetBlockCountFromService() (uint64, error) {
	blockHeight, err := u.getBlockHeightFromBlockStream()
	if err != nil {
		return 0, errs.NewError(err)
	}
	return blockHeight, nil
}

type UTXO struct {
	TxHash    string `json:"tx_hash"`
	TxOutputN int    `json:"tx_output_n"`
	Value     uint64 `json:"value"`
}

type UTXOIns struct {
	InscID    string `json:"inscID"`
	TxHash    string `json:"tx_hash"`
	TxOutputN int    `json:"tx_output_n"`
	Value     uint64 `json:"value"`
	Offset    uint64 `json:"offset"`
}

type BTCFullnodeConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsHttps  string `json:"is_https"`
}

func (u *Client) BuildBTCClient(cfg *BTCFullnodeConfig) (*rpcclient.Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         cfg.Host,
		User:         cfg.Username,
		Pass:         cfg.Password,
		HTTPPostMode: true,                     // Bitcoin core only supports HTTP POST mode
		DisableTLS:   !(cfg.IsHttps == "true"), // Bitcoin core does not provide TLS by default
	}
	return rpcclient.New(connCfg, nil)
}

func (u *Client) GetUTXOs(btcClient *rpcclient.Client, addresses []string, chainCfg *chaincfg.Params, minConfirmation, maxConfirmation int) ([]btcjson.ListUnspentResult, error) {
	trackingBTCAddresses := []btcutil.Address{}
	for _, addressStr := range addresses {
		address, err := btcutil.DecodeAddress(addressStr, chainCfg)
		if err != nil {
			continue
		}
		trackingBTCAddresses = append(trackingBTCAddresses, address)
	}
	listUnspentResults, err := btcClient.ListUnspentMinMaxAddresses(minConfirmation, maxConfirmation, trackingBTCAddresses)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return listUnspentResults, nil
}

type BitcoinParams struct {
	FirstScannedBTCBlkHeight uint64

	MasterPubKeys         [][]byte
	GeneralMultisigWallet string
	NumRequiredSigs       int
	TotalSigs             int
	MinDepositAmount      uint64
	MinWithdrawAmount     uint64
	DepositFee            uint64

	ChainParam *chaincfg.Params
	InputSize  int // for multisig
	OutputSize int
	MaxTxSize  int
	MaxFeeRate int

	TaprootInputSize int
}

var BitcoinParamsMaintest = &BitcoinParams{
	FirstScannedBTCBlkHeight: uint64(787496),
	MasterPubKeys: [][]byte{
		[]byte{0x2, 0xe, 0x8, 0xa, 0xe3, 0xcf, 0xf5, 0x1d, 0xc3, 0xc0, 0x83, 0xaf, 0xa9, 0x24, 0x71, 0x9c, 0x2f, 0xca, 0x62, 0x89, 0x74, 0x70, 0xb4, 0x8b, 0x9, 0x51, 0x3, 0x6f, 0x32, 0x9e, 0xdb, 0x5f, 0xe7},
		[]byte{0x3, 0xc2, 0x3c, 0x3d, 0x6f, 0x83, 0xbe, 0xc9, 0x56, 0xde, 0x6a, 0x54, 0x90, 0xac, 0x2d, 0xe7, 0xee, 0x5c, 0xf8, 0x63, 0x22, 0x84, 0x9c, 0x61, 0xed, 0x62, 0x5b, 0x69, 0x8f, 0x4a, 0x4a, 0xee, 0x49},
		[]byte{0x3, 0x3b, 0x75, 0x80, 0x77, 0x8f, 0x4d, 0x2e, 0x46, 0x20, 0x6a, 0xd5, 0x32, 0x66, 0x18, 0xb6, 0xd6, 0x4c, 0x46, 0x1a, 0xe, 0x47, 0xb2, 0x5a, 0x77, 0xad, 0x72, 0xdc, 0x56, 0x4e, 0xa6, 0xca, 0xdb},
		[]byte{0x2, 0xfc, 0x81, 0x32, 0xba, 0xb2, 0x85, 0x71, 0x82, 0x3f, 0x82, 0x3d, 0x74, 0xe5, 0xd4, 0xa2, 0xff, 0xcb, 0xb, 0xe7, 0x2c, 0x49, 0x63, 0xb3, 0x73, 0x75, 0xf7, 0xc4, 0x41, 0xf9, 0x3e, 0xda, 0x96},
		[]byte{0x2, 0xb2, 0x14, 0xd8, 0x19, 0x77, 0x31, 0x59, 0xa3, 0xae, 0x9c, 0x30, 0xf9, 0x85, 0xa4, 0xe0, 0x56, 0x1d, 0x98, 0x9d, 0xf6, 0x27, 0xfb, 0xbd, 0xd6, 0x9d, 0x33, 0xdd, 0xa7, 0x25, 0x38, 0x35, 0xb5},
		[]byte{0x3, 0xb8, 0xc6, 0xf2, 0x80, 0x9d, 0xe5, 0xd, 0x6a, 0x43, 0x57, 0xb9, 0xac, 0xce, 0xaa, 0xd, 0x8, 0xda, 0xd4, 0x75, 0xd9, 0x6a, 0xbf, 0x70, 0x14, 0xe7, 0x2a, 0xeb, 0x68, 0xe0, 0xb1, 0xa4, 0xb4},
		[]byte{0x2, 0xee, 0x76, 0x26, 0xa3, 0x4f, 0xd, 0xb7, 0x57, 0x21, 0xa4, 0x44, 0x8c, 0xa8, 0x6e, 0x59, 0xa1, 0x32, 0x2d, 0xa6, 0xbe, 0xf8, 0x86, 0xbf, 0x64, 0xa5, 0x94, 0xa7, 0xed, 0x20, 0xfd, 0xc2, 0x52},
	},
	GeneralMultisigWallet: "bc1qajyp9ekpepmhftxq8aeps4cv8gjkat00nfk9lplqec7mevhv4z6qxy37z8",
	NumRequiredSigs:       5,
	TotalSigs:             7,
	MinDepositAmount:      uint64(0),
	MinWithdrawAmount:     uint64(0),
	DepositFee:            uint64(10000),

	ChainParam:       &chaincfg.MainNetParams,
	InputSize:        192,
	OutputSize:       43,
	MaxTxSize:        51200, // 50 KB
	TaprootInputSize: 68,
	// MaxFeeRate: 150,
}

func toSat(amount float64) uint64 {
	amtFloat := new(big.Float).SetFloat64(amount*1e8 + 0.5)
	res, _ := amtFloat.Uint64()
	return res
}

type GetTxInfoFromBlockStream struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid    string `json:"txid"`
		Vout    int    `json:"vout"`
		Prevout struct {
			Scriptpubkey        string `json:"scriptpubkey"`
			ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
			ScriptpubkeyType    string `json:"scriptpubkey_type"`
			ScriptpubkeyAddress string `json:"scriptpubkey_address"`
			Value               int    `json:"value"`
		} `json:"prevout"`
		Scriptsig             string   `json:"scriptsig"`
		ScriptsigAsm          string   `json:"scriptsig_asm"`
		Witness               []string `json:"witness"`
		IsCoinbase            bool     `json:"is_coinbase"`
		Sequence              int      `json:"sequence"`
		InnerWitnessscriptAsm string   `json:"inner_witnessscript_asm"`
	} `json:"vin"`
	Vout []struct {
		Scriptpubkey        string `json:"scriptpubkey"`
		ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
		ScriptpubkeyType    string `json:"scriptpubkey_type"`
		ScriptpubkeyAddress string `json:"scriptpubkey_address"`
		Value               int    `json:"value"`
	} `json:"vout"`
	Size   int `json:"size"`
	Weight int `json:"weight"`
	Fee    int `json:"fee"`
	Status struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight int    `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   int    `json:"block_time"`
	} `json:"status"`
}

func (u *Client) getAddressMempoolTxsFromBlockStream(address string) ([]GetTxInfoFromBlockStream, error) {
	var result []GetTxInfoFromBlockStream
	time.Sleep(100 * time.Millisecond)
	url := u.BlockstreamUrl + fmt.Sprintf("/api/address/%s/txs/mempool", address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, "RPC error") {
		return nil, errors.New(bodyStr)
	}
	respUTXOs := []GetTxInfoFromBlockStream{}
	err = json.Unmarshal(body, &respUTXOs)
	if err != nil {
		return nil, err
	}
	result = append(result, respUTXOs...)

	return result, nil
}

type UTXOFromBlockStream struct {
	Txid   string `json:"txid"`
	Vout   int    `json:"vout"`
	Status struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight int    `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   int    `json:"block_time"`
	} `json:"status"`
	Value int `json:"value"`
}

func (u *Client) getUTXOFromBlockStream(address string, minConfirm, maxConfirm uint64, curBlockHeight uint64) ([]*UTXO, error) {
	url := u.BlockstreamUrl + "/api/address/" + address + "/utxo"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 429 {
		return nil, errors.New("429 Too Many Requests")
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, "RPC error") {
		return nil, errors.New(bodyStr)
	}

	respUTXOs := []UTXOFromBlockStream{}
	err = json.Unmarshal(body, &respUTXOs)
	if err != nil {
		return nil, err
	}

	result := []*UTXO{}
	for _, utxo := range respUTXOs {
		confirm := uint64(0)
		if utxo.Status.Confirmed {
			confirm = uint64(int(curBlockHeight) - int(utxo.Status.BlockHeight) + 1)
		}

		if confirm >= minConfirm && confirm <= maxConfirm {
			result = append(result, &UTXO{
				TxHash:    utxo.Txid,
				TxOutputN: utxo.Vout,
				Value:     uint64(utxo.Value),
			})
		}
	}

	return result, nil
}

func (u *Client) getUTXONoPendingFromBlockStreamV2(address string, minConfirm uint64, maxConfirm uint64, curBlockHeight uint64) ([]*UTXO, error) {
	utxos, err := u.getUTXOFromBlockStream(address, minConfirm, maxConfirm, curBlockHeight)
	if err != nil {
		fmt.Println("getUTXOFromBlockStream err", err)
		return nil, err
	}
	pendingTxs, err := u.getAddressMempoolTxsFromBlockStream(address)
	if err != nil {
		fmt.Println("getAddressMempoolTxsFromBlockStream err", err)
		return nil, err
	}
	pendingUTXOMap := map[string]bool{}
	for _, tx := range pendingTxs {
		for _, input := range tx.Vin {
			if input.Prevout.ScriptpubkeyAddress != address {
				continue
			}

			key := fmt.Sprintf("%v:%v", input.Txid, input.Vout)
			pendingUTXOMap[key] = true
		}
	}
	utxosNoPending := []*UTXO{}
	for _, utxo := range utxos {
		key := fmt.Sprintf("%v:%v", utxo.TxHash, utxo.TxOutputN)
		if !pendingUTXOMap[key] {
			utxosNoPending = append(utxosNoPending, utxo)
		}
	}
	return utxosNoPending, nil
}

func (u *Client) convertUTXOToSDKFormat(utxos []*UTXO) []Utxo {
	utxoSDK := []Utxo{}
	for _, utxo := range utxos {
		utxoSDK = append(utxoSDK, Utxo{
			TxHash:    utxo.TxHash,
			TxOutputN: int64(utxo.TxOutputN),
			Value:     fmt.Sprintf("%v", utxo.Value),
		})
	}
	return utxoSDK
}

func (u *Client) CheckBVMInscs(txIDs []string) ([]bool, error) {
	var resp struct {
		Result []bool `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	url := fmt.Sprintf("https://tc-node-public-webstat.trustless.computer/check-validtc?btc-txs=%s", strings.Join(txIDs, ","))
	// fmt.Printf("CheckBVMInscs URL: %v", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Result == nil || resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	return resp.Result, nil
}

type TCInscription struct {
	Offset int64  `json:"offset"` //BigNumber
	ID     string `json:"id"`
}

func (u *Client) GetListUnspentBVMInscsForSend(btcAddress string, inscUTXOs []*UTXOIns) ([]InscriptionT1, []Utxo, map[string][]InscriptionT2, error) {
	btcBestBlockHeight, err := u.GetBlockCountFromService()
	if err != nil {
		return nil, nil, nil, errs.NewError(err)
	}
	utxos, err := u.getUTXONoPendingFromBlockStreamV2(btcAddress, uint64(1), uint64(9999999), btcBestBlockHeight)
	if err != nil {
		return nil, nil, nil, errs.NewError(err)
	}
	// for returning api
	bvmInsc := []InscriptionT1{}
	// for create tx
	utxosSDK := u.convertUTXOToSDKFormat(utxos)
	inscriptionMapSDK := make(map[string][]InscriptionT2)
	for _, u := range inscUTXOs {
		key := fmt.Sprintf("%v:%v", u.TxHash, u.TxOutputN)
		inscriptionMapSDK[key] = []InscriptionT2{
			{
				ID:     u.InscID,
				Offset: int64(u.Offset),
			},
		}
	}
	return bvmInsc, utxosSDK, inscriptionMapSDK, nil
}

func (u *Client) GetListUnspentBVMInscsForMint(btcAddress string) ([]InscriptionT1, []Utxo, map[string][]InscriptionT2, error) {
	btcBestBlockHeight, err := u.GetBlockCountFromService()
	if err != nil {
		return nil, nil, nil, errs.NewError(err)
	}
	utxos, err := u.getUTXONoPendingFromBlockStreamV2(btcAddress, uint64(1), uint64(9999999), btcBestBlockHeight)
	if err != nil {
		return nil, nil, nil, errs.NewError(err)
	}
	// for returning api
	bvmInsc := []InscriptionT1{}
	// for create tx
	utxosSDK := u.convertUTXOToSDKFormat(utxos)
	inscriptionMapSDK := make(map[string][]InscriptionT2)
	//
	return bvmInsc, utxosSDK, inscriptionMapSDK, nil
}

type FeeRates struct {
	FastestFee  int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
	EconomyFee  int `json:"economyFee"`
	MinimumFee  int `json:"minimumFee"`
}

func (u *Client) getFeeRateFromChain() (*FeeRates, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get("https://mempool.space/api/v1/fees/recommended")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	feeRateObj := &FeeRates{}
	err = json.Unmarshal(responseData, &feeRateObj)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return feeRateObj, nil

}

func (u *Client) GetFeeRate() (int, error) {
	// cache here:
	feeRate := 30
	feeRateCurrent, err := u.getFeeRateFromChain()
	if err != nil {
		return 0, errs.NewError(err)
	}
	feeRate = feeRateCurrent.FastestFee
	return feeRate, nil
}

func (u *Client) SendBVMInscs(btcAddress string, privateKey string, inscIDs []string, addresses []string, inscUTXOs []*UTXOIns, feeRate int) (string, string, error) {
	_, utxoSDK, inscriptionMapSDK, err := u.GetListUnspentBVMInscsForSend(btcAddress, inscUTXOs)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	if len(utxoSDK) <= 0 {
		return "", "", errs.NewError(errors.New("utxoSDK is empty"))
	}
	if feeRate <= 0 {
		feeRate, err = u.GetFeeRate()
		if err != nil {
			return "", "", errs.NewError(err)
		}
	}
	inscPaymentInfos := []InscPaymentInfo{}
	sendInscIDsMap := map[string]bool{}
	for i, inscID := range inscIDs {
		inscPaymentInfos = append(inscPaymentInfos, InscPaymentInfo{
			Address: addresses[i],
			InscID:  inscID,
		})
		sendInscIDsMap[inscID] = true
	}

	network := 1
	// if u.Config.ENV == "develop" {
	// 	network = 3
	// }

	param := CreateSendTxMultiInscRequest{
		PrivateKeyStr:    privateKey,
		SenderAddress:    btcAddress,
		Utxos:            utxoSDK,
		Inscriptions:     inscriptionMapSDK,
		InscPaymentInfos: inscPaymentInfos,
		PaymentInfos:     []PaymentInfo{},
		FeeRatePerByte:   feeRate,
		Network:          network,
	}
	respData, err := u.CreateSendTxSendMultiInsc(param)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	if respData.Error != "" {
		return "", "", errs.NewError(errors.New(respData.Error))
	}
	_, err = u.ParseTx(respData.Data.TxHex)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	_, err = u.BroadcastBTCTxByMempool(respData.Data.TxHex)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	return respData.Data.TxID, respData.Data.TxHex, nil
}

func (u *Client) ParseTx(data string) (*wire.MsgTx, error) {
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	tx := &wire.MsgTx{}
	err = tx.Deserialize(strings.NewReader(string(dataBytes)))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (u *Client) BroadcastBTCTxByMempool(hexTx string) (string, error) {
	url := u.MempoolUrl + "/api/tx"
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer([]byte(hexTx)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, "RPC error") {
		return "", errors.New(bodyStr)
	}
	return bodyStr, nil
}

type CreateOrdInscImgDtoRequest struct {
	Network         int                        `json:"network"`
	PrivateString   string                     `json:"privateString"`
	SenderAddress   string                     `json:"senderAddress"`
	Utxos           []Utxo                     `json:"utxos"`
	Inscriptions    map[string][]InscriptionT2 `json:"inscriptions"`
	Data            string                     `json:"data"`
	ContentType     string                     `json:"contentType"`
	FeeRatePerByte  int                        `json:"feeRatePerByte"`
	ReceiverAddress string                     `json:"receiverAddress"`
}

func (u *Client) CreateOrdInscImg(btcAddress string, privateKey string, feeRate int, receiverAddress string, data []byte) (string, string, error) {
	_, utxoSDK, _, err := u.GetListUnspentBVMInscsForMint(btcAddress)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	if len(utxoSDK) <= 0 {
		return "", "", errs.NewError(errors.New("utxo list empty"))
	}
	if feeRate <= 0 {
		feeRate, err = u.GetFeeRate()
		if err != nil {
			return "", "", errs.NewError(err)
		}
	}
	network := 1
	// if u.Config.ENV == "develop" {
	// 	network = 3
	// }
	// filter list utxos to reduce payload
	contentType := detectContentType(data)
	param := CreateOrdInscImgDtoRequest{
		PrivateString:   privateKey,
		SenderAddress:   btcAddress,
		Utxos:           utxoSDK,
		Inscriptions:    map[string][]InscriptionT2{},
		FeeRatePerByte:  feeRate,
		Network:         network,
		ReceiverAddress: receiverAddress,
		ContentType:     contentType,
		Data:            hex.EncodeToString(data),
	}
	respData, err := u.CreateOrdInscImgParam(param)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	var txHash1, txHash2 string
	txHash1, err = u.BroadcastBTCTxByMempool(respData.Data.CommitTxHex)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	time.Sleep(5 * time.Second)
	txHash2, err = u.BroadcastBTCTxByMempool(respData.Data.RevealTxHex)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	return respData.Data.RevealTxID, fmt.Sprintf("%s_%s", txHash1, txHash2), nil
}
