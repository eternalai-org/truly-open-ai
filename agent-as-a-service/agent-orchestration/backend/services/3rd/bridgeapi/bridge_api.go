package bridgeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/pkg/errors"
)

type BridgeApi struct {
	url string
}

func NewBridgeApi(url string) *BridgeApi {
	return &BridgeApi{
		url: url,
	}
}

func (m *BridgeApi) request(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, int, error) {
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
	req.Header.Add("x-api-key", "a71e3753381842")
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

type SolanaEAITxResp struct {
	ID                   string         `json:"id"`
	DeletedAt            *time.Time     `json:"deleted_at"`
	CreatedAt            *time.Time     `json:"created_at"`
	UpdatedAt            *time.Time     `json:"updated_at"`
	FromTokenAddress     string         `json:"fromTokenAddress"`
	FromNativeAddress    string         `json:"fromNativeAddress"`
	DepositTokenAddress  string         `json:"depositTokenAddress"`
	DepositNativeAddress string         `json:"depositNativeAddress"`
	TxReceivedDeposit    string         `json:"txReceivedDeposit"`
	Token                string         `json:"token"`
	Amount               numeric.BigInt `json:"amount"`
	Block                int            `json:"block"`
}

func (m *BridgeApi) GetSolanaEAITxs(fromBlock uint64) ([]*SolanaEAITxResp, error) {
	path := fmt.Sprintf("%s/api/internal/get-solana-eai-txs?fromBlock=%d", m.url, fromBlock)
	data, code, err := m.request(path, "GET", map[string]string{}, nil)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New(string(data))
	}
	var res struct {
		Status bool               `json:"status"`
		Data   []*SolanaEAITxResp `json:"data"`
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
