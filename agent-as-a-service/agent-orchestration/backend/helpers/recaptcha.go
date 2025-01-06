package helpers

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func ValidateRecaptcha(secret string, response string) (bool, error) {
	apiURL := "https://www.google.com/recaptcha/api/siteverify"
	form := url.Values{}
	form.Add("secret", secret)
	form.Add("response", response)
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	fmt.Println(req.Header)
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("read body error", err.Error())
		return false, err
	}
	if res.StatusCode > 300 {
		log.Println("read body error", string(body))
		return false, fmt.Errorf("%s: %s request error", res.Status, string(body))
	}
	respMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &respMap); err != nil {
		return false, errors.New("Unmarshal error")
	}
	success, ok := respMap["success"]
	if !ok {
		return false, errors.New("Bad response")
	}
	return success.(bool), nil
}
