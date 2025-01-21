package btcapi

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/blockcypher/gobcy/v2"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/pkg/errors"
)

type Client struct {
	Chain             string
	Network           string
	Token             string
	client            *gobcy.API
	QNUrl             string
	SdkUrl            string
	BlockstreamUrl    string
	MempoolUrl        string
	HirosoUrl         string
	BlockchainInfoUrl string
}

const (
	// This will calculate and include appropriate fees for your transaction to be included in the next 1-2 blocks
	PreferenceHigh = "high"
	// This will calculate and include appropriate fees for your transaction to be included in the next 3-6 blocks
	PreferenceMedium = "medium"
	// This will calculate and include appropriate fees for your transaction to be included in the next 7 or more blocks
	PreferenceLow = "low"
	// No fee
	PreferenceZero = "zero"
)

func (c *Client) getClient() (*gobcy.API, error) {
	if c.client == nil {
		client := gobcy.API{c.Token, c.Chain, c.Network}
		c.client = &client
	}
	return c.client, nil
}

func (c *Client) doWithAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) doWithoutAuth(req *http.Request) (*http.Response, error) {
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

func (c *Client) getJSON(url string, headers map[string]string, result interface{}) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.doWithAuth(req)
	if err != nil {
		return 0, fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return resp.StatusCode, json.NewDecoder(resp.Body).Decode(result)
	}
	return resp.StatusCode, nil
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

func (c *Client) Address() (string, string, error) {
	var networkParams *chaincfg.Params
	switch c.Chain {
	case "btc":
		{
			switch c.Network {
			case "main":
				{
					networkParams = &chaincfg.MainNetParams
				}
			case "test3":
				{
					networkParams = &chaincfg.TestNet3Params
				}
			default:
				{
					networkParams = &chaincfg.MainNetParams
				}
			}
		}
	}
	privateKeyBtc, _, receiveAddressBtc, err := c.GenerateAddressSegwit(networkParams)
	if err != nil {
		return "", "", err
	}
	return receiveAddressBtc, privateKeyBtc, nil
}

func (c *Client) GenerateAddressSegwit(networkParams *chaincfg.Params, privateKey ...string) (privKey, pubKey, addressSegwit string, err error) {
	var secret *btcec.PrivateKey
	if len(privateKey) == 0 {
		secret, err = btcec.NewPrivateKey()
		if err != nil {
			err = errors.Wrap(err, "c.GenerateAddressSegwit")
			return
		}
	} else {
		secret, _ = btcec.PrivKeyFromBytes([]byte(privateKey[0]))
		if secret == nil {
			err = errors.Wrap(err, "invalid secret")
			return
		}
	}

	wif, err := btcutil.NewWIF(secret, networkParams, true)
	if err != nil {
		err = errors.Wrap(err, "c.GenerateAddressSegwit")
		return
	}

	privKey = wif.String()

	witnessProg := btcutil.Hash160(wif.PrivKey.PubKey().SerializeCompressed())
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, networkParams)

	if err != nil {
		err = errors.Wrap(err, "btcutil.NewAddressWitnessPubKeyHash")
		return
	}

	addressSegwit = addressWitnessPubKeyHash.EncodeAddress()

	return
}

func (c *Client) GenerateAddressTaproot() (string, string, error) {
	var networkParams *chaincfg.Params
	switch c.Chain {
	case "btc":
		{
			switch c.Network {
			case "main":
				{
					networkParams = &chaincfg.MainNetParams
				}
			case "test3":
				{
					networkParams = &chaincfg.TestNet3Params
				}
			default:
				{
					panic(errors.New("invalid network"))
				}
			}
		}
	}
	privateKeyBtc, _, receiveAddressBtc, err := GenerateAddressTaproot(networkParams)
	if err != nil {
		return "", "", err
	}
	return receiveAddressBtc, privateKeyBtc, nil
}

func GenerateAddressTaproot(chainParam *chaincfg.Params, seed ...string) (privKey, pubKey, addressTaproot string, err error) {
	var secret *btcec.PrivateKey
	if len(seed) == 0 {
		secret, err = btcec.NewPrivateKey()
		if err != nil {
			err = errors.Wrap(err, "c.GenerateAddressTaproot generate new private key error")
			return
		}
	} else {
		hash := sha256.Sum256([]byte(seed[0]))
		secret, _ = btcec.PrivKeyFromBytes(hash[:])
		if secret == nil {
			err = errors.Wrap(err, "c.GenerateAddressTaproot invalid secret")
			return
		}
	}

	wif, err := btcutil.NewWIF(secret, chainParam, true)
	if err != nil {
		err = errors.Wrap(err, "c.GenerateAddressTaproot invalid secret")
		return
	}

	tapKey := txscript.ComputeTaprootKeyNoScript(wif.PrivKey.PubKey())
	spendTapPubKey := schnorr.SerializePubKey(tapKey)
	taprootAddress, err := btcutil.NewAddressTaproot(spendTapPubKey, chainParam)
	if err != nil {
		err = errors.Wrap(err, "btcutil.NewAddressTaproot")
		return
	}

	privKey = wif.String()
	addressTaproot = taprootAddress.String()
	pubKey = hex.EncodeToString(wif.PrivKey.PubKey().SerializeCompressed())
	return
}

