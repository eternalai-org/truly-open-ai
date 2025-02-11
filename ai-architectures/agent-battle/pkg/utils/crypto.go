package utils

import (
	"strings"

	"agent-battle/pkg/encrypt"
	"agent-battle/pkg/eth"
)

// return encryptedPrivateKeyKey, address,  error
func GenerateAddress(secretKey string) (string, string, error) {
	var err error
	privateKey := ""
	encryptedPrivateKeyKey := ""
	address := ""
	privateKey, _, address, err = eth.GenerateAddress()
	if err != nil {
		return "", "", err
	}

	encryptedPrivateKeyKey, err = encrypt.EncryptToString(privateKey, secretKey)
	if err != nil {
		return "", "", err
	}

	return encryptedPrivateKeyKey, strings.ToLower(address), nil
}

// BeatifyWalletAddress beautify wallet address
// Currently, it just converts the address to lowercase
func BeatifyWalletAddress(address string) string {
	return strings.ToLower(address)
}