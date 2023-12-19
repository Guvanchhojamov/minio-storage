package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func main() {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	minOpts := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	}
	minioClient, err := minio.New(endpoint, minOpts)
	if err != nil {
		log.Fatal(err)
	}
	buckOpts := minio.MakeBucketOptions{ObjectLocking: false}
	err = minioClient.MakeBucket(context.Background(), "first", buckOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test sotrage minio: %v\n %v", minioClient, buckOpts)
}
