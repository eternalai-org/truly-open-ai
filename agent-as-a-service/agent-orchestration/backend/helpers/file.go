package helpers

import (
	cryptorand "crypto/rand"
	"math/big"
	"os"
	"strings"
)

func WriteFileEternalTemp(body []byte) (string, error) {
	if _, err := os.Stat("/tmp/eternal-data/"); os.IsNotExist(err) {
		err := os.MkdirAll("/tmp/eternal-data/", os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	buf := make([]byte, 32)
	_, err := cryptorand.Read(buf)
	if err != nil {
		return "", err
	}
	hash := big.NewInt(0).SetBytes(buf).Text(16)
	err = os.WriteFile("/tmp/eternal-data/"+hash, body, 0644)
	if err != nil {
		return "", err
	}
	return "file://" + hash, nil
}

func ReadFileEternalTemp(hash string) ([]byte, error) {
	hash = strings.TrimPrefix(hash, "file://")
	body, err := os.ReadFile("/tmp/eternal-data/" + hash)
	if err != nil {
		return nil, err
	}
	return body, nil
}
