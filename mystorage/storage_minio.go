package mystorage

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
)

type StorageMinio struct {
	storageClient *minio.Client
}

func NewStorageMinio(storageClient *minio.Client) *StorageMinio {
	return &StorageMinio{storageClient: storageClient}
}

func (s *StorageMinio) CreateBucket(bucketName string) (bool, error) {

	makeOpts := minio.MakeBucketOptions{
		ObjectLocking: false,
	}
	err := s.storageClient.MakeBucket(context.Background(), bucketName, makeOpts)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *StorageMinio) UploadFile() (minio.UploadInfo, error) {

	var (
		bucketName = "images"
		objectName = "image3.png"
		filePath   = "./tmp/image3.png"
		putOpts    = minio.PutObjectOptions{ContentType: "image/png"}
	)
	return s.storageClient.FPutObject(context.Background(), bucketName, objectName, filePath, putOpts)
}

func (s *StorageMinio) DownloadFile() error {
	return errors.New("some error")
}
func (s *StorageMinio) GetBuckedFiles() error {
	return errors.New("some error")
}
func (s *StorageMinio) GetFileInfo() error {
	return errors.New("some error")
}
func (s *StorageMinio) RemoveFileFromStorage() error {
	return errors.New("some error")
}
func (s *StorageMinio) GetFileLink() error {
	return errors.New("some error")
}
