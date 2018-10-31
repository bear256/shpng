package azsb

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type Blobs struct {
	azblob.ServiceURL
}

func NewBlobs(accountName, accountKey string) *Blobs {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal(err)
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", accountName))
	serviceURL := azblob.NewServiceURL(*u, p)
	return &Blobs{serviceURL}
}

func (blobs *Blobs) Read(containerName, blobName string) ([]byte, error) {
	ctx := context.Background()
	containerURL := blobs.NewContainerURL(containerName)
	blobURL := containerURL.NewBlobURL(blobName)
	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	if err != nil {
		return nil, err
	}
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})
	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)
	if err != nil {
		return nil, err
	}
	return downloadedData.Bytes(), err
}

func (blobs *Blobs) Write(containerName, blobName string, data []byte) error {
	ctx := context.Background()
	containerURL := blobs.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)
	_, err := blobURL.Upload(ctx, bytes.NewReader(data), azblob.BlobHTTPHeaders{}, azblob.Metadata{}, azblob.BlobAccessConditions{})
	return err
}
