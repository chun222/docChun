/*
 * @Date: 2022-02-16 10:07:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-03-02 17:03:05
 * @FilePath: \gin-antd\server\app\middleware\jwt.go
 */

package middleware

import (
	"time"

	"chunDoc/system/core/config"
	"chunDoc/system/core/response"
	"chunDoc/system/util/jwtUtil"
	"chunDoc/system/util/str"
	"github.com/gin-gonic/gin"
)

func JwtMid() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := str.After(c.Request.Header.Get("Authorization"), "Bearer ")

		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "error token", c)
			c.Abort()
			return
		}

		j := jwtUtil.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwtUtil.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "TokenExpired", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		//即将过期的生成新token
		if claims.ExpiresAt-time.Now().Unix() < config.Instance().Jwt.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + config.Instance().Jwt.ExpiresTime
			newToken, e := j.CreateToken(*claims)
			if newToken != "" {
				c.Header("Authorization", "Bearer "+newToken)
			} else {
				response.FailWithDetailed(gin.H{"reload": true}, e.Error(), c)
				c.Abort()
				return
			}

		}
		c.Set("claims", claims)

		c.Next()
	}
}
