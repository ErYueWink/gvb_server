package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
	"time"
)

// GenToken 生成token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires))), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,                                         // 签发者
			IssuedAt:  jwt.At(time.Now()),                                               // 签发时间
		},
	}
	// 根据签名类型生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 根据密钥进行编码，返回token
	return token.SignedString(MySecret)
}
