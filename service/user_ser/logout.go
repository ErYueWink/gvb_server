package user_ser

import (
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwt"
	"time"
)

// Logout 用户退出登录
func (UserService) Logout(claims *jwt.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