func GenerateAddressTaprootFromPrivateKey(chainParam *chaincfg.Params, privateKey string) (addressTaproot string, err error) {
	wif, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		err = errors.Wrap(err, "c.GenerateAddressTaproot invalid secret")
		return
	}
	tapKey := txscript.ComputeTaprootKeyNoScript(wif.PrivKey.PubKey())
	spendTapPubKey := schnorr.SerializePubKey(tapKey)
	taprootAddress, err := btcutil.NewAddressTaproot(spendTapPubKey, chainParam)
	if err != nil {
		err = errors.Wrap(err, "btcutil.NewAddressTaproot")
		return
	}
	addressTaproot = taprootAddress.String()
	return
}

type TxRef struct {
	TxHash        string    `json:"tx_hash"`
	BlockHeight   int       `json:"block_height"`
	TxInputN      int       `json:"tx_input_n"`
	TxOutputN     int       `json:"tx_output_n"`
	Value         int       `json:"value"`
	RefBalance    int       `json:"ref_balance"`
	Spent         bool      `json:"spent"`
	Confirmations int       `json:"confirmations"`
	Confirmed     time.Time `json:"confirmed"`
	DoubleSpend   bool      `json:"double_spend"`
	// SatRanges     [][]uint64 `json:"sat_ranges"`
}

type BlockCypherWalletInfo struct {
	Address            string  `json:"address"`
	TotalReceived      int     `json:"total_received"`
	TotalSent          int     `json:"total_sent"`
	Balance            int     `json:"balance"`
	UnconfirmedBalance int     `json:"unconfirmed_balance"`
	FinalBalance       int     `json:"final_balance"`
	NTx                int     `json:"n_tx"`
	UnconfirmedNTx     int     `json:"unconfirmed_n_tx"`
	FinalNTx           int     `json:"final_n_tx"`
	Txrefs             []TxRef `json:"txrefs"`
	TxURL              string  `json:"tx_url"`
	Error              string  `json:"error"`
}

type QuickNodeUTXO_Resp struct {
	ID      int             `json:"id"`
	Result  []QuickNodeUTXO `json:"result"`
	Jsonrpc string          `json:"jsonrpc"`
}

type QuickNodeUTXO struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Value         string `json:"value"`
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
}

func (c *Client) GetBalanceFromQuickNode(address string) (*BlockCypherWalletInfo, error) {
	var respond QuickNodeUTXO_Resp
	var result BlockCypherWalletInfo
	payload := strings.NewReader(fmt.Sprintf("{\n\t\"method\": \"bb_getutxos\",\n\t\"params\": [\n\t\t\"%v\"\n\t, {\"confirmed\": true}]\n}", address))
	req, err := http.NewRequest("POST", c.QNUrl, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respond)
	if err != nil {
		log.Println(string(body))
		return nil, err
	}
	totalBalance := 0
	convertedUTXOList := []TxRef{}
	for _, utxo := range respond.Result {
		value, err := strconv.ParseUint(utxo.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		totalBalance += int(value)
		newTxReft := TxRef{
			TxHash:      utxo.Txid,
			TxOutputN:   utxo.Vout,
			Value:       int(value),
			BlockHeight: utxo.Height,
		}
		convertedUTXOList = append(convertedUTXOList, newTxReft)
	}
	result.Address = address
	result.Balance = totalBalance
	result.FinalBalance = totalBalance
	result.Txrefs = convertedUTXOList
	return &result, nil
}

func (c *Client) GetBalanceFromBlockcypler(addr string) (uint64, error) {
	client, err := c.getClient()
	if err != nil {
		return 0, err
	}
	addrInfo, err := client.GetAddrBal(addr, nil)
	if err != nil {
		return 0, err
	}
	return addrInfo.FinalBalance.Uint64(), nil
}

func (c *Client) Balance(addr string) (uint64, error) {
	switch c.Chain {
	case "btc":
		{
			switch c.Network {
			case "main":
				{
					data, err := c.GetBalanceFromQuickNode(addr)
					if err != nil {
						return 0, err
					}
					return uint64(data.Balance), nil
				}
			case "test3":
				{
					return c.GetBalanceFromBlockcypler(addr)
				}
			default:
				{
					panic(errors.New("invalid network"))
				}
			}
		}
	}
	return 0, nil
}

func (bs *Client) Transfer(secret string, from string, destination string, amount int) (string, error) {
	balance, err := bs.Balance(from)
	if err != nil {
		return "", err
	}
	if amount > 0 && int(balance) == amount {
		amount = -1
	}
	chain, err := bs.getClient()
	if err != nil {
		return "", err
	}
	wif, err := btcutil.DecodeWIF(secret)
	if err != nil {
		return "", err
	}
	pkHex := hex.EncodeToString(wif.PrivKey.Serialize())
	tx := gobcy.TempNewTX(from, destination, *big.NewInt(int64(amount)))
	tx.Preference = PreferenceHigh
	skel, err := chain.NewTX(tx, false) // gobcy.TX
	if err != nil {
		return "", err
	}
	prikHexs := []string{}
	for i := 0; i < len(skel.ToSign); i++ {
		prikHexs = append(prikHexs, pkHex)
	}
	err = skel.Sign(prikHexs)
	if err != nil {
		return "", err
	}
	// add this one with segwit address:
	for i := range skel.Signatures {
		skel.Signatures[i] = skel.Signatures[i] + "01"
	}
	skel, err = chain.SendTX(skel)
	if err != nil {
		return "", err
	}
	return skel.Trans.Hash, nil
}

func (c *Client) CheckTXHash(hash string, confirmedBlock int) (bool, error) {
	client, err := c.getClient()
	if err != nil {
		return false, err
	}
	tnx, err := client.GetTX(hash, nil)
	if err != nil {
		return false, err
	}
	if tnx.Confirmations < confirmedBlock {
		return false, errors.New("tx is not yet confirmed")
	}
	return true, nil
}
