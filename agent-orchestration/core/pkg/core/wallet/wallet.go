package wallet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/domain"
)

// Wallet ...
type Wallet struct {
	Endpoint string
}

// Init ...
func Init(e string) *Wallet {
	return &Wallet{
		Endpoint: e,
	}
}

// Generate ...
func (r *Wallet) Generate(req domain.WalletGenerateRequest) (*domain.Wallet, error) {
	apiURL := fmt.Sprintf("%s/wallet/generate", r.Endpoint)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		Result *domain.Wallet
		Error  *domain.Error
	}
	json.Unmarshal(b, &data)

	if data.Error != nil {
		fmt.Println(data.Error, data.Error.Error())
		return nil, data.Error
	}

	return data.Result, nil
}

// GetPrivateKey ...
func (r *Wallet) GetPrivateKey(req domain.WalletGetPrivateKeyRequest) (string, error) {
	apiURL := fmt.Sprintf("%s/wallet/get-private-key", r.Endpoint)

	body, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data struct {
		Result string
		Error  *domain.Error
	}
	json.Unmarshal(b, &data)

	if data.Error != nil {
		fmt.Println(data.Error, data.Error.Error())
		return "", data.Error
	}

	return data.Result, nil
}

func (r *Wallet) Check(req domain.WalletCheckRequest) error {
	apiURL := fmt.Sprintf("%s/wallet/check", r.Endpoint)

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var data struct {
		Error *domain.Error
	}
	json.Unmarshal(b, &data)

	if data.Error != nil {
		fmt.Println(data.Error, data.Error.Error())
		return data.Error
	}

	return nil
}
