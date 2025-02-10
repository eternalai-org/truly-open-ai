package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Environment string

const Production Environment = "production"

func IsEnvProduction(env string) bool {
	return env == string(Production)
}

func filenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func CreateFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		log.Printf("Created directory: %s", path)
	} else if err != nil {
		return err
	} else {
		log.Printf("Directory %s already exists", path)
	}
	return nil
}

func DownloadFileByUrl(link string, localPath string) (string, string, error) {
	parsedUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return "", "", err
	}

	// Get the data
	client := &http.Client{}
	req, _ := http.NewRequest("GET", link, nil)

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", errors.New("received non 200 response code")
	}

	fileName := strings.Replace(resp.Header.Get("Content-Disposition"), "attachment; filename=", "", -1)
	fileName = strings.Replace(fileName, "attachment;filename=", "", -1)
	if fileName == "" {
		filePath := parsedUrl.Path
		fileName = path.Base(filePath)
	}

	_ = CreateFolderIfNotExists(filepath.Join(localPath, filenameWithoutExtension(fileName)))

	pathFile := filepath.Join(localPath, filenameWithoutExtension(fileName), fileName)
	// Create the file
	out, err := os.Create(pathFile)
	if err != nil {
		return "", "", err
	}
	defer out.Close()
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", "", err
	}

	return pathFile, fileName, nil
}
