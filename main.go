package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"minio-test/mystorage"
)

func main() {
	storage := mystorage.NewMyStorage()
	storageClient, err := storage.ConnectStorage()
	if err != nil {
		log.Fatal(err)
	}
	buckOpts := minio.MakeBucketOptions{ObjectLocking: false}
	err = storageClient.MakeBucket(context.Background(), "first", buckOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("minio storage as :9000 console :9090")
}
