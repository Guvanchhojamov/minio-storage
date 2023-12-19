package mystorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpointMinio   = "play.min.io"
	accessKeyID     = "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL          = true
)

type MyStorage struct {
}

func NewMyStorage() *MyStorage {
	return &MyStorage{}
}

func (ms *MyStorage) ConnectStorage() (*minio.Client, error) {
	minOpts := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	}
	return minio.New(endpointMinio, minOpts)

}
