package secret_manager

import (
	"context"
	"crypto/cipher"
	"encoding/base64"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"golang.org/x/crypto/blowfish"
)

func GetGoogleSecretKey(ctx context.Context, name string) (string, error) {
	// Create the client.
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{Name: name}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}

// decryptToByte : value, cipherKey
func decryptToByte(value string, cipherKey string) ([]byte, error) {
	var returnMe, keyByteArr []byte
	keyByteArr = []byte(cipherKey)
	decodeB64, err1 := decodeBase64(value)
	if err1 != nil {
		return nil, err1
	}

	returnMe, err2 := blowfishDecrypt(decodeB64, keyByteArr)
	if err2 != nil {
		return nil, err2
	}
	return returnMe, nil
}

// DecryptToString : value, cipherKey
func DecryptToString(value string, cipherKey string) (string, error) {
	if cipherKey == "" {
		return value, nil
	}

	decryptedByteArr, err := decryptToByte(value, cipherKey)
	if decryptedByteArr == nil {
		return "", err
	}
	var returnMe string = string(decryptedByteArr[:])

	returnMe = strings.TrimRight(returnMe, "\x00")

	return returnMe, nil
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func blowfishDecrypt(value, key []byte) ([]byte, error) {
	dcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	div := value[:blowfish.BlockSize]
	decrypted := value[blowfish.BlockSize:]
	if len(decrypted)%blowfish.BlockSize != 0 {
		return nil, err
	}
	dcbc := cipher.NewCBCDecrypter(dcipher, div)
	dcbc.CryptBlocks(decrypted, decrypted)
	return decrypted, nil
}
