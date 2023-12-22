package models

type Create struct {
	BucketName string `json:"name" db:"name" binding:"required"`
}
