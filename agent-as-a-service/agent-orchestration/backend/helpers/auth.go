package helpers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

type AuthToken struct {
	Address   string `json:"address"`
	Exp       int64  `json:"exp"`
	SessionID string `json:"session_id"`
}

func DecryptAndVerifyAuthToken(authToken, serverKey string) (*AuthToken, error) {
	encrypted, signed := SplitAuthToken(authToken)
	if signed == "" {
		return nil, fmt.Errorf("invalid token")
	}

	masterWallet, err := crypto.HexToECDSA(serverKey)
	if err != nil {
		return nil, err
	}

	publicKey := masterWallet.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	serverAddress := crypto.PubkeyToAddress(*publicKeyECDSA).String()

	err = VerifyAuthToken(encrypted, signed, serverAddress)
	if err != nil {
		return nil, err
	}

	decrypted, err := DecryptAuthToken(encrypted, serverKey)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now().Unix()
	if decrypted.Exp < currentTime && decrypted.Exp != 0 {
		return nil, fmt.Errorf("token expired")
	}

	return decrypted, nil
}

func SplitAuthToken(authToken string) (string, string) {
	parts := strings.Split(authToken, ".")
	if len(parts) != 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

func VerifyAuthToken(authTokenStr, signature, address string) error {
	dataHash := crypto.Keccak256Hash([]byte(authTokenStr))
	err := VerifySig(dataHash.Bytes(), signature, address)
	if err != nil {
		return err
	}
	return nil
}

func DecryptAuthToken(authTokenStr, serverKey string) (*AuthToken, error) {
	var authToken AuthToken
	err := Decrypt(authTokenStr, serverKey, &authToken)
	if err != nil {
		return nil, err
	}

	return &authToken, nil
}

func VerifySig(dataHash []byte, sig64, address string) error {
	signature, err := base64.StdEncoding.DecodeString(sig64)
	if err != nil {
		return err
	}

	if signature[crypto.RecoveryIDOffset] == 27 || signature[crypto.RecoveryIDOffset] == 28 {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	sigPub, err := crypto.SigToPub(dataHash, signature)
	if err != nil {
		return errors.New("invalid signature: " + err.Error())
	}

	sigAddress := crypto.PubkeyToAddress(*sigPub).String()

	if !strings.EqualFold(sigAddress, address) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

func Decrypt(dataStr, privKey string, result interface{}) error {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return err
	}
	encryptedBytes, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		return err
	}

	decryptedBytes, err := ecies.ImportECDSA(masterWallet).Decrypt(encryptedBytes, nil, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(decryptedBytes, result)
	if err != nil {
		return err
	}
	return nil
}

func VerifySignature(fromAddress, signatureHex, message string) error {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}

	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	messageHash := accounts.TextHash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return err
	}

	if common.HexToAddress(fromAddress) != crypto.PubkeyToAddress(*pubKey) {
		return fmt.Errorf("failed to verify signature")
	}

	return nil
}

func EncryptAndSignAuthToken(authToken AuthToken, serverKey string) (string, error) {
	encrypted, err := EncryptAuthToken(authToken, serverKey)
	if err != nil {
		return "", err
	}

	signed, err := SignAuthToken(encrypted, serverKey)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%s.%s", encrypted, signed)
	return result, nil
}

func EncryptAuthToken(authToken AuthToken, serverKey string) (string, error) {
	encrypted, err := Encrypt(authToken, serverKey)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func Encrypt(data interface{}, privKey string) (string, error) {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	publicKey := masterWallet.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	pubkey := ecies.ImportECDSAPublic(publicKeyECDSA)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	ct, err := ecies.Encrypt(rand.Reader, pubkey, dataBytes, nil, nil)
	if err != nil {
		return "", err
	}

	base64CT := base64.StdEncoding.EncodeToString(ct)
	return base64CT, nil
}

func SignAuthToken(authTokenStr, privateKey string) (string, error) {
	signature, err := Sign(authTokenStr, privateKey)
	if err != nil {
		return "", err
	}

	return signature, nil
}

func Sign(dataStr, privKey string) (string, error) {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	dataHash := crypto.Keccak256Hash([]byte(dataStr))
	signature, err := crypto.Sign(dataHash.Bytes(), masterWallet)
	if err != nil {
		return "", err
	}

	signatureStr := base64.StdEncoding.EncodeToString(signature)
	return signatureStr, nil
}
