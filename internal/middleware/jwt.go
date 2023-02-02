package middleware

import (
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	//"fmt"
	//"github.com/hyperledger/gin-demo/common"
	//"github.com/hyperledger/gin-demo/response"
	//"github.com/gin-gonic/gin"
	//"net/http"
)

type Myclaims struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	jwt.StandardClaims
}

var Secret = []byte("bytedance051/")

// jwt过期时间, 按照实际环境设置
const TokenExpireDuration = time.Hour * 1

// GenToken 生成Token
func GenToken(username string, id int64) (string, error) {
	c := Myclaims{
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bytedance",
		},
	}
	// 用指定的哈希方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GenToken(claims.Username, claims.ID)
	}
	return "", errors.New("Couldn't handle this token")
}

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		//tokenStr := ctx.Request.Header.Get("Authorization")
		tokenStr := ctx.Query("token")
		//fmt.Println(tokenStr)
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Have not token",
			})
			return
		}
		claims, err := ParseToken(tokenStr)
		//fmt.Println(claims)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "ERR_AUTH_INVALID",
			})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "ERR_AUTH_EXPIRED",
			})
			return
		}
		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		ctx.Set("username", claims.Username)
		ctx.Set("id", claims.ID)
		ctx.Next()
	}
}
