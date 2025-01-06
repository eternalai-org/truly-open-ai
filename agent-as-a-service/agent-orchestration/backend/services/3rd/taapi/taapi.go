package taapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TaApi struct {
	serverURL string
	apiKey    string
}

func NewTaApi(apiKey string) *TaApi {
	return &TaApi{
		serverURL: "https://api.taapi.io",
		apiKey:    apiKey,
	}
}

func (m *TaApi) generateUrl(path string) string {
	fullUrl := fmt.Sprintf("%s/%s", m.serverURL, path)
	return fullUrl
}

func (m *TaApi) request(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest(method, fullUrl, reqBody)
	if err != nil {
		return nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return body, res.StatusCode, nil
}

type FibonacciRetracementResp struct {
	Value          float64 `json:"value"`
	Trend          string  `json:"trend"`
	StartPrice     float64 `json:"startPrice"`
	EndPrice       float64 `json:"endPrice"`
	StartTimestamp int64   `json:"startTimestamp"`
	EndTimestamp   int64   `json:"endTimestamp"`
}

func (m *TaApi) GetFibonacciRetracement(tokenSymbol string) (*FibonacciRetracementResp, error) {
	params := url.Values{}
	params.Add("secret", m.apiKey)
	params.Add("exchange", "binance")
	params.Add("symbol", fmt.Sprintf(`%s/USDT`, tokenSymbol))
	params.Add("interval", "1d")

	path := fmt.Sprintf("fibonacciretracement?%s", params.Encode())
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &FibonacciRetracementResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *TaApi) GetRsi(tokenSymbol string) (*FibonacciRetracementResp, error) {
	params := url.Values{}
	params.Add("secret", m.apiKey)
	params.Add("exchange", "binance")
	params.Add("symbol", fmt.Sprintf(`%s/USDT`, tokenSymbol))
	params.Add("interval", "1d")

	path := fmt.Sprintf("rsi?%s", params.Encode())
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &FibonacciRetracementResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *TaApi) GetEma(tokenSymbol string) (*FibonacciRetracementResp, error) {
	params := url.Values{}
	params.Add("secret", m.apiKey)
	params.Add("exchange", "binance")
	params.Add("symbol", fmt.Sprintf(`%s/USDT`, tokenSymbol))
	params.Add("interval", "1d")
	params.Add("period", "12")

	path := fmt.Sprintf("ema?%s", params.Encode())
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &FibonacciRetracementResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *TaApi) GetSma(tokenSymbol string) (*FibonacciRetracementResp, error) {
	params := url.Values{}
	params.Add("secret", m.apiKey)
	params.Add("exchange", "binance")
	params.Add("symbol", fmt.Sprintf(`%s/USDT`, tokenSymbol))
	params.Add("interval", "1d")
	params.Add("period", "20")

	path := fmt.Sprintf("sma?%s", params.Encode())
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &FibonacciRetracementResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type BulkRequestDetailResp struct {
	Result FibonacciRetracementResp `json:"result"`
	ID     string                   `json:"id"`
}

type BulkRequestResp struct {
	Data []*BulkRequestDetailResp `json:"data"`
}

func (m *TaApi) BulkRequest(tokenSymbol string) (*BulkRequestResp, error) {
	bodyReq := map[string]interface{}{
		"secret": m.apiKey,
		"construct": map[string]interface{}{
			"exchange": "binance",
			"symbol":   fmt.Sprintf(`%s/USDT`, tokenSymbol),
			"interval": "1d",
			"indicators": []map[string]interface{}{
				{
					"id":        "fibonacciretracement",
					"indicator": "fibonacciretracement",
				},
				{
					"id":        "rsi",
					"indicator": "rsi",
				},
				{
					"id":        "sma",
					"indicator": "sma",
					"period":    "20",
				},
				{
					"id":        "ema",
					"indicator": "ema",
					"period":    "12",
				},
			},
		},
	}

	fullUrl := m.generateUrl("bulk")
	bodyBytes, _ := json.Marshal(bodyReq)
	data, _, err := m.request(fullUrl, "POST", nil, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	resp := &BulkRequestResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
