package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/Dstack-TEE/dstack/sdk/go/tappd"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	var tappdClientProvider *tappd.TappdClient
	if os.Getenv("TEE_NODE") == "PRODUCTION" {
		tappdClientProvider = tappd.NewTappdClient()
	} else {
		tappdClientProvider = tappd.NewTappdClient(
			tappd.WithEndpoint(os.Getenv("TEE_URL")),
			tappd.WithLogger(slog.Default()),
		)
	}

	ctx := context.Background()
	deriveKeyResp, err := tappdClientProvider.DeriveKey(ctx, "/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deriveKeyResp) // &{-----BEGIN PRIVATE KEY--- ...

	router := gin.Default()

	router.GET("/api/get-data", func(c *gin.Context) {
		id := c.Query("id")
		post, err := callExternalAPI(tappdClientProvider, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, post)
	})

	router.Run(":8080")
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func callExternalAPI(client *tappd.TappdClient, id string) (*Post, error) {
	// https://jsonplaceholder.typicode.com/posts/1
	response, err := http.Get(os.Getenv("TEST_URL") + id)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return nil, err
	}

	hash := sha256.New()
	hash.Write(body)
	checksum := hash.Sum(nil)
	tdxQuoteResp, err := client.TdxQuote(context.Background(), checksum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tdxQuoteResp) // &{0x0000000000000000000 ...

	rtmrs, err := tdxQuoteResp.ReplayRTMRs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rtmrs) // map[0:00000000000000000 ...

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
