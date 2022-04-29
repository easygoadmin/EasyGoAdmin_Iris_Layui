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

var Link = new(LinkController)

type LinkController struct{}

func (c *LinkController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("link/index.html")
}

func (c *LinkController) List(ctx iris.Context) {
	// 参数
	var req dto.LinkPageReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取列表方法
	lists, count, err := service.Link.GetList(req)
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

func (c *LinkController) Edit(ctx iris.Context) {
	// 查询记录
	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		info := &model.Link{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 友链图片
		if info.Image != "" {
			info.Image = utils.GetImageUrl(info.Image)
		}

		// 数据绑定
		ctx.ViewData("info", info)
	}
	// 常量绑定
	ctx.ViewData("typeList", constant.LINK_TYPE_LIST)
	ctx.ViewData("formList", constant.LINK_FORM_LIST)
	ctx.ViewData("platformList", constant.LINK_PLATFORM_LIST)
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("link/edit.html")
}

func (c *LinkController) Add(ctx iris.Context) {
	// 添加对象
	var req dto.LinkAddReq
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
	rows, err := service.Link.Add(req, utils.Uid(ctx))
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

func (c *LinkController) Update(ctx iris.Context) {
	// 更新对象
	var req dto.LinkUpdateReq
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
	rows, err := service.Link.Update(req, utils.Uid(ctx))
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

func (c *LinkController) Delete(ctx iris.Context) {
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
	rows, err := service.Link.Delete(ids)
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

func (c *LinkController) Status(ctx iris.Context) {
	// 设置对象
	var req dto.LinkStatusReq
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
	rows, err := service.Link.Status(req, utils.Uid(ctx))
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
