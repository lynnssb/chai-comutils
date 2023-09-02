/**
 * @author:       wangxuebing
 * @fileName:     passwordhandle.go
 * @date:         2023/1/14 17:04
 * @description:
 */

package encryptionutil

import (
	"bytes"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword bcrypt根据密码和盐值生成密码
func GeneratePassword(password, salt string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(password)
	buf.WriteString(salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(buf.String()), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Created Password Error")
	}
	return string(hash), err
}

// ComparePassword 根据密文和密码+盐值验证密码是否正确
// returns:
//
//	true:  密码正确
//	false: 密码错误
func ComparePassword(sourcePassword, password, salt string) bool {
	var buf bytes.Buffer
	buf.WriteString(password)
	buf.WriteString(salt)
	err := bcrypt.CompareHashAndPassword([]byte(sourcePassword), buf.Bytes())
	if err != nil {
		return false
	}
	return true
}
