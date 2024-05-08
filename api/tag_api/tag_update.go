package tag_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest  true "查询参数"
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/tags/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagCreateRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var tagModel models.TagModel
	count := global.DB.Take(&tagModel, id).RowsAffected
	if count == 0 {
		res.FailWithMsg(fmt.Sprintf("ID为%v的标签不存在", id), c)
		return
	}
	// 标签存在结构体转map修改标签
	maps := structs.Map(cr)
	err = global.DB.Model(&tagModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改标签失败", c)
		return
	}
	res.OKWithMsg("修改标签成功", c)
}
