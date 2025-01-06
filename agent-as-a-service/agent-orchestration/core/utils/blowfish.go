package utils

import (
	"bytes"
	"crypto/cipher"
	"encoding/base64"

	"golang.org/x/crypto/blowfish"
)

/*
   Reference:
       https://github.com/ipfans/golang-sample/blob/master/blowfish.go
*/

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func blowfishChecksizeAndPad(value []byte) []byte {
	modulus := len(value) % blowfish.BlockSize
	if modulus != 0 {
		padnglen := blowfish.BlockSize - modulus
		for i := 0; i < padnglen; i++ {
			value = append(value, 0)
		}
	}
	return value
}

func blowfishEncrypt(value, key []byte) ([]byte, error) {
	bcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	returnMe := make([]byte, blowfish.BlockSize+len(value))
	eiv := returnMe[:blowfish.BlockSize]
	ecbc := cipher.NewCBCEncrypter(bcipher, eiv)
	ecbc.CryptBlocks(returnMe[blowfish.BlockSize:], value)
	return returnMe, nil
}

// EncryptToByte : value, cipherKey
func EncryptToByte(value string, cipherKey string) ([]byte, error) {
	var returnMe, valueInByteArr, paddedByteArr, keyByteArr []byte
	valueInByteArr = []byte(value)
	keyByteArr = []byte(cipherKey)
	paddedByteArr = blowfishChecksizeAndPad(valueInByteArr)
	returnMe, err := blowfishEncrypt(paddedByteArr, keyByteArr)
	if err != nil {
		return nil, err
	}
	return returnMe, nil
}

// EncryptToString : value, cipherKey
func EncryptToString(value string, cipherKey string) (string, error) {
	encryptedByteArr, err := EncryptToByte(value, cipherKey)
	if err != nil {
		return "", err
	}
	returnMe := encodeBase64(encryptedByteArr)
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

// DecryptToByte : value, cipherKey
func DecryptToByte(value string, cipherKey string) ([]byte, error) {
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
	decryptedByteArr, err := DecryptToByte(value, cipherKey)
	if decryptedByteArr == nil {
		return "", err
	}
	// remove padding character
	removedNilChars := bytes.Trim(decryptedByteArr, "\x00")
	var returnMe = string(removedNilChars[:])
	return returnMe, nil
}
