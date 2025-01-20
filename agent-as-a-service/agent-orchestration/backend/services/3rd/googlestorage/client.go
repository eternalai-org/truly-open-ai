package googlestorage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"google.golang.org/api/option"
)

type Client struct {
	gsClient *storage.BucketHandle
}

func InitClient(credentialsFile string, bucketName string) *Client {
	opt := option.WithCredentialsFile(credentialsFile)
	ctx := context.Background()
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return &Client{
			gsClient: nil,
		}
	}
	return &Client{
		gsClient: client.Bucket(bucketName),
	}
}

func (c *Client) UploadPublicDataFile(folder string, fileName string, fileBytes []byte) (string, error) {
	fullPathImg := fmt.Sprintf("%s/%s", folder, fileName)
	if c.gsClient == nil {
		return "", errs.NewError(errs.ErrBadRequest)
	}

	w := c.gsClient.Object(fullPathImg).NewWriter(context.Background())
	if _, err := io.Copy(w, bytes.NewReader(fileBytes)); err != nil {
		return "", errs.NewError(err)
	}
	if err := w.Close(); err != nil {
		return "", errs.NewError(err)
	}
	return fullPathImg, nil
}

func (c *Client) UploadPublicMultipartFile(folder string, fileName string, file multipart.File, handle *multipart.FileHeader) (string, error) {
	fullPathImg := fmt.Sprintf("%s/%s", folder, fileName)
	if c.gsClient == nil {
		return "", errs.NewError(errs.ErrBadRequest)
	}

	w := c.gsClient.Object(fullPathImg).NewWriter(context.Background())
	w.ContentType = handle.Header.Get("Content-Type")
	if _, err := io.Copy(w, file); err != nil {
		return "", errs.NewError(err)
	}
	if err := w.Close(); err != nil {
		return "", errs.NewError(err)
	}
	return fullPathImg, nil
}

func (c *Client) UploadPublicDataBase64(folder string, fileName string, stringBase64 string) (string, error) {
	fullPathImg := fmt.Sprintf("%s/%s", folder, fileName)
	if c.gsClient == nil {
		return "", errs.NewError(errs.ErrBadRequest)
	}

	w := c.gsClient.Object(fullPathImg).NewWriter(context.Background())
	decoded, err := base64.StdEncoding.DecodeString(string([]byte(stringBase64)))
	if err != nil {
		fmt.Println(err)
	}

	if _, err := io.Copy(w, bytes.NewReader(decoded)); err != nil {
		return "", errs.NewError(err)
	}
	if err := w.Close(); err != nil {
		return "", errs.NewError(err)
	}
	return fullPathImg, nil
}
