/**
 * @author:       wangxuebing
 * @fileName:     miniostorage.go
 * @date:         2023/5/24 0:30
 * @description:  Minio存储相关API接口
 */

package miniostorage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func createMinIOClient(endPoint, accessKeyId, secretAccessKey, location string, useSSL bool) (*minio.Client, error) {
	mc, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
		Region: location,
	})
	if err != nil {
		return nil, err
	}

	return mc, nil
}

/**
 * 文件上传
 * @param endPoint:         MinIO服务地址
 * @param accessKeyId:      MinIO服务访问ID
 * @param secretAccessKey:  MinIO服务访问密钥
 * @param location:         MinIO服务区域
 * @param useSSL:           是否使用SSL
 * @param bucketName:       存储桶名称
 * @param objectName:       存储对象名称
 * @param filePath:         本地文件路径
 * @return UploadInfo
 * @return error
 */
func UploadFile(endPoint, accessKeyId, secretAccessKey, location string, useSSL bool, bucketName, objectName, filePath string) (*minio.UploadInfo, error) {
	mc, err := createMinIOClient(endPoint, accessKeyId, secretAccessKey, location, useSSL)
	if err != nil {
		return nil, err
	}

	if err := mc.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
		return nil, err
	}
	exists, errBucketExists := mc.BucketExists(context.Background(), bucketName)
	if errBucketExists != nil && !exists {
		return nil, errBucketExists
	}

	fileInfo, err := mc.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: ""})
	if err != nil {
		return nil, err
	}

	return &fileInfo, nil
}

/**
 * 删除文件
 * @param endPoint:         MinIO服务地址
 * @param accessKeyId:      MinIO服务访问ID
 * @param secretAccessKey:  MinIO服务访问密钥
 * @param location:         MinIO服务区域
 * @param useSSL:           是否使用SSL
 * @param bucketName:       存储桶名称
 * @param objectName:       存储对象名称
 * @return error
 */
func RemoveFile(endPoint, accessKeyId, secretAccessKey, location string, useSSL bool, bucketName, objectName string) error {
	mc, err := createMinIOClient(endPoint, accessKeyId, secretAccessKey, location, useSSL)
	if err != nil {
		return err
	}

	err = mc.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        "",
	})

	return err
}
