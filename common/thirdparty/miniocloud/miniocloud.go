package miniocloud

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// InitMinioClient 初始化Minio客户端
func InitMinioClient(endpoint, accessKeyId, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}
