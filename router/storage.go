package router

import (
	"fmt"
	"minio-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) craeteBucket(ctx *gin.Context) {
	var input models.Create
	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	isCreated, err := r.mystorage.CreateBucket(input.BucketName)
	if err != nil || !isCreated {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("created Bucket - %s", input.BucketName))
}
func (r *Router) uploadFile(ctx *gin.Context) {
	putInfo, err := r.mystorage.UploadFile()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Upload Info": putInfo,
	})
	return
}
