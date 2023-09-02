/**
 * @author:       wangxuebing
 * @fileName:     qiniu.go
 * @date:         2023/5/20 1:48
 * @description:  七牛云存储相关API接口
 */

package qiniustorage

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

/**
 * 简单上传凭证
 * @param accessKey
 * @param secretKey
 * @param bucket
 * @return token
 */
func GetQiNiuUploadToken(accessKey, secretKey, bucket string) string {
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //在不指定上传凭证的有效时间情况下，默认有效期为1个小时
	return putPolicy.UploadToken(mac)
}
