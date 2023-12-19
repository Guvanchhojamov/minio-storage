package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) craeteBucket(ctx *gin.Context) {
	r.minioClient.MakeBucket()
	ctx.JSON(http.StatusOK, "created")
}
