package http_client

import (
	"bufio"
	"decentralized-inference/internal/types"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func RequestHttp(fullUrl string, method string, headers map[string]string, reqBody io.Reader, timeout int64) ([]byte, int, error) {

	req, err := http.NewRequest(method, fullUrl, reqBody)
	if err != nil {
		return nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	return body, res.StatusCode, nil
}

func RequestHttpWithStream(fullUrl string, method string, headers map[string]string, reqBody io.Reader, timeout int64, out chan types.StreamData) {
	var err error
	defer func() {
		if err != nil {
			out <- types.StreamData{
				Stop:     true,
				Err:      err,
				StreamID: -1,
				Data:     []byte{},
			}

			return
		}

		out <- types.StreamData{
			Stop:     true,
			Err:      nil,
			StreamID: -1,
			Data:     []byte{},
		}

		close(out)
	}()

	req, err := http.NewRequest(method, fullUrl, reqBody)
	if err != nil {
		return
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	// Read the response body in a buffered manner
	scanner := bufio.NewScanner(res.Body)
	index := 0
	for scanner.Scan() {
		streamData := scanner.Text()
		streamData = strings.ReplaceAll(streamData, "data: ", "")
		if streamData != "[DONE]" && streamData != "" {
			out <- types.StreamData{
				StreamID: index,
				Data:     []byte(streamData),
				Stop:     false,
				Err:      nil,
			}

			index++

			// Process each line of the response
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

}
