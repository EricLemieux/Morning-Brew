package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

const EnvVarGoogleApplicationCredentials = "GOOGLE_APPLICATION_CREDENTIALS"

type GoogleCloudStorage struct {
	BucketId string
}

func NewGoogleCloudStorage(bucketId string) (GoogleCloudStorage, error) {
	value, exists := os.LookupEnv(EnvVarGoogleApplicationCredentials)
	if !exists || len(value) == 0 {
		return GoogleCloudStorage{}, fmt.Errorf("google cloud key not found")
	}

	return GoogleCloudStorage{
		BucketId: bucketId,
	}, nil
}

func (store GoogleCloudStorage) Upload(file []byte) (string, error) {
	ctx := context.Background()

	// TODO: this should be the byte array, probably, or whatever we get back from the audio api
	f, err := os.Open("notes.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: better error
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := client.Bucket(store.BucketId).Object("test-file-name").NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	// TODO: get the public url

	return "", nil
}
