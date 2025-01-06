package coingecko

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CoinGeckoAPI struct {
	serverURL string
}

func NewCoinGeckoAPI() *CoinGeckoAPI {
	serverURL := "https://api.coingecko.com/api/v3/coins"
	return &CoinGeckoAPI{
		serverURL: serverURL,
	}
}

type SolanaTokenInfo struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"price"`
}

type PriceData struct {
	Timestamp time.Time `json:"timestamp"`
	Price     float64   `json:"price"`
}

func (m *CoinGeckoAPI) GetSolanaTokenInfo(contract string) (*SolanaTokenInfo, error) {
	url := fmt.Sprintf("%s/solana/contract/%s", m.serverURL, contract)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", err)
	}

	id, _ := result["id"].(string)
	name, _ := result["name"].(string)
	symbol, _ := result["symbol"].(string)
	marketData, ok := result["market_data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Market data not found")
	}
	currentPrice, ok := marketData["current_price"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Current price not found")
	}
	price, _ := currentPrice["usd"].(float64)

	tokenInfo := &SolanaTokenInfo{
		ID:           id,
		Name:         name,
		Symbol:       symbol,
		CurrentPrice: price,
	}

	return tokenInfo, nil
}
func (m *CoinGeckoAPI) GetCoinMarketChart(coinID string, currency string) ([]PriceData, error) {
	url := fmt.Sprintf("%s/%s/market_chart?vs_currency=%s&days=1", m.serverURL, coinID, currency)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", err)
	}
	prices := result["prices"].([]interface{})
	var priceDataArray []PriceData

	for _, priceData := range prices {
		pricePair := priceData.([]interface{})
		timestamp := int64(pricePair[0].(float64))
		price := pricePair[1].(float64)

		timeStamp := time.Unix(timestamp/1000, 0)

		priceDataArray = append(priceDataArray, PriceData{
			Timestamp: timeStamp,
			Price:     price,
		})
	}

	return priceDataArray, nil
}

// func main() {
// 	client := NewCoinGeckoAPI()
// 	info, _ := client.GetSolanaTokenInfo("2KgAN8nLAU74wjiyKi85m4ZT6Z9MtqrUTGfse8Xapump")
// 	b, _ := json.Marshal(info)
// 	fmt.Println(string(b))

// 	chart, _ := client.GetCoinMarketChart(info.ID, "USD")
// 	a, _ := json.Marshal(chart)
// 	fmt.Println(string(a))
// }
