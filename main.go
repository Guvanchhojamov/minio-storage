package main

import (
	"log"
	"minio-test/mystorage"
	"minio-test/router"
	"minio-test/server"
	"net/http"
)

func main() {
	storage := mystorage.NewMyStorage()
	storageClient, err := storage.ConnectStorage()
	if err != nil {
		log.Fatal(err)
	}
	myStorageMinio := mystorage.NewStorageMinio(storageClient)
	handler := router.NewRouter(myStorageMinio)
	srv := server.NewServer(&http.Server{})

	err = srv.Run(":8085", handler.InitRoutes())
	if err != nil {
		log.Fatal("server run error")
	}

}
