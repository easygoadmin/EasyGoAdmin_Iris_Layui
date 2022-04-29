package controller

import (
	"easygoadmin/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/kataras/iris/v12"
)

// 控制器管理对象
var Upload = new(UploadController)

type UploadController struct{}

func (c *UploadController) UploadImage(ctx iris.Context) {
	// 调用上传方法
	result, err := service.Upload.UploadImage(ctx)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 拼接图片地址
	result.FileUrl = utils.GetImageUrl(result.FileUrl)
	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "上传成功",
		Data: result,
	})
}

func (c *UploadController) UploadEditImage(ctx iris.Context) {
	// 调用上传方法
	result, err := service.Upload.UploadImage(ctx)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 拼接图片地址
	fileUrl := utils.GetImageUrl(result.FileUrl)
	// 返回结果
	ctx.JSON(common.JsonEditResult{
		Error: 0,
		Url:   fileUrl,
	})
}
