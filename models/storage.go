package models

import (
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"time"
)

type Create struct {
	BucketName string `json:"name" db:"name" binding:"required"`
}
type Object struct {
	Name         string          `json:"name"`
	Size         int64           `json:"size"`
	LastModified time.Time       `json:"last_modified"`
	Type         minio.StringMap `json:"type"`
}
type DownloadResponse struct {
	Download bool   `json:"download"`
	Saved    string `json:"saved"`
}

type UploadFile struct {
	BucketName string                `form:"bucket_name" binding:"required"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}
