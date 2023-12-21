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
	handler := router.NewRouter(storageClient)
	srv := server.NewServer(&http.Server{})
	err = srv.Run(":8085", handler.InitRoutes())
	if err != nil {
		log.Fatal("server run error")
	}

}
