package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type TagCreateRequest struct {
	Title string `json:"title" binding:"required" msg:"标签标题"`
}

// TagCreateView 发布标签
// @Tags 标签管理
// @Summary 发布标签
// @Description 发布标签
// @Param data body TagRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	var tagModel models.TagModel
	count := global.DB.Take(&tagModel, "title = ?", cr.Title).RowsAffected
	if count != 0 {
		res.FailWithMsg(fmt.Sprintf("%s标签已存在", cr.Title), c)
		return
	}
	// 创建标签
	err = global.DB.Create(&models.TagModel{Title: cr.Title}).Error
	if err != nil {
		res.FailWithMsg("创建标签失败", c)
		return
	}
	res.OKWithMsg("创建标签成功", c)
}
