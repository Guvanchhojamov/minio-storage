package router

import (
	"github.com/minio/minio-go/v7"
)

type Features struct {
}

func (r *Router) MakeBucket(storageClient minio.Client, bucketName string) error {

}
