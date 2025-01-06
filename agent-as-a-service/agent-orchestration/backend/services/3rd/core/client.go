package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	BaseURL string
}

func (c *Client) buildUrl(resourcePath string) string {
	if resourcePath != "" {
		return c.BaseURL + "/" + resourcePath
	}
	return c.BaseURL
}

func (c *Client) doWithoutAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
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

func (c *Client) StoreAddress(addr string, prk string, walletType int, walletId string) (string, string, error) {
	resp := struct {
		Result struct {
			Address       string
			SecretVersion string
		}
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("wallet/update"),
		map[string]interface{}{
			"WalletId":   walletId,
			"WalletType": walletType,
			"Address":    addr,
			"PrivateKey": prk,
		},
		&resp,
	)
	if err != nil {
		return "", "", err
	}
	return resp.Result.Address, resp.Result.SecretVersion, nil
}

func (c *Client) GetAddressPrk(addr string) (string, error) {
	resp := struct {
		Result string
	}{}
	err := c.methodJSON(
		http.MethodPost,
		c.buildUrl("wallet/get-private-key"),
		map[string]interface{}{
			"Address": addr,
		},
		&resp,
	)
	if err != nil {
		return "", err
	}
	resp.Result = strings.TrimPrefix(resp.Result, "0x")
	return resp.Result, nil
}
