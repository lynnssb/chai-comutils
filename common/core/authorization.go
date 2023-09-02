/**
 * @author:       wangxuebing
 * @fileName:     authorization.go
 * @date:         2023/1/14 16:33
 * @description:
 */

package core

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

const (
	CtxAuth = "Authorization"
)
const (
	CtxAuthKeyByJwtAdminUserId  = "JwtAdminUserId"  //管理端用户ID
	CtxAuthKeyByJwtClientUserId = "JwtClientUserId" //客户端用户ID
	CtxAuthKeyByJwtIsAdmin      = "JwtIsAdmin"      //是否是管理员
)

// GeneratorJwtToken Generator Token
// Params:
//
//	val: (默认可以传:CtxAuthKeyByJwtUserId <如需其他可以自定义>)
//	secretKey:
//	iat:
//	seconds:
//	userId:
func GeneratorJwtToken(val string, secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["iss"] = "lynn"
	claims[val] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// GetAuthJwtKeyAdminUserId 根据ctx解析token中的jwtTenantId
func GetAuthJwtKeyAdminUserId(ctx context.Context) string {
	adminUserId := GetAuthJwtKeyValue[string](ctx, CtxAuthKeyByJwtAdminUserId)

	return adminUserId
}

// GetAuthJwtKeyClientUserId 根据ctx解析token中的ClientUserId
func GetAuthJwtKeyClientUserId(ctx context.Context) string {
	clientUserId := GetAuthJwtKeyValue[string](ctx, CtxAuthKeyByJwtClientUserId)

	return clientUserId
}

func GetAuthJwtKeyValue[T any](ctx context.Context, val string) T {
	value := ctx.Value(val)
	return value.(T)
}
