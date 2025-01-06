package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	tokenLength                      uint = 32
	letters                               = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	verificationTokenExpiredDuration      = 24 * time.Hour
	referralLetters                       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GenerateReferenceCode : number, order id
func GenerateReferenceCode(n int, ID uint) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) + strconv.Itoa(int(ID))
}

// GenerateVerificationToken ...
func GenerateVerificationToken(n uint) string {
	length := tokenLength
	if n != 0 {
		length = n
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// GenerateVerificationCode ...
func GenerateVerificationCode() int {
	return random(1000, 9999)
}

// GenerateReferralCode ...
func GenerateReferralCode() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = referralLetters[rand.Int63()%int64(len(referralLetters))]
	}
	return string(b)
}

// GenerateReferralNumber ...
func GenerateReferralNumber() string {
	n := rand.Int() % 1000
	return fmt.Sprintf("%d", n)
}

// GenerateAPIToken ...
func GenerateAPIToken() string {
	b := make([]byte, 36)
	for i := range b {
		b[i] = referralLetters[rand.Int63()%int64(len(referralLetters))]
	}
	return string(b)
}

// GenerateUsername ...
func GenerateUsername(username string) string {
	nstr := strconv.Itoa(random(0, 10000))
	return fmt.Sprintf("%s%s", username, nstr)
}

// IsValidEmail : email
func IsValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

// GenerateEthereumWallet : ...
// func GenerateEthereumWallet() (string, string, error) {
// 	key, err := crypto.GenerateKey()
// 	if err != nil {
// 		return "", "", err
// 	}
// 	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
// 	privateKey := hex.EncodeToString(key.D.Bytes())
// 	return address, privateKey, nil
// }

// // GenerateBitcoinWallet : ...
// func GenerateBitcoinWallet(isTestNet bool) (string, string, error) {
// 	networkParams := &chaincfg.MainNetParams

// 	if isTestNet {
// 		networkParams = &chaincfg.TestNet3Params
// 	}

// 	secret, err := btcec.NewPrivateKey(btcec.S256())
// 	if err != nil {
// 		return "", "", err
// 	}

// 	wif, err := btcutil.NewWIF(secret, networkParams, true)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	address, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), networkParams)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return address.EncodeAddress(), wif.String(), nil
// }

// Index ...
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// SpaceMap : string
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// CheckStringIsNumber ...
func CheckStringIsNumber(s string) bool {
	_, err := strconv.ParseFloat(SpaceMap(s), 64)
	return err == nil
}

// CheckValidAccountBankNumber ...
func CheckValidAccountBankNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	rs, err := regexp.MatchString("^[a-zA-Z0-9]*$", s)
	if err != nil {
		return false
	}
	return rs
}

// SlackHook ...
func SlackHook(slackURL string, text string) error {
	bodyRequest, err := json.Marshal(map[string]interface{}{
		"text": text,
	})
	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(bodyRequest))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}

func ParseDBURL(dbURL string) string {
	var respURL string
	respMap := map[string]string{}
	strArr := strings.Split(dbURL, ";;;")
	for _, str := range strArr {
		idx := strings.Index(str, "=")
		if idx > 0 {
			respMap[str[:idx]] = str[idx+1:]
		}
	}
	respURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", respMap["username"], respMap["password"], respMap["host"], respMap["port"], respMap["database"], respMap["charset"], respMap["parseTime"], respMap["loc"])
	return respURL
}

func ParseRefID(reference string) uint {
	var refID uint
	refIDs := strings.Split(reference, "_")
	if len(refIDs) >= 2 {
		refIDTmp, _ := strconv.ParseUint(refIDs[1], 10, 64)
		refID = uint(refIDTmp)
	} else if len(refIDs) == 1 {
		refIDTmp, _ := strconv.ParseUint(refIDs[0], 10, 64)
		refID = uint(refIDTmp)
	}
	return refID
}

func GenerateMD5(v string) string {
	data := []byte(v)
	return fmt.Sprintf("%x", md5.Sum(data))
}
