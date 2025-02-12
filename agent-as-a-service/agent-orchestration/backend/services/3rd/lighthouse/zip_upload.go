package lighthouse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
)

var (
	ZipChunkSize = 200 // MB
	BASH_EXEC    = getBashExecutable()
	THREADS      = runtime.NumCPU() - 1
)

type FileDetail struct {
	Name  string `json:"name"`
	Hash  string `json:"hash"`
	Index int    `json:"index"`
}

type FileInLightHouse struct {
	Name       string        `json:"name"`
	IsInserted bool          `json:"is_inserted"`
	IsPart     bool          `json:"is_part"`
	CountPart  int           `json:"count_part"`
	Files      []*FileDetail `json:"files"`
}

func getBashExecutable() string {
	bashPath, err := exec.LookPath("bash")
	if err != nil {
		fmt.Println("Bash not found. Please install bash")
		return ""
	}
	return bashPath
}

func executeCommand(fileCmd string) ([]byte, error) {
	commandId := strconv.FormatInt(time.Now().UnixMicro(), 10)
	fileLog := fmt.Sprintf("/tmp/log_%v.txt", commandId)
	execCmd := fmt.Sprintf("%v %v  2>&1 | /usr/bin/tee %v", BASH_EXEC, fileCmd, fileLog)
	fileExec := fmt.Sprintf("/tmp/bash_%v.sh", commandId)

	if err := os.WriteFile(fileExec, []byte(execCmd), 0o644); err != nil {
		return nil, err
	}

	command := exec.Command(BASH_EXEC, fileExec)
	out, err := command.Output()
	if err != nil {
		return out, err
	}
	return os.ReadFile(fileLog)
}

func getScriptZipFile(fileFolder string, baseDir string) (string, error) {
	tmpPath := path.Join("/tmp", fileFolder)
	filePath := fmt.Sprintf("%s/zip-file-%v.sh", tmpPath, fileFolder)
	if err := removeFileIfExists(filePath); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	commands := []string{
		fmt.Sprintf("cd %v", baseDir),
		fmt.Sprintf("rm -Rf %v.zip.part-*", fileFolder),
		fmt.Sprintf("tar -cf - %v | pigz --best -p %v | split -b %vM - %v.zip.part-", fileFolder, THREADS, ZipChunkSize, fileFolder),
	}

	for _, cmd := range commands {
		if _, err := file.WriteString(cmd + " \n "); err != nil {
			return "", fmt.Errorf("error writing to file: %v", err)
		}
	}

	return filePath, nil
}

func removeFileIfExists(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("error removing file: %v", err)
		}
	}
	return nil
}

func getScriptUnZipFile(fileFolder string, parentDir string) (string, error) {
	model := fmt.Sprintf("unzip-%v.sh", fileFolder)
	filePath := filepath.Join("/tmp", fileFolder, model)
	if err := removeFileIfExists(filePath); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file:%v", err)
	}
	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("cd %v \n ", parentDir)); err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}

	_, err = file.WriteString(
		fmt.Sprintf("cat %v.zip.part-* | pigz -p %v -d | tar -xf -", fileFolder, 2),
	)
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}
	return filePath, nil
}

