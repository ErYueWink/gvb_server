package article_api

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/jwt"
	"gvb_server/utils/res"
	"math/rand"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string      `json:"title" structs:"title" binding:"required" msg:"请输入文章标题"`                // 文章标题
	Abstract string      `json:"abstract" structs:"abstract"`                                           // 文章简介
	Content  string      `json:"content,omit(list)" structs:"content" binding:"required" msg:"请输入文章内容"` // 文章内容
	Category string      `json:"category" structs:"category"`                                           // 文章分类
	Source   string      `json:"source" structs:"source"`                                               // 文章来源
	Link     string      `json:"link" structs:"link"`                                                   // 原文链接
	BannerID uint        `json:"banner_id" structs:"banner_id"`                                         // 文章封面id
	Tags     ctype.Array `json:"tags" structs:"tags"`                                                   // 文章标签
}

// ArticleCreateView 发布文章
// @Tags 文章管理
// @Summary 发布文章
// @Description 发布文章
// @Param data body ArticleRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCreateView(c *gin.Context) {
	// 参数绑定
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 获取Cookie中的载荷信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 获取用户ID、用户昵称
	userId := claims.UserID
	userNickName := claims.NickName
	// 处理content 防止xss攻击
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	// 判断有没有script标签
	if len(nodes) > 0 {
		// 移除script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		// 获取html文本
		html, _ := doc.Html()
		// 将html文本转换为markdown格式
		mdContent, _ := converter.ConvertString(html)
		cr.Content = mdContent
	}
	// 如果用户没有输入简介则根据内容进行补全
	if cr.Abstract == "" {
		abs := []rune(cr.Content)
		if len(abs) < 100 {
			cr.Abstract = string(abs[:])
		} else {
			cr.Abstract = string(abs[:100])
		}
	}
	// 如果不传bannerID就从后台随机选取一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMsg("没有图片数据", c)
			return
		}
		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}
	// 根据图片ID查询图片路径
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("图片ID为%d的图片不存在", cr.BannerID), c)
		return
	}
	// 根据用户ID查询用户头像
	var userAvatar string
	err = global.DB.Model(&models.UserModel{}).Where("id = ?", userId).Select("avatar").Scan(&userAvatar).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("用户ID为%d的用户不存在", userId), c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		Keyword:      cr.Title,
		UserID:       userId,
		UserNickName: userNickName,
		UserAvatar:   userAvatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}
	// 在文章入库之前应该先判断文章是否已经存在
	if article.IsExistData() {
		res.FailWithMsg("文章已存在", c)
		return
	}
	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加文章失败", c)
		return
	}
	res.OKWithMsg("添加文章成功", c)

}
