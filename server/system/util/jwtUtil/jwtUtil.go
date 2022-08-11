/*
 * @Date: 2022-02-16 10:40:08
 * @LastEditors: 邓春宜
 * @Desc:
 * @LastEditTime: 2022-02-23 13:22:09
 * @FilePath: \gin-antd\server\app\util\jwtUtil\jwtUtil.go
 */
package jwtUtil

import (
	"errors"
	"time"

	"chunDoc/system/core/config"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	JwtSecret []byte
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID          uint
	UserName    string
	LoginName   string
	AuthorityId string
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.Instance().Jwt.JwtSecret),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: config.Instance().Jwt.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + config.Instance().Jwt.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    config.Instance().Jwt.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtSecret)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.JwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
