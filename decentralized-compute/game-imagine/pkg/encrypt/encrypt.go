package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

// EncryptToByte encrypts the given value using AES-GCM.
func EncryptToByte(value, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, value, nil)
	return ciphertext, nil
}

// EncryptToString encrypts the given value using AES-GCM and encodes the result in base64.
func EncryptToString(value, key string) (string, error) {
	if key == "" {
		return value, nil
	}

	keyBytes := []byte(key)
	valueBytes := []byte(value)
	ciphertext, err := EncryptToByte(valueBytes, keyBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptToByte decrypts the given ciphertext using AES-GCM.
func DecryptToByte(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce := ciphertext[:gcm.NonceSize()]
	plaintext, err := gcm.Open(nil, nonce, ciphertext[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// DecryptToString decrypts the given base64-encoded ciphertext using AES-GCM.
func DecryptToString(ciphertext, key string) (string, error) {
	if key == "" {
		return ciphertext, nil
	}

	keyBytes := []byte(key)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plaintextBytes, err := DecryptToByte(ciphertextBytes, keyBytes)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}

// EncryptToByteChaCha20Poly1305 encrypts the given value using ChaCha20-Poly1305.
func EncryptToByteChaCha20Poly1305(value, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aead.Seal(nonce, nonce, value, nil)
	return ciphertext, nil
}

// EncryptToStringChaCha20Poly1305 encrypts the given value using ChaCha20-Poly1305 and encodes the result in base64.
func EncryptToStringChaCha20Poly1305(value, key string) (string, error) {
	if key == "" {
		return value, nil
	}

	keyBytes := []byte(key)
	valueBytes := []byte(value)
	ciphertext, err := EncryptToByteChaCha20Poly1305(valueBytes, keyBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptToByteChaCha20Poly1305 decrypts the given ciphertext using ChaCha20-Poly1305.
func DecryptToByteChaCha20Poly1305(ciphertext, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aead.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce := ciphertext[:aead.NonceSize()]
	plaintext, err := aead.Open(nil, nonce, ciphertext[aead.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// DecryptToStringChaCha20Poly1305 decrypts the given base64-encoded ciphertext using ChaCha20-Poly1305.
func DecryptToStringChaCha20Poly1305(ciphertext, key string) (string, error) {
	if key == "" {
		return ciphertext, nil
	}

	keyBytes := []byte(key)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plaintextBytes, err := DecryptToByteChaCha20Poly1305(ciphertextBytes, keyBytes)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}

// GenerateAESKey generates a random AES key with the given size.
// The key size must be 16 (128 bits), 24 (192 bits), or 32 (256 bits).
func GenerateAESKey(keySize int) (string, error) {
	if keySize != 16 && keySize != 24 && keySize != 32 {
		return "", fmt.Errorf("invalid key size. Must be 16 (128 bits), 24 (192 bits), or 32 (256 bits)")
	}

	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return string(key), nil
}
