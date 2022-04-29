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

var Ad = new(AdController)

type AdController struct{}

func (c *AdController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("ad/index.html")
}

func (c *AdController) List(ctx iris.Context) {
	// 参数
	var req dto.AdPageReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取列表方法
	lists, count, err := service.Ad.GetList(req)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回结果集
	ctx.JSON(common.JsonResult{
		Code:  0,
		Data:  lists,
		Msg:   "操作成功",
		Count: count,
	})
}

func (c *AdController) Edit(ctx iris.Context) {
	// 查询记录
	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		info := &model.Ad{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 广告图片
		if info.Cover != "" {
			info.Cover = utils.GetImageUrl(info.Cover)
		}

		// 广告位列表
		list := make([]model.AdSort, 0)
		utils.XormDb.Where("mark=1").Find(&list)
		adSortList := make(map[int]string, 0)
		for _, v := range list {
			adSortList[v.Id] = v.Description
		}
		// 数据绑定
		ctx.ViewData("info", info)
		ctx.ViewData("adSortList", adSortList)
	}
	// 绑定数据源
	ctx.ViewData("typeList", constant.AD_TYPE_LIST)
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("ad/edit.html")
}

func (c *AdController) Add(ctx iris.Context) {
	// 添加对象
	var req dto.AdAddReq
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
	rows, err := service.Ad.Add(req, utils.Uid(ctx))
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

func (c *AdController) Update(ctx iris.Context) {
	// 更新对象
	var req dto.AdUpdateReq
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
	rows, err := service.Ad.Update(req, utils.Uid(ctx))
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

func (c *AdController) Delete(ctx iris.Context) {
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
	rows, err := service.Ad.Delete(ids)
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

func (c *AdController) Status(ctx iris.Context) {
	// 设置对象
	var req dto.AdStatusReq
	// 请求验证
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
	// 调用设置状态方法
	rows, err := service.Ad.Status(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 设置成功
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}
