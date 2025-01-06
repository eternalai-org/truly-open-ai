package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func CurlURL(apiURL string, method string, headers map[string]string, postData interface{}, respData interface{}) error {
	var err error
	mt := http.MethodGet
	if method != "" {
		mt = method
	}
	var bytesBuffer io.Reader
	if postData != nil {
		var bodyBytes []byte
		bodyBytes, err = json.Marshal(postData)
		if err != nil {
			return err
		}
		bytesBuffer = bytes.NewBuffer(bodyBytes)
	}
	var req *http.Request
	req, err = http.NewRequest(mt, apiURL, bytesBuffer)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	client := &http.Client{}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var respBytes []byte
		respBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = fmt.Errorf("request is bad response status code %d ( %s )", res.StatusCode, string(respBytes))
		return err
	}
	if respData != nil {
		var respBytes []byte
		respBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(respBytes, respData)
		if err != nil {
			return err
		}
	}
	return nil
}

func MakeSeoURL(title string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	prettyurl := reg.ReplaceAllString(title, "-")
	prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))
	return prettyurl
}

func ConvertImageDataURL(tokenURL string) string {
	var ipfsReplaces = []string{
		"https://cloudflare-ipfs.com/ipfs",
		"https://ipfs.fleek.co/ipfs",
		"https://nearnaut.mypinata.cloud/ipfs",
		"ipfs:/",
	}
	for _, ipfsReplace := range ipfsReplaces {
		if strings.HasPrefix(tokenURL, ipfsReplace) {
			tokenURL = strings.Replace(tokenURL, ipfsReplace, "https://ipfs.io/ipfs", -1)
		}
	}
	return tokenURL
}

func CurlURLString(apiURL string, method string, headers map[string]string, postData interface{}) (string, error) {
	var err error
	mt := http.MethodGet
	if method != "" {
		mt = method
	}
	var bytesBuffer io.Reader
	if postData != nil {
		var bodyBytes []byte
		bodyBytes, err = json.Marshal(postData)
		if err != nil {
			return "", err
		}
		bytesBuffer = bytes.NewBuffer(bodyBytes)
	}
	var req *http.Request
	req, err = http.NewRequest(mt, apiURL, bytesBuffer)
	if err != nil {
		return "", err
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	client := &http.Client{}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var respBytes []byte
		respBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		err = fmt.Errorf("request is bad response status code %d ( %s )", res.StatusCode, string(respBytes))
		return "", err
	}
	respBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(respBytes), nil
}

func CurlBase64String(apiURL string) (string, error) {
	var err error
	mt := http.MethodGet
	var req *http.Request
	req, err = http.NewRequest(mt, apiURL, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var respBytes []byte
		respBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		err = fmt.Errorf("request is bad response status code %d ( %s )", res.StatusCode, string(respBytes))
		return "", err
	}
	respBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(respBytes), nil
}

func RemoveImageURLs(text string) string {
	// Regex pattern to match media URLs with specific extensions
	mediaURLPattern := `https?://\S+\.(jpg|jpeg|png|gif|svg)`
	re := regexp.MustCompile(mediaURLPattern)
	return re.ReplaceAllString(text, "")
}

func RemoveURLs(text string) string {
	// Regex pattern to match media URLs
	mediaURLPattern := `https?://t\.co/\S+`
	re := regexp.MustCompile(mediaURLPattern)
	return re.ReplaceAllString(text, "")
}
