package controller

import (
	"easygoadmin/constant"
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

var City = new(CityController)

type CityController struct{}

func (c *CityController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("city/index.html")
}

func (c *CityController) List(ctx iris.Context) {
	// 参数
	var req dto.CityPageReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取列表方法
	lists, err := service.City.GetList(req)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回结果集
	ctx.JSON(common.JsonResult{
		Code: 0,
		Data: lists,
		Msg:  "操作成功",
	})
}

func (c *CityController) Edit(ctx iris.Context) {
	// 查询记录
	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		info := &model.City{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 数据绑定
		ctx.ViewData("info", info)
	}
	ctx.ViewData("levelList", constant.CITY_LEVEL)
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("city/edit.html")
}

func (c *CityController) Add(ctx iris.Context) {
	// 添加对象
	var req dto.CityAddReq
	// 参数绑定
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 参数校验
	v := validate.Struct(&req)
	if !v.Validate() {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
		return
	}
	// 调用添加方法
	rows, err := service.City.Add(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 添加成功
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *CityController) Update(ctx iris.Context) {
	// 更新对象
	var req dto.CityUpdateReq
	// 参数绑定
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 参数校验
	v := validate.Struct(&req)
	if !v.Validate() {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
		return
	}
	// 调用更新方法
	rows, err := service.City.Update(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 更新成功
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *CityController) Delete(ctx iris.Context) {
	// 记录ID
	ids := ctx.Params().GetString("id")
	if ids == "" {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}
	// 调用删除方法
	rows, err := service.City.Delete(ids)
	if err != nil || rows == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 删除成功
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (c *CityController) GetChilds(ctx iris.Context) {
	// 参数验证
	var req dto.CityChildReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回结果
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取子级城市
	list, err := service.City.GetChilds(req.CityCode)
	if err != nil {
		// 返回结果
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}
