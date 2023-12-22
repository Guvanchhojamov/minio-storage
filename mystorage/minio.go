package mystorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	EndpointMinio   = "127.0.0.1:9000"
	accessKeyID     = "admin"
	secretAccessKey = "admin12345"
	useSSL          = false
)

type MyStorage struct {
	storageClient *minio.Client
}

func NewMyStorage() *MyStorage {
	return &MyStorage{}
}
func (ms *MyStorage) ConnectStorage() (*minio.Client, error) {
	minOpts := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	}
	return minio.New(EndpointMinio, minOpts)

}
