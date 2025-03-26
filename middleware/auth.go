package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yangsai7/demo_api/config"
)

const (
	UID         = "X-UID"
	WeproOpenId = "X-Wepro-Open-Id"
	Token       = "token"
)

func SetUser(ctx *gin.Context) {

	tokenString := ctx.GetHeader(Token)
	if tokenString == "" {
		ctx.Next()
		return
	}

	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证算法是否匹配
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于解密
		return []byte(config.GlobalCfg.Wepro.JwtSign), nil
	})
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"msg":  "jwt parse error:" + err.Error(),
		})
		ctx.Abort()
		return
	}

	// 处理解析后的令牌
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Set(UID, int64(claims["uid"].(float64)))
		ctx.Set(WeproOpenId, claims["openid"].(string))
		ctx.Next()
		return
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"msg":  "Token is not valid",
		})
		ctx.Abort()
		return
	}
}
