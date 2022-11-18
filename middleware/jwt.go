package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"qiublog/utils"
	"qiublog/utils/errmsg"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(id uint, username string, role int) (string, int) {
	expireTime := time.Now().Add(time.Hour * 24 * 3)
	SetClaims := MyClaims{
		id,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Qiudie",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	cRes := func(c *gin.Context, code int) {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Abort()
	}
	return func(c *gin.Context) {
		ckToken, err := c.Cookie("token")
		if err != nil {
			//认证字符串判断 !没有认证字符串
			code = errmsg.ERROR_TOKEN_EXIST
			cRes(c, code)
			return
		}
		key, tCode := CheckToken(ckToken)
		if tCode == errmsg.ERROR {
			//认证字符串判断 !内容不对
			code = errmsg.ERROR_TOKEN_WRONG
			cRes(c, code)
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			//认证字符串时间判断 !过期
			code = errmsg.ERROR_TOKEN_RUNTIME
			cRes(c, code)
			return
		}
		c.Set("id", key.Id)
		c.Next()
	}
}
