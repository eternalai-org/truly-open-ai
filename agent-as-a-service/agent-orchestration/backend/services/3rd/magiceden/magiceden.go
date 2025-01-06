package magiceden

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
)

type MagicEdenService struct {
	url     string
	apiKey  string //b5400021ebc9489a9fe8a4544f663b53
	headers map[string]string
}

func NewMagicedenService() *MagicEdenService {
	obj := &MagicEdenService{
		url:     "https://api-mainnet.magiceden.io/v2", //os.Getenv("MAGICEDEN_URL"),
		headers: make(map[string]string),
	}
	//obj.headers["Authorization"] = "Bearer " + os.Getenv("MAGICEDEN_API_TOKEN")
	obj.headers["Accept"] = "*/*"
	obj.headers["User-Agent"] = "PostmanRuntime/7.37.3"
	obj.headers["Cookie"] = "__cf_bm=hL5OTJxF7noZcx.misHDUHHwPb71OMEH4Q_vf_400nA-1733464224-1.0.1.1-ubLBHYlzscNudJVTQQpzGKbBZ0oVQj.hgrwKpY.fuFiPvcoBpn8INX0Esfb_s9J9oxA3dLBIAvL9YlSDk9amgQ; _cfuvid=oq3NYPdYmReJ8dyHHPYNeLXeHTKXnzcURY2qIH1.g1A-1733464224460-0.0.1.1-604800000"
	return obj
}

func (s *MagicEdenService) GetInscriptionInfo(inscriptionId string) (*MagicedenInscriptionInfo, error) {
	url := fmt.Sprintf(s.url+"/ord/btc/tokens/%s", inscriptionId)
	respBytes, _, httpStatus, err := helpers.HttpRequest(url, "GET", s.headers, nil)
	if err != nil {
		return nil, err
	}
	if httpStatus != 200 {
		return nil, errors.New(fmt.Sprintf("status %d", httpStatus))
	}
	resp := &MagicedenInscriptionInfo{}
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// used for get inscriptions
type Inscription struct {
	Cohort                        string  `json:"cohort"`
	Name                          string  `json:"name"`
	CollectionSymbol              string  `json:"collectionSymbol"`
	CollectionId                  string  `json:"collectionId"`
	Vol                           float64 `json:"vol"`
	TotalVol                      float64 `json:"totalVol"`
	TotalTxns                     int     `json:"totalTxns"`
	VolPctChg                     float64 `json:"volPctChg,omitempty"`
	Txns                          int     `json:"txns"`
	TxnsPctChg                    float64 `json:"txnsPctChg,omitempty"`
	Fp                            float64 `json:"fp"`
	FpPctChg                      float64 `json:"fpPctChg,omitempty"`
	FpListingPrice                float64 `json:"fpListingPrice"`
	FpListingCurrency             string  `json:"fpListingCurrency"`
	HighestGlobalOfferBidCurrency string  `json:"highestGlobalOfferBidCurrency"`
	MarketCap                     float64 `json:"marketCap"`
	TotalSupply                   int     `json:"totalSupply"`
	ListedCount                   int     `json:"listedCount"`
	OwnerCount                    int     `json:"ownerCount"`
	UniqueOwnerRatio              float64 `json:"uniqueOwnerRatio"`
	Image                         string  `json:"image"`
	IsCompressed                  bool    `json:"isCompressed"`
	HasInscriptions               bool    `json:"hasInscriptions"`
	Currency                      string  `json:"currency"`
	Pending                       int     `json:"pending"`
	CurrencyUsdRate               float64 `json:"currencyUsdRate"`
	MarketCapUsd                  float64 `json:"marketCapUsd"`
	FpSparkLinePath               string  `json:"fpSparkLinePath"`
	Description                   string  `json:"description"`
}

func (s *MagicEdenService) GetInscriptionHardCode() ([]Inscription, error) {
	bytes := []byte(INSC_COLLECTIONS)

	resp := []Inscription{}
	err := json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *MagicEdenService) GetCollectionInfo(collection string) (*CollectionInfo, error) {
	url := fmt.Sprintf("https://api-mainnet.magiceden.io/v2/ord/btc/collections/%s", collection)
	respBytes, _, httpStatus, err := helpers.HttpRequest(url, "GET", s.headers, nil)
	if httpStatus != 200 {
		return nil, errors.New(fmt.Sprintf("status %d", httpStatus))
	}
	if err != nil {
		return nil, err
	}
	resp := &CollectionInfo{}
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
