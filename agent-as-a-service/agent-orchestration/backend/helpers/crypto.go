package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func ValidateMessageSignature(msg string, signatureHex string, signer string) error {
	msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgBytes := []byte(msg)
	msgHash := crypto.Keccak256Hash(
		msgBytes,
	)
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}
	if signature[crypto.RecoveryIDOffset] > 1 {
		signature[crypto.RecoveryIDOffset] -= 27
	}
	sigPublicKey, err := crypto.SigToPub(msgHash.Bytes(), signature)
	if err != nil {
		return err
	}
	pbkHex := crypto.PubkeyToAddress(*sigPublicKey)
	if !strings.EqualFold(pbkHex.Hex(), signer) {
		return errors.New("not valid signer")
	}
	return nil
}
