package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils/res"
	"io/fs"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImagesUploadView 上传单/多个文件
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	// 获取MultipartForm
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg("上传文件失败", c)
		return
	}
	// 获取form表单中的图片
	files, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("上传图片失败", c)
		return
	}
	// 判断文件上传目录是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 不存在则创建目录
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
			res.FailWithMsg("创建目录失败", c)
			return
		}
	}
	// 响应数据
	var fileUploadResponse []FileUploadResponse
	// 遍历上传的文件
	for _, file := range files {
		// 文件保存位置
		filePath := path.Join(basePath, file.Filename)
		// 计算上传的文件占几M
		size := float64(file.Size) / float64(1024*1024)
		// 如果上传文件的大小大于等于2M的话，上传失败
		if size >= float64(global.Config.Upload.Size) {
			global.Log.Error("上传文件失败")
			fileUploadResponse = append(fileUploadResponse, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传文件大小为%.2fM,超过预定大小%dM", size, global.Config.Upload.Size),
			})
			continue // 开始第二次循环
		}
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error("保存文件失败")
			fileUploadResponse = append(fileUploadResponse, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "保存文件失败",
			})
			continue // 开启下一次循环
		}
		// 文件保存成功
		fileUploadResponse = append(fileUploadResponse, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	res.OKWithData(fileUploadResponse, c)
}
