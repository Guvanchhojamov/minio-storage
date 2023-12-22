package mystorage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio-test/models"
	"net/url"
	"time"
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
		objectName = "image1.png"
		filePath   = "./tmp/image1.png"
		putOpts    = minio.PutObjectOptions{ContentType: "image/png"}
	)
	return s.storageClient.FPutObject(context.Background(), bucketName, objectName, filePath, putOpts)
}

func (s *StorageMinio) DownloadFile() error {
	var (
		bucketName = "images"
		objectName = "image3.png"
		filePath   = "./download/image3.png"
		putOpts    = minio.GetObjectOptions{}
	)
	return s.storageClient.FGetObject(context.Background(), bucketName, objectName, filePath, putOpts)
}
func (s *StorageMinio) GetBucketFiles() (interface{}, error) {
	var (
		bucketName = "images"
		obj        models.Object
		objects    []models.Object
	)
	listOpt := minio.ListObjectsOptions{
		WithMetadata: true,
	}
	objectsChannel := s.storageClient.ListObjects(context.Background(), bucketName, listOpt)
	for object := range objectsChannel {
		if object.Err != nil {
			return nil, object.Err
		}
		obj.Name = object.Key
		obj.LastModified = object.LastModified
		obj.Type = object.UserMetadata
		obj.Size = object.Size
		objects = append(objects, obj)
	}
	return objects, nil
}
func (s *StorageMinio) GetFileInfo() (minio.ObjectInfo, error) {
	var (
		bucketName = "images"
		objcetName = "image1.png"
	)
	objOpts := minio.StatObjectOptions{}

	objectInfo, err := s.storageClient.StatObject(context.Background(), bucketName, objcetName, objOpts)
	if err != nil {
		return minio.ObjectInfo{}, err
	}
	return objectInfo, err
}
func (s *StorageMinio) RemoveFileFromStorage() (bool, error) {
	var (
		bucketName = "images"
		objcetName = "image1.png"
	)
	removeOpts := minio.RemoveObjectOptions{}
	err := s.storageClient.RemoveObject(context.Background(), bucketName, objcetName, removeOpts)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (s *StorageMinio) GetFileLink() (*url.URL, error) {
	var (
		bucketName = "images"
		objectName = "image1.png"
		expires    = time.Hour * 24
	)
	urlParams := make(url.Values)
	urlParams.Set("response-content-disposition", fmt.Sprintf(`attachment; filename="%s"`, objectName))
	u, err := s.storageClient.PresignedGetObject(context.Background(), bucketName, objectName, expires, urlParams)
	if err != nil {
		return nil, err
	}
	return u, err
}
