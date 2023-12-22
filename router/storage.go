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
func (r *Router) downloadFile(ctx *gin.Context) {
	err := r.mystorage.DownloadFile()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Download": "success",
		"Saved in": "/download",
	})
	return
}
func (r *Router) getFiles(ctx *gin.Context) {
	data, err := r.mystorage.GetBucketFiles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
	return
}
func (r *Router) getFileInfo(ctx *gin.Context) {
	data, err := r.mystorage.GetFileInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
	return
}

func (r *Router) deleteFile(ctx *gin.Context) {
	isDeleted, err := r.mystorage.RemoveFileFromStorage()
	if err != nil || !isDeleted {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Deleted": true,
	})
	return
}
func (r *Router) getFileLink(ctx *gin.Context) {
	fileUrl, err := r.mystorage.GetFileLink()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"File URL": fileUrl.String(),
	})
	return
}
