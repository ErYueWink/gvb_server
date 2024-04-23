package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/utils/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithMsg("请求成功", c)
}
