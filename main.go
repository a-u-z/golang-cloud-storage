package main

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func getCloudStorageByFileName(ctx context.Context, fileName string) (byteData []byte, err error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}

	r, err := client.Bucket("bucketName").Object(fileName).NewReader(ctx)
	if err != nil {
		return
	}

	return ioutil.ReadAll(r)
}
