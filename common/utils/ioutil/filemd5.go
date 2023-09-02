/**
 * @author:       wangxuebing
 * @fileName:     filemd5.go
 * @date:         2023/1/16 22:43
 * @description:
 */

package ioutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// FileMD5 计算文件MD5值
// Params:
//
//	filePath:文件路径
func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
