package router

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type Router struct {
	minioClient *minio.Client
}

func NewRouter(minioClient *minio.Client) *Router {
	return &Router{minioClient: minioClient}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/create", r.craeteBucket)
	return router
}
