package miniocloud

import (
	"chai-comutils/common/utils/characterutil"
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"net/url"
	"time"
)

// MinioPresignedGetObject 为 HTTP GET 操作生成预签名 URL
// 即使存储桶是私有的，浏览器/移动客户端也可以指向此 URL 来直接下载对象
// 此预签名 URL 可以有一个关联的过期时间（以秒为单位），过期后将不再可操作
// 最长过期时间为 604800 秒（即 7 天），最短过期时间为 1 秒
// time.Second * 24 * 60 * 60 //1 day
func MinioPresignedGetObject(arg *MinioBasicParam, bucketName, objectName string, expiry time.Duration, reqParamss string) (*url.URL, error) {
	var (
		minioClient  *minio.Client
		presignedURL *url.URL
		err          error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	presignedURL, err = minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, reqParams)
	return presignedURL, err
}

// MinioPresignedPutObject 为 HTTP PUT 操作生成预签名 URL
// 浏览器/移动客户端可能会指向此 URL 以将对象直接上传到存储桶，即使它是私有的
// 此预签名 URL 可以有一个关联的过期时间（以秒为单位），过期后将不再可操作。默认有效期设置为 7 天
// time.Second * 24 * 60 * 60 //1 day
func MinioPresignedPutObject(arg *MinioBasicParam, bucketName, objectName string, day int) (string, error) {
	var (
		minioClient  *minio.Client
		presignedURL *url.URL
		err          error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return "", errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}
	expiry := time.Second * 24 * 60 * 60 * time.Duration(day)
	presignedURL, err = minioClient.PresignedPutObject(context.Background(), bucketName, objectName, expiry)
	return presignedURL.String(), err
}

func MinioPresignedPostPolicy(arg *MinioBasicParam, bucketName, objectName string, day int) (string, map[string]string, error) {
	var (
		minioClient  *minio.Client
		presignedURL *url.URL
		formData     map[string]string
		err          error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return "", nil, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	policy := minio.NewPostPolicy()
	policy.SetBucket(bucketName)
	policy.SetKey(objectName)
	policy.SetExpires(time.Now().UTC().AddDate(0, 0, day))

	presignedURL, formData, err = minioClient.PresignedPostPolicy(context.Background(), policy)
	if err != nil {
		return "", nil, errors.New(characterutil.StitchingBuilderStr("minio presigned post policy err:", err.Error()))
	}

	return presignedURL.String(), formData, nil
}
