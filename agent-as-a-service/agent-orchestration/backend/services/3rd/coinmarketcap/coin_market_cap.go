package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type CoinMarketCap struct {
	serverURL string
	apiKey    string
}

type PriceConversionResponse struct {
	Status interface{}                 `json:"status"`
	Data   PriceConversionDataResponse `json:"data"`
}

type PriceConversionDataResponse struct {
	Id          int       `json:"id"`
	Symbol      string    `json:"symbol"`
	Name        string    `json:"name"`
	Amount      int       `json:"amount"`
	LastUpdated time.Time `json:"last_updated"`
	Quote       Quote     `json:"quote"`
}

type HistoricalDataCoin struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quotes []*struct {
		Timestamp string `json:"timestamp"`
		Quote     *struct {
			USD *struct {
				Price     numeric.BigFloat `json:"price"`
				Timestamp time.Time        `json:"timestamp"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"quotes"`
}

type HistoricalData struct {
	Data map[string]*HistoricalDataCoin `json:"data"`
}

type Quote struct {
	USD USD `json:"usd"`
}

type USD struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

func NewCoinMarketCap(apiKey string) *CoinMarketCap {
	apiURL := "https://pro-api.coinmarketcap.com"
	return &CoinMarketCap{
		serverURL: apiURL,
		apiKey:    apiKey,
	}
}

func (m *CoinMarketCap) generateUrl(path string) string {
	fullUrl := fmt.Sprintf("%s/%s", m.serverURL, path)
	return fullUrl
}

func (m *CoinMarketCap) request(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, int, error) {

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
	req.Header.Add("X-CMC_PRO_API_KEY", m.apiKey)

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

func (m *CoinMarketCap) GetHistoricalPrice(id string, unixEnd int64) (*big.Float, error) {
	urlQueries := url.Values{}
	urlQueries.Set("id", id)
	urlQueries.Set("time_start", fmt.Sprintf("%d", unixEnd-1))
	urlQueries.Set("time_end", fmt.Sprintf("%d", unixEnd))
	urlQueries.Set("interval", "hourly")
	path := fmt.Sprintf("v3/cryptocurrency/quotes/historical?%s", urlQueries.Encode())
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &HistoricalData{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	if resp.Data[id] == nil {
		return nil, errors.New("not found coin")
	}
	var price numeric.BigFloat
	for _, v := range resp.Data[id].Quotes {
		if v.Quote.USD.Timestamp.Unix() == unixEnd {
			price = v.Quote.USD.Price
		}
	}
	if price.Float.Cmp(big.NewFloat(0)) <= 0 {
		return nil, errors.New("not found price")
	}
	return &price.Float, nil
}

type QuotesLatestDataCoin struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  *struct {
		USD *struct {
			Price     numeric.BigFloat `json:"price"`
			MarketCap numeric.BigFloat `json:"market_cap"`
			Volume24h numeric.BigFloat `json:"volume_24h"`
		} `json:"USD"`
	} `json:"quote"`
}

type QuotesLatestData struct {
	Data map[string]*QuotesLatestDataCoin `json:"data"`
}

func (m *CoinMarketCap) GetQuotesLatest(arrayIDs []string) (map[string]*QuotesLatestDataCoin, error) {
	// urlQueries := url.Values{}
	// urlQueries.Set("id", strings.Join(arrayIDs, ","))
	path := fmt.Sprintf("v2/cryptocurrency/quotes/latest?id=%s", strings.Join(arrayIDs, ","))
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	resp := &QuotesLatestData{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type MapDataCoin struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
}

type MapData struct {
	Data []*MapDataCoin `json:"data"`
}

func (m *CoinMarketCap) GetCryptocurrencyMap() (map[string]*MapDataCoin, error) {
	path := "v1/cryptocurrency/map?sort=cmc_rank"
	fullUrl := m.generateUrl(path)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fullUrl)
	resp := &MapData{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	coinMap := map[string]*MapDataCoin{}
	for _, v := range resp.Data {
		coinMap[fmt.Sprintf("%d", v.ID)] = v
	}
	return coinMap, nil
}

type DexPairsResp struct {
	PlatformId          uint   `json:"platformId"`
	PlatformName        string `json:"platformName"`
	BaseTokenSymbol     string `json:"baseTokenSymbol"`
	QuoteTokenSymbol    string `json:"quoteTokenSymbol"`
	Liquidity           string `json:"liquidity"`
	PairContractAddress string `json:"pairContractAddress"`
	PlatFormCryptoId    string `json:"platFormCryptoId"`
	ExchangeId          uint   `json:"exchangeId"`
	PoolId              uint   `json:"poolId"`
	BaseTokenName       string `json:"baseTokenName"`
	MarketCap           string `json:"marketCap"`
	PriceUsd            string `json:"priceUsd"`
	PriceChange24h      string `json:"priceChange24h"`
	BaseToken           *struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Symbol   string `json:"symbol"`
		Decimals uint   `json:"decimals"`
	} `json:"baseToken"`
	QuoteToken *struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Symbol   string `json:"symbol"`
		Decimals uint   `json:"decimals"`
	} `json:"quoteToken"`
	Volume24h      string `json:"volume24h"`
	VolumeQuote24h string `json:"volumeQuote24h"`
}

type DexSearchResp struct {
	Data *struct {
		Total uint            `json:"total"`
		Pairs []*DexPairsResp `json:"pairs"`
	} `json:"data"`
}

func (m *CoinMarketCap) DexSearch(keyword string) ([]*DexPairsResp, error) {
	fullUrl := fmt.Sprintf(`https://api.coinmarketcap.com/dexer/v3/dexer/search/main-site?keyword=%s&all=false`, keyword)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fullUrl)
	resp := &DexSearchResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.Pairs, nil
}

type DexSpotPairsLatestDetail struct {
	ScrollId                  string `json:"scroll_id"`
	ContractAddress           string `json:"contract_address"`
	Name                      string `json:"name"`
	BaseAssetId               string `json:"base_asset_id"`
	BaseAssetUcid             string `json:"base_asset_ucid"`
	BaseAssetName             string `json:"base_asset_name"`
	BaseAssetSymbol           string `json:"base_asset_symbol"`
	BaseAssetContractAddress  string `json:"base_asset_contract_address"`
	QuoteAssetId              string `json:"quote_asset_id"`
	QuoteAssetUcid            string `json:"quote_asset_ucid"`
	QuoteAssetName            string `json:"quote_asset_name"`
	QuoteAssetSymbol          string `json:"quote_asset_symbol"`
	QuoteAssetContractAddress string `json:"quote_asset_contract_address"`
	DexId                     string `json:"dex_id"`
	DexSlug                   string `json:"dex_slug"`
	NetworkId                 string `json:"network_id"`
	NetworkSlug               string `json:"network_slug"`
	LastUpdated               string `json:"last_updated"`
	CreatedAt                 string `json:"created_at"`
	Quote                     []*struct {
		ConvertId             string  `json:"convert_id"`
		Price                 float64 `json:"price"`
		PriceByQuoteAsset     float64 `json:"price_by_quote_asset"`
		LastUpdated           string  `json:"last_updated"`
		Volume24h             float64 `json:"volume_24h"`
		PercentChangePrice1h  float64 `json:"percent_change_price_1h"`
		PercentChangePrice24h float64 `json:"percent_change_price_24h"`
		Liquidity             float64 `json:"liquidity"`
		FullyDilutedValue     float64 `json:"fully_diluted_value"`
	} `json:"quote"`
}

type DexSpotPairsLatestResp struct {
	Data *[]*DexSpotPairsLatestDetail `json:"data"`
}

func (m *CoinMarketCap) DexSpotPairsLatest(quoteAssetSymbol, networkSlug string) (*DexSpotPairsLatestResp, error) {
	fullUrl := fmt.Sprintf(`https://pro-api.coinmarketcap.com/v4/dex/spot-pairs/latest?network_slug=%s&quote_asset_symbol=%s`, networkSlug, quoteAssetSymbol)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fullUrl)
	resp := &DexSpotPairsLatestResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type DexPairsTradeLatestDetail struct {
	ScrollId                  string `json:"scroll_id"`
	ContractAddress           string `json:"contract_address"`
	Name                      string `json:"name"`
	BaseAssetId               string `json:"base_asset_id"`
	BaseAssetUcid             string `json:"base_asset_ucid"`
	BaseAssetName             string `json:"base_asset_name"`
	BaseAssetSymbol           string `json:"base_asset_symbol"`
	BaseAssetContractAddress  string `json:"base_asset_contract_address"`
	QuoteAssetId              string `json:"quote_asset_id"`
	QuoteAssetUcid            string `json:"quote_asset_ucid"`
	QuoteAssetName            string `json:"quote_asset_name"`
	QuoteAssetSymbol          string `json:"quote_asset_symbol"`
	QuoteAssetContractAddress string `json:"quote_asset_contract_address"`
	DexId                     string `json:"dex_id"`
	DexSlug                   string `json:"dex_slug"`
	NetworkId                 string `json:"network_id"`
	NetworkSlug               string `json:"network_slug"`
	LastUpdated               string `json:"last_updated"`
	CreatedAt                 string `json:"created_at"`
	Trades                    []*struct {
		Date      string `json:"date"`
		TradeType string `json:"type"`
		Quote     []*struct {
			Price             float64 `json:"price"`
			Total             float64 `json:"total"`
			ConvertId         string  `json:"convert_id"`
			PriceByQuoteAsset float64 `json:"price_by_quote_asset"`
			AmountBaseAsset   float64 `json:"amount_base_asset"`
			AmountQuoteAsset  float64 `json:"amount_quote_asset"`
		} `json:"quote"`
	} `json:"trades"`
}

type DexPairsTradeLatestResp struct {
	Data *[]*DexPairsTradeLatestDetail `json:"data"`
}

func (m *CoinMarketCap) DexPairsTradeLatest(contractAddress, networkSlug string) (*DexPairsTradeLatestResp, error) {
	fullUrl := fmt.Sprintf(`https://pro-api.coinmarketcap.com/v4/dex/pairs/trade/latest?contract_address=%s&network_slug=%s`, contractAddress, networkSlug)
	data, _, err := m.request(fullUrl, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fullUrl)
	resp := &DexPairsTradeLatestResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	client := NewCoinMarketCap("")
	// data, err := client.GetHistoricalPrice("1", time.Now().Truncate(1*time.Hour).Unix()-5*3600)
	// fmt.Println(data, err)

	data1, _ := client.GetQuotesLatest([]string{"1"})
	for _, v := range data1 {
		price, _ := v.Quote.USD.Price.Float64()
		fmt.Println(v.Symbol, price)
	}
}