func getListZipFile(fileFolder string, parentDir string) ([]string, error) {
	tmpPath := path.Join("/tmp", fileFolder)
	filePath := fmt.Sprintf("%s/list-zip-file-%v.sh", tmpPath, fileFolder)
	if err := removeFileIfExists(filePath); err != nil {
		return nil, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating file:%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("rm %s/list_file_%v.txt \n", tmpPath, fileFolder))
	if err != nil {
		return nil, fmt.Errorf("error write file:%v", err)
	}

	_, err = file.WriteString(fmt.Sprintf("cd %v \n", parentDir))
	if err != nil {
		return nil, fmt.Errorf("error write file:%v", err)
	}

	_, err = file.WriteString(fmt.Sprintf("ls %v.zip.part-* > %s/list_file_%v.txt \n", fileFolder, tmpPath, fileFolder))
	if err != nil {
		return nil, fmt.Errorf("error write file:%v", err)
	}

	output, err := executeCommand(fmt.Sprintf("%s/list-zip-file-%v.sh ", tmpPath, fileFolder))
	if err != nil {
		return nil, fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	file, err = os.Open(fmt.Sprintf("%s/list_file_%v.txt", tmpPath, fileFolder))
	if err != nil {
		return nil, fmt.Errorf("error opening file:%v", err)
	}

	scanner := bufio.NewScanner(file)
	listFile := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		listFile = append(listFile, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file:,%v", err)
	}
	return listFile, nil
}

func uploadListZipFileToLightHouse(fileFolder string, baseDir string, apiKey string) (*FileInLightHouse, error) {
	listFile, err := getListZipFile(fileFolder, baseDir)
	if err != nil {
		return nil, err
	}

	if len(listFile) == 0 {
		return nil, fmt.Errorf("no files pattern %v.zip.part-*  found in folder %v", fileFolder, baseDir)
	}

	result := &FileInLightHouse{
		Name:      fileFolder,
		CountPart: len(listFile),
		IsPart:    true,
		Files:     make([]*FileDetail, 0),
	}

	for i, file := range listFile {
		log.Println("Start upload model: ", fileFolder, "chunk: ", i, "file: ", file)
		for j := 0; j < 10; i++ {
			filePath := fmt.Sprintf("%v/%v", baseDir, file)
			cid, err := UploadFile(apiKey, file, filePath)
			if err != nil {
				log.Println("Error when upload model ", fileFolder, "retry", j, "chunk", i, "file", file, "err", err)
				time.Sleep(2 * time.Minute)
				continue
			} else {
				log.Println("Finish upload model: ", fileFolder, "chunk: ", i, "file: ", file, "==> hash: ", cid)
				result.Files = append(result.Files, &FileDetail{Name: file, Hash: cid, Index: i + 1})
				break
			}
		}
	}
	return result, nil
}

func uploadFileResultToLightHouse(info *FileInLightHouse, apiKey string) (string, error) {
	data, _ := json.Marshal(info)
	return UploadData(apiKey, info.Name, data)
}

func getFileResultFromLightHouse(hash string) (*FileInLightHouse, error) {
	data, _, err := DownloadDataSimple(hash)
	if err != nil {
		return nil, err
	}

	result := &FileInLightHouse{}
	if err = json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return result, nil
}

func downloadZipFileFromLightHouse(info *FileInLightHouse, baseDir string) error {
	for _, file := range info.Files {
		log.Println(
			"Start download: ",
			"file: ", file.Name,
			"hash: ", file.Hash,
			"path: ", baseDir,
		)
		for {
			filePath := filepath.Join(baseDir, file.Name)
			if err := DownloadToFile(file.Hash, filePath); err != nil {
				log.Println(
					"Error when try down file from light house",
					"file", file.Name,
					"hash", file.Hash,
					"err", err.Error(),
				)
				time.Sleep(2 * time.Minute)
				continue
			}
			break
		}
	}
	log.Println("Success download all zip file:", "file name", info.Name)
	return nil
}

func DownloadFileFromLightHouse(hash string, hfDir string) error {
	info, err := getFileResultFromLightHouse(hash)
	if err != nil {
		return fmt.Errorf("error when get model info from light house hash : %v err :%v ", hash, err)
	}

	if err = downloadZipFileFromLightHouse(info, hfDir); err != nil {
		return fmt.Errorf("error when download zip chunk file:%v ", err)
	}

	scriptFile, err := getScriptUnZipFile(info.Name, hfDir)
	if err != nil {
		return fmt.Errorf("error when get unzip script file:%v ", err)
	}
	log.Println("Start unzip list files")

	output, err := executeCommand(scriptFile)
	if err != nil {
		return fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	log.Println("Success unzip model ", info.Name)
	unzipFolder := filepath.Join(hfDir, info.Name)
	files, err := os.ReadDir(filepath.Join(hfDir, info.Name))
	if err != nil {
		return fmt.Errorf("error when read dir:%v , err:%v", unzipFolder, err.Error())
	}

	for _, file := range files {
		fmt.Printf("%s/%s\n", info.Name, file.Name())
	}
	return nil
}

func ZipAndUploadFileInMultiplePartsToLightHouse(fileFolder string, baseDir string, apiKey string) (string, error) {
	tmpPath := path.Join("/tmp", fileFolder)
	_ = utils.CreateFolderIfNotExists(tmpPath)
	scriptFile, err := getScriptZipFile(fileFolder, baseDir)
	if err != nil {
		return "", fmt.Errorf("error when get script zip file:%v ", err)
	}

	log.Println("Start compress file")
	output, err := executeCommand(scriptFile)
	if err != nil {
		return "", fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	log.Println("Finish compress file . Start upload file")
	result, err := uploadListZipFileToLightHouse(fileFolder, baseDir, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload list zip file to light house :%v ", err)
	}

	hash, err := uploadFileResultToLightHouse(result, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload model result to light house :%v ", err)
	}
	return hash, nil
}

func UploadFileInMultiplePartsToLightHouse(fileFolder string, baseDir string, apiKey string) (string, error) {
	log.Println("Start upload model")
	tmpPath := path.Join("/tmp", fileFolder)
	_ = utils.CreateFolderIfNotExists(tmpPath)
	result, err := uploadListZipFileToLightHouse(fileFolder, baseDir, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload list zip file to light house :%v ", err)
	}

	hash, err := uploadFileResultToLightHouse(result, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload model result to light house :%v ", err)
	}
	return hash, nil
}

func fileSizeInMB(filePath string) (float64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err // Return the error from os.Stat
	}

	fileSize := fileInfo.Size()                     // Size in bytes
	fileSizeMB := float64(fileSize) / (1024 * 1024) // Convert to MB

	return fileSizeMB, nil
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	buffer := make([]byte, fileSize)
	_, err = io.ReadFull(file, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func ZipAndUploadFileInMultiplePartsToLightHouseByUrl(url string, baseDir string, apiKey string) (*FileInLightHouse, error) {
	_ = utils.CreateFolderIfNotExists(baseDir)
	fullFilePath, fileName, err := utils.DownloadFileByUrl(url, baseDir)
	if err != nil {
		return nil, fmt.Errorf("error when download file:%v ", err)
	}

	fileSize, err := fileSizeInMB(fullFilePath)
	if err != nil {
		return nil, fmt.Errorf("error when get file size:%v ", err)
	}
	fileFolder := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	if fileSize < float64(ZipChunkSize) {
		bytes, err := readFile(fullFilePath)
		if err != nil {
			return nil, fmt.Errorf("error when read file:%v ", err)
		}

		hash, err := UploadDataWithRetry(apiKey, fileName, bytes)
		if err != nil {
			return nil, fmt.Errorf("error when upload file:%v ", err)
		}
		return &FileInLightHouse{
			Name:   fileName,
			IsPart: false,
			Files: []*FileDetail{
				{Name: fileName, Hash: hash, Index: 1},
			},
		}, nil
	}

	tmpPath := path.Join("/tmp", fileFolder)
	_ = utils.CreateFolderIfNotExists(tmpPath)
	scriptFile, err := getScriptZipFile(fileFolder, baseDir)
	if err != nil {
		return nil, fmt.Errorf("error when get script zip file:%v ", err)
	}

	log.Println("Start compress file")
	output, err := executeCommand(scriptFile)
	if err != nil {
		return nil, fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	log.Println("Finish compress file . Start upload file")
	result, err := uploadListZipFileToLightHouse(fileFolder, baseDir, apiKey)
	if err != nil {
		return nil, fmt.Errorf("error when upload list zip file to light house :%v ", err)
	}

	return result, nil
}
