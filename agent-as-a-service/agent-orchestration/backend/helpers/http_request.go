package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type AllowedCode struct {
	Code map[string]string
}

type RelyErrorMessage struct {
	Code    string      `json:"code"`
	Message *string     `json:"message"`
	Error   interface{} `json:"error"`
}

type RelyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAllowedCode() AllowedCode {
	ac := new(AllowedCode)
	codes := make(map[string]string)
	codes["200_ok"] = "200 OK"
	codes["201_created"] = "201 Created"
	codes["202_accepted"] = "202 Accepted"
	ac.Code = codes
	return *ac
}

func HttpRequest(fullUrl string, method string, headers map[string]string, reqBody interface{}) ([]byte, *http.Header, int, error) {
	fullUrl = strings.TrimSpace(fullUrl)
	bff := new(bytes.Buffer)
	if reqBody != nil {
		byteData, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, 0, err
		}
		bff = bytes.NewBuffer(byteData)
	}

	req, err := http.NewRequest(method, fullUrl, bff)
	if err != nil {
		return nil, nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, 0, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	isAllowed := isAllowed(res.Status)
	if !isAllowed {
		data := &RelyErrorMessage{}
		err = json.Unmarshal(body, data)
		if err != nil {
			return nil, nil, res.StatusCode, err
		}

		dataErrorString, ok := data.Error.(string)
		if ok {
			return nil, &res.Header, res.StatusCode, errors.New(dataErrorString)
		}

		dataError := &RelyError{}
		byteArray, err := json.Marshal(data.Error)
		if err != nil {
			return nil, nil, res.StatusCode, err
		}

		err = json.Unmarshal(byteArray, dataError)
		if err != nil {
			if data.Message != nil {
				return nil, &res.Header, res.StatusCode, errors.New(*data.Message)
			}
		}
		return nil, &res.Header, res.StatusCode, errors.New(dataError.Message)

	}

	if err != nil {
		return nil, nil, res.StatusCode, err
	}

	return body, &res.Header, res.StatusCode, nil
}

func HttpRequestFullResponse(fullUrl string, method string, headers map[string]string, reqBody interface{}) ([]byte, *http.Header, int, error) {
	fullUrl = strings.TrimSpace(fullUrl)
	bff := new(bytes.Buffer)
	if reqBody != nil {
		byteData, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, 0, err
		}
		bff = bytes.NewBuffer(byteData)
	}

	req, err := http.NewRequest(method, fullUrl, bff)
	if err != nil {
		return nil, nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &res.Header, res.StatusCode, err
	}
	return body, &res.Header, res.StatusCode, nil
}

func JsonRequest(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, *http.Header, int, error) {
	// headers["accept"] = "application/json"
	// headers["content-type"] = "application/json"

	return HttpRequest(fullUrl, method, headers, reqBody)
}

func isAllowed(code string) bool {
	ac := NewAllowedCode()
	code = strings.ReplaceAll(code, " ", "_")
	code = strings.ToLower(code)
	getCode, ok := ac.Code[code]
	if !ok {
		return false
	}

	_ = getCode
	return true
}
