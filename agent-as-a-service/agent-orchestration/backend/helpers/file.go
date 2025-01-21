package helpers

import "os"

func WriteFileTemp(body []byte) (string, error) {
	if _, err := os.Stat("/tmp/data/"); os.IsNotExist(err) {
		err := os.MkdirAll("/tmp/data/", os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	hash := RandomReferralCode(64)
	err := os.WriteFile("/tmp/data/"+hash, body, 0644)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func ReadFileTemp(hash string) ([]byte, error) {
	body, err := os.ReadFile("/tmp/data/" + hash)
	if err != nil {
		return nil, err
	}
	return body, nil
}
