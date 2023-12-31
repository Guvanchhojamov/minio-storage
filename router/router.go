package router

import (
	"github.com/gin-gonic/gin"
	"minio-test/mystorage"
)

type Router struct {
	mystorage *mystorage.StorageMinio
}

func NewRouter(mystorage *mystorage.StorageMinio) *Router {
	return &Router{mystorage: mystorage}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/create", r.craeteBucket)
	router.POST("/upload", r.uploadFile)
	router.GET(":bucketname/:filename", r.downloadFile)
	router.GET("/objects", r.getFiles)
	router.GET("/info", r.getFileInfo)
	router.GET("/delete", r.deleteFile)
	router.GET("/getlink", r.getFileLink)
	return router
}
