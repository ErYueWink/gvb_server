package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/res"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/tags [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.TagModel]}
func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBind(&cr); err != nil {
		global.Log.Error(err)
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 查询标签列表
	list, count, err := common.CommList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询标签列表失败", c)
		return
	}
	res.OKWithList(list, count, c)

}
