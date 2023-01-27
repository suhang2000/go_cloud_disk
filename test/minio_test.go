package test

import (
	"errors"
	"fmt"
	"github.com/minio/minio-go"
	"go_cloud_disk/core/define"
	"log"
	"testing"
)

func setup() (*minio.Client, error) {
	endpoint := "127.0.0.1:9000"
	accessKeyID := define.MinioId
	secretAccessKey := define.MinioKey
	useSSL := false

	// Initialize minio client object.
	client, err := minio.NewV4(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return nil, errors.New("failed to initialize")
	}
	return client, nil
}

func TestMinIOBucket(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%#v\n", client) // minioClient is now set up

	// list all buckets
	buckets, err := client.ListBuckets()
	if err != nil {
		t.Fatal(err)
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	// check if bucket exists
	found, err := client.BucketExists("bucket-cloud-disk")
	if err != nil {
		t.Fatal(err)
	}
	if found {
		println("bucket exists")
	} else {
		println("bucket not found")
	}
}

func TestUploadFile(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	n, err := client.FPutObject("bucket-cloud-disk", "cover.jpg", "./image/welcome-cover.jpg",
		minio.PutObjectOptions{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Successfully uploaded bytes: ", n)
}
