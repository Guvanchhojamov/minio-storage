package router

import (
	"errors"
	"fmt"
	"minio-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	tempFolder = "temp"
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
	var uploadInput models.UploadFile
	err := ctx.ShouldBind(&uploadInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	contentType := fileHeader.Header.Get("Content-Type")
	bucketName := ctx.PostForm("bucket_name")
	filePath := fmt.Sprintf("./%s/%s", tempFolder, fileHeader.Filename)
	err = ctx.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	putInfo, err := r.mystorage.UploadFile(fileHeader.Filename, filePath, bucketName, contentType)
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
	bucketName := ctx.Param("bucketname")
	filename := ctx.Param("filename")
	if bucketName == "" || filename == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("not valid request"))
		return
	}
	err := r.mystorage.DownloadFile(bucketName, filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, models.DownloadResponse{Download: true, Saved: "/downloads"})
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
