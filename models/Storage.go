package models

import (
	"github.com/minio/minio-go/v7"
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
