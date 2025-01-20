package utils

import (
	"strings"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/pkg/encrypt"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/pkg/eth"
)

// return encryptedPrivateKeyKey, address,  error
func GenerateAddress(secretKey string) (string, string, error) {
	privateKey, _, address, err := eth.GenerateAddress()
	if err != nil {
		return "", "", err
	}

	encryptedPrivateKeyKey, err := encrypt.EncryptToString(privateKey, secretKey)
	if err != nil {
		return "", "", err
	}

	return encryptedPrivateKeyKey, strings.ToLower(address), nil
}
