package middleware

import (
	"Blog/utils"
	"Blog/utils/errormsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,

		StandardClaims: jwt.StandardClaims{
			//注意时间格式转换
			//ExpiresAt: expireTime.Unix(),
			ExpiresAt: jwt.NewTime(float64(expireTime.Unix())),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errormsg.ERROR
	}
	return token, errormsg.SUCCSE
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errormsg.SUCCSE
	} else {
		return nil, errormsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errormsg.SUCCSE
		if tokenHeader == "" {
			code = errormsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"mssage1": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errormsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"mssage2": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errormsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"mssage3": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt.Unix() {
			c.JSON(http.StatusOK, gin.H{
				"code":   code,
				"mssage": errormsg.GetErrMsg(code),
			})
			code = errormsg.ERROR_TOKEN_RUNTIME
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
