package dexscreener

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DexScreenerAPI struct {
	BaseURL string
}

func NewDexScreenerAPI() *DexScreenerAPI {
	return &DexScreenerAPI{
		BaseURL: "https://api.dexscreener.com",
	}
}

func (m DexScreenerAPI) request(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, int, error) {
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

type PairsDetailResp struct {
	NetworkID   uint64 `json:"network_id"`
	ChainId     string `json:"chainId"`
	Description string `json:"description"`
	DexId       string `json:"dexId"`
	Url         string `json:"url"`
	PairAddress string `json:"pairAddress"`
	PriceNative string `json:"priceNative"`
	PriceUsd    string `json:"priceUsd"`
	Volume      *struct {
		H24 float64 `json:"h24"`
		H6  float64 `json:"h6"`
		H1  float64 `json:"h1"`
		M5  float64 `json:"m5"`
	} `json:"volume"`
	PriceChange *struct {
		H24 float64 `json:"h24"`
		H6  float64 `json:"h6"`
		H1  float64 `json:"h1"`
		M5  float64 `json:"m5"`
	} `json:"priceChange"`
	Liquidity *struct {
		Usd float64 `json:"usd"`
	} `json:"liquidity"`
	QuoteToken *struct {
		Symbol  string `json:"symbol"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"quoteToken"`
	BaseToken *struct {
		Symbol  string `json:"symbol"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"baseToken"`
	Fdv       uint64 `json:"fdv"`
	MarketCap uint64 `json:"marketCap"`
	Info      *struct {
		ImageUrl  string `json:"imageUrl"`
		Header    string `json:"header"`
		OpenGraph string `json:"openGraph"`
		Websites  []struct {
			Label string `json:"label"`
			Url   string `json:"url"`
		} `json:"websites"`
		Socials []struct {
			Type string `json:"type"`
			Url  string `json:"url"`
		} `json:"socials"`
	} `json:"info"`
}

type PairsResp struct {
	Pairs []*PairsDetailResp `json:"pairs"`
}

func (m DexScreenerAPI) SearchPairs(tokenContractAddress string) (*PairsDetailResp, error) {
	fullUrl := fmt.Sprintf(`%s/latest/dex/tokens/%s`, m.BaseURL, tokenContractAddress)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fullUrl)
	resp := &PairsResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	pair := &PairsDetailResp{}
	if len(resp.Pairs) > 0 {
		liquidity := resp.Pairs[0].Liquidity.Usd
		pair = resp.Pairs[0]
		for _, item := range resp.Pairs {
			if item.Liquidity != nil && item.Liquidity.Usd > liquidity {
				liquidity = item.Liquidity.Usd
				pair = item
			}
		}
	}
	return pair, nil
}
