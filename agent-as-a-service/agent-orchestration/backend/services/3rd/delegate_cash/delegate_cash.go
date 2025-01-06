package delegate_cash

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type DelegateCashAPIService struct {
	Url    string
	APIKey string
}

func NewDelegateCashAPIService(Url, APIKey string) *DelegateCashAPIService {
	return &DelegateCashAPIService{
		//Url: "https://api.delegate.xyz",
		//APIKey: "eternala-bf8c-46ed-ba49-dfaf776d1e3f",
		Url:    Url,
		APIKey: APIKey,
	}
}

// V1
func (s *DelegateCashAPIService) CheckDelegateForWalletV1(delegateAddr string, vaultAddr string) (bool, error) {
	url := fmt.Sprintf(s.Url+"/registry/v1/check/all?delegate=%s&vault=%s&contract=%s", delegateAddr, vaultAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	req.Header.Set("X-API-KEY", s.APIKey)
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}

func (s *DelegateCashAPIService) CheckDelegateForContractV1(delegateAddr string, vaultAddr string, contractAddr string) (bool, error) {
	checkWallet, _ := s.CheckDelegateForWalletV1(delegateAddr, vaultAddr)
	if checkWallet {
		return true, nil
	}
	url := fmt.Sprintf(s.Url+"/registry/v1/check/contract?delegate=%s&vault=%s&contract=%s", delegateAddr, vaultAddr, contractAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	req.Header.Set("X-API-KEY", s.APIKey)
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}

func (s *DelegateCashAPIService) CheckDelegateForTokenERC721V1(delegateAddr string, vaultAddr string, contractAddr string, tokenId string, chainID int) (bool, error) {
	checkContract, err := s.CheckDelegateForContractV1(delegateAddr, vaultAddr, contractAddr)
	if checkContract {
		return true, nil
	}

	//ref: https://docs.delegate.xyz/technical-documentation/rest-api/v1#returns-true-if-the-address-is-delegated-to-act-on-your-behalf-for-a-specific-token-the-tokens-contr
	url := fmt.Sprintf(s.Url+"/registry/v1/check/token?delegate=%s&vault=%s&contract=%s&tokenId=%s&chainId=%d", delegateAddr, vaultAddr, contractAddr, tokenId, chainID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	req.Header.Set("X-API-KEY", s.APIKey)
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}

// V2
func (s *DelegateCashAPIService) CheckDelegateForWalletV2(delegateAddr string, vaultAddr string) (bool, error) {
	url := fmt.Sprintf(s.Url+"/registry/v2/check/all?to=%s&from=%s&contract=%s", delegateAddr, vaultAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	req.Header.Set("X-API-KEY", s.APIKey)
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}

func (s *DelegateCashAPIService) CheckDelegateForContractV2(delegateAddr string, vaultAddr string, contractAddr string) (bool, error) {
	checkWallet, _ := s.CheckDelegateForWalletV2(delegateAddr, vaultAddr)
	if checkWallet {
		return true, nil
	}
	url := fmt.Sprintf(s.Url+"/registry/v2/check/contract?to=%s&from=%s&contract=%s", delegateAddr, vaultAddr, contractAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	req.Header.Set("X-API-KEY", s.APIKey)
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}

func (s *DelegateCashAPIService) CheckDelegateForTokenERC721V2(delegateAddr string, vaultAddr string, contractAddr string, tokenId string, chainID int) (bool, error) {
	checkContract, _ := s.CheckDelegateForContractV2(delegateAddr, vaultAddr, contractAddr)
	if checkContract {
		return true, nil
	}
	//ref: https://docs.delegate.xyz/technical-documentation/rest-api/v2#returns-true-if-delegate-is-granted-to-act-on-froms-behalf-for-entire-wallet-that-contract-or-that-s
	url := fmt.Sprintf(s.Url+"/registry/v2/check/erc721?to=%s&from=%s&contract=%s&tokenId=%s&chainId=%d", delegateAddr, vaultAddr, contractAddr, tokenId, chainID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	req.Header.Set("X-API-KEY", s.APIKey)
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}
