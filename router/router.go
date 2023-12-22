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
	return router
}
