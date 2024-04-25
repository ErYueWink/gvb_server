package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"gvb_server/utils/res"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

// ImagesUploadDataView 上传单个图片
func (ImagesApi) ImagesUploadDataView(c *gin.Context) {
	file, err := c.FormFile("images")
	if err != nil {
		global.Log.Error(err.Error())
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	basePath := global.Config.Upload.Path // 文件上传目录
	maxSize := global.Config.Upload.Size
	// 获取文件名，截取文件后缀名
	fileName := file.Filename
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1]) // 文件后缀名
	// 判断上传的文件和白名单中预先设置好的图片是否一致
	if !utils.InList(suffix, WhiteImageList) {
		res.FailWithMsg("文件非法", c)
		return
	}
	// 判断上传文件大小
	fileSize := float64(file.Size) / float64(1024*1024)
	if fileSize > float64(maxSize) {
		res.FailWithMsg(fmt.Sprintf("上传文件大小为%.2fM,超过预定大小%dM", fileSize, global.Config.Upload.Size), c)
		return
	}
	// 判断文件保存的目录是否存在
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 创建这个目录
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
			res.FailWithMsg("目录创建失败", c)
			return
		}
	}
	// 计算出图片的hash值
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("获取图片数据出错", c)
		return
	}
	fileBytes, err := io.ReadAll(fileObj)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("获取图片数据出错", c)
		return
	}
	imageHash := utils.Md5(fileBytes)
	count := global.DB.Take(&models.BannerModel{}, "hash = ?", imageHash).RowsAffected
	if count != 0 {
		res.FailWithMsg("图片重复，请重新上传", c)
		return
	}
	if global.Config.QiNiu.Enable {
		qiniuPath, err := qiniu.UploadImage(fileBytes, fileName, "gvb")
		if err != nil {
			res.FailWithMsg("图片上传失败", c)
			return
		}
		global.DB.Create(&models.BannerModel{
			Name:      fileName,
			Path:      qiniuPath,
			Hash:      imageHash,
			ImageType: ctype.QINIU,
		})
		res.OKWithData(qiniuPath, c)
		return
	}
	// 保存图片
	filePath := path.Join(basePath, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMsg("文件保存失败", c)
		return
	}
	global.DB.Create(&models.BannerModel{
		Name:      fileName,
		Path:      filePath,
		Hash:      imageHash,
		ImageType: ctype.LOCAL,
	})
	// 文件上传成功，返回文件路径
	res.OKWithData("/"+filePath, c)
}
