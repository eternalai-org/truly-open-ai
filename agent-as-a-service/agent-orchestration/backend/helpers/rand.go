package helpers

import (
	cryptorand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func RandomStringWithLength(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"abcdefghijklmnopqrstuvwxyz" +
			"0123456789",
	)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// return hex.EncodeToString(bytes), nil
	return fmt.Sprintf("0x%s", hex.EncodeToString(bytes)), nil
}

func GenShareCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"0123456789",
	)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func RandomReferralCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"abcdefghijklmnopqrstuvwxyz" +
			"0123456789",
	)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func RandomBigInt(numBytes int) *big.Int {
	buf := make([]byte, numBytes)
	_, err := cryptorand.Read(buf)
	if err != nil {
		panic(err)
	}
	return new(big.Int).SetBytes((buf))
}

func RandomNonceNumber(numberOfDigits int) (string, error) {
	resp := ""
	n := 0
	for n < numberOfDigits {
		randomNumber := rand.Intn(9)
		if (n == 0 && randomNumber != 0) || n > 0 {
			resp += strconv.Itoa(randomNumber)
			n += 1
		}
	}
	return resp, nil
}

func Sha256ToNonce(s string) *big.Int {
	h := sha256.New()
	h.Write([]byte(s))
	buf := h.Sum(nil)
	return new(big.Int).SetBytes((buf[:32]))
}

func RandHash() string {
	buf := make([]byte, 32)
	_, err := cryptorand.Read(buf)
	if err != nil {
		panic(err)
	}
	return common.BytesToHash(buf).Hex()
}

func RandInArray(arry []string) string {
	rand.Seed(time.Now().UnixNano())    // seed or it will be set to 1
	randomIndex := rand.Intn(len(arry)) // generate a random int in the range 0 to 9
	pick := arry[randomIndex]
	return pick
}
