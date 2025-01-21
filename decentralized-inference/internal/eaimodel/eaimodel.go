package eaimodel

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"math/big"
	"os"
)

type StorageType string

const (
	LightHouseStorageType StorageType = "lighthouse-filecoin"
	EaiChainStorageType   StorageType = "eai-chain"
)

type TaskResult struct {
	ResultURI string      `json:"result_uri"`
	Storage   StorageType `json:"storage"`
	Data      []byte      `json:"data"`
}

type ModelInfoContract struct {
	ModelID   *big.Int
	ModelAddr string
	OwnerAddr string

	Metadata ModelMetadata
}

type ModelMetadata struct {
	Version         uint64       `json:"version"`
	ModelName       string       `json:"model_name"`
	ModelType       ModelType    `json:"model_type"`
	ModelURL        string       `json:"model_url"`
	ModelFileHash   string       `json:"model_file_hash"`
	MinHardwareTier HardwareTier `json:"min_hardware"`

	VerifierURL      string `json:"verifier_url"`
	VerifierFileHash string `json:"verifier_file_hash"`
}

type ModelType string

const (
	ModelTypeText  ModelType = "text"
	ModelTypeImage ModelType = "image"
)

type HardwareTier int64

const (
	HardwareTier_1 HardwareTier = 1
	HardwareTier_2 HardwareTier = 2
	HardwareTier_3 HardwareTier = 3
)

func CheckModelFileHash(modelHash string, filePath string) (bool, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer f.Close()

	buf := make([]byte, 1024*1024)
	h := sha256.New()

	for {
		bytesRead, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				return false, err
			}
			break
		}
		h.Write(buf[:bytesRead])
	}

	log.Printf("checksum: %s\n", hex.EncodeToString(h.Sum(nil)))

	if hex.EncodeToString(h.Sum(nil)) != modelHash {
		return false, nil
	}

	return true, nil
}
