package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
)

// JwtAuth 路由中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			res.FailWithMsg("token不存在", c)
			c.Abort()
			return
		}
		// 解析token
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMsg("解析token失败", c)
			c.Abort()
			return
		}
		// 判断用户是否退出登录
		if redis_ser.CheckLogout(token) {
			res.FailWithMsg("用户退出登录", c)
			c.Abort()
			return
		}
		// 将载荷信息保存到cookie
		c.Set("claims", claims)
	}
}
