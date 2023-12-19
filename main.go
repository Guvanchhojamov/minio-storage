package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
)

func main() {
	endpoint := "127.0.0.1:9000"
	useSSl := false
	minioClient, err := minio.New(endpoint, &minio.Options{Secure: useSSl})
	if err != nil {
		log.Fatal(err)
	}
	buckOpts := minio.MakeBucketOptions{ObjectLocking: false}
	err = minioClient.MakeBucket(context.Background(), "first", buckOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test sotrage minio: %v\n %b", minioClient)
}
