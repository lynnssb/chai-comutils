package miniocloud

import (
	"chai-comutils/common/utils/characterutil"
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/tags"
)

// MinioMakeBucket 创建一个新存储桶
func MinioMakeBucket(arg *MinioBasicParam, bucketName, region string) error {
	var (
		minioClient *minio.Client
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
		Region:        region,
		ObjectLocking: true,
	})
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("create bucket err:", err.Error()))
	}
	return nil
}

// MinioListBuckets 列出所有存储桶
func MinioListBuckets(arg *MinioBasicParam) ([]minio.BucketInfo, error) {
	var (
		minioClient *minio.Client
		buckets     []minio.BucketInfo
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	buckets, err = minioClient.ListBuckets(context.Background())
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio list_buckets err:", err.Error()))
	}
	return buckets, nil
}

// MinioBucketExists 检查存储桶是否存在
func MinioBucketExists(arg *MinioBasicParam, bucketName string) (bool, error) {
	var (
		minioClient *minio.Client
		found       bool
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return false, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}
	found, err = minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return false, errors.New(characterutil.StitchingBuilderStr("minio bucket_exists err:", err.Error()))
	}
	return found, nil
}

// MinioRemoveBucket 移除一个桶 (桶必须是空的才能成功移除)
func MinioRemoveBucket(arg *MinioBasicParam, bucketName string) error {
	var (
		minioClient *minio.Client
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}
	err = minioClient.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("remove bucket err:", err.Error()))
	}
	return nil
}

// MinioListObjects 列出存储桶中的对象
func MinioListObjects(arg *MinioBasicParam, bucketName string) ([]minio.ObjectInfo, error) {
	var (
		minioClient *minio.Client
		objects     []minio.ObjectInfo
		err         error
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    "",
		Recursive: true,
	})
	for obj := range objectCh {
		objects = append(objects, minio.ObjectInfo{
			ETag:         obj.ETag,
			Key:          obj.Key,
			LastModified: obj.LastModified,
			Size:         obj.Size,
			ContentType:  obj.ContentType,
		})
	}
	return objects, nil
}

// MinioSetBucketTagging 为存储桶设置标签
func MinioSetBucketTagging(arg *MinioBasicParam, bucketName string, v map[string]string) error {
	var (
		minioClient *minio.Client
		_tags       *tags.Tags
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}

	_tags, err = tags.NewTags(v, false)
	if err != nil {
		return err
	}
	err = minioClient.SetBucketTagging(context.Background(), bucketName, _tags)
	if err != nil {
		return err
	}
	return nil
}

// MinioGetBucketTagging 获取桶的标签
func MinioGetBucketTagging(arg *MinioBasicParam, bucketName string) (*tags.Tags, error) {
	var (
		minioClient *minio.Client
		_tags       *tags.Tags
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}
	_tags, err = minioClient.GetBucketTagging(context.Background(), bucketName)
	if err != nil {
		return nil, errors.New(characterutil.StitchingBuilderStr("minio get_bucket_tag err:", err.Error()))
	}
	return _tags, nil
}

// MinioRemoveBucketTagging 删除存储桶上的所有标签
func MinioRemoveBucketTagging(arg *MinioBasicParam, bucketName string) error {
	var (
		minioClient *minio.Client
		err         error
	)
	minioClient, err = InitMinioClient(arg.Endpoint, arg.AccessKeyId, arg.SecretAccessKey, arg.UseSSL)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio init err:", err.Error()))
	}
	err = minioClient.RemoveBucketTagging(context.Background(), bucketName)
	if err != nil {
		return errors.New(characterutil.StitchingBuilderStr("minio remove_bucket_tag err:", err.Error()))
	}
	return nil
}
