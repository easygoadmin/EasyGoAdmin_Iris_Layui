package controller

import (
	"easygoadmin/dto"
	"easygoadmin/service"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
)

var Generate = new(GenerateController)

type GenerateController struct{}

func (c *GenerateController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("generate/index.html")
}

func (c *GenerateController) List(ctx iris.Context) {
	// 参数验证
	var req dto.GeneratePageReq
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用查询列表方法
	list, err := service.Generate.GetList(req)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  list,
		Count: gconv.Int64(len(list)),
	})
}

func (c *GenerateController) Generate(ctx iris.Context) {
	// 参数验证
	var req dto.GenerateFileReq
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用生成方法
	err := service.Generate.Generate(req, ctx)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "模块生成成功",
	})
}

func (c *GenerateController) BatchGenerate(ctx iris.Context) {
	// 生成对象
	var req dto.BatchGenerateFileReq
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 参数分析
	tableList := strings.Split(req.Tables, ",")
	count := 0
	for _, item := range tableList {
		itemList := strings.Split(item, "|")
		// 组装参数对象
		var param dto.GenerateFileReq
		param.Name = itemList[0]
		param.Comment = itemList[1]
		// 调用生成方法
		err := service.Generate.Generate(param, ctx)
		if err != nil {
			continue
		}
		count++
	}
	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "本次共生成【" + strconv.Itoa(count) + "】个模块文件",
	})
}
