package miniocloud

import (
	"chai-comutils/common/utils/characterutil"
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
)

// MinioGetObject 返回对象数据流
func MinioGetObject(arg *MinioBasicParam, bucketName, objectName string) (*minio.Object, error) {
	var (
		minioClient *minio.Client
		object      *minio.Object
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	object, err = minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio get object err:", err.Error()))
	}

	return object, nil
}

// MinioFGetObject 下载对象并将其保存为本地文件系统中的文件
func MinioFGetObject(arg *MinioBasicParam, bucketName, objectName string, localFilePath string) error {
	var (
		minioClient *minio.Client
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}
	err = minioClient.FGetObject(context.Background(), bucketName, objectName, localFilePath, minio.GetObjectOptions{})
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio f_get object err:", err.Error()))
	}
	return nil
}
