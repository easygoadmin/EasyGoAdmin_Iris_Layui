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

var Menu = new(MenuController)

type MenuController struct{}

func (c *MenuController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("menu/index.html")
}

func (c *MenuController) List(ctx iris.Context) {
	// 参数
	var req dto.MenuPageReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取列表方法
	lists, err := service.Menu.GetList(req)
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

func (c *MenuController) Edit(ctx iris.Context) {
	// 获取菜单列表
	menuTreeList, _ := service.Menu.GetTreeList()
	// 数据源转换
	menuList := service.Menu.MakeList(menuTreeList)
	// 查询记录
	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		info := &model.Menu{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		// 获取节点
		funcList := make([]model.Menu, 0)
		utils.XormDb.Where("pid=? and type=1 and mark=1", id).Find(&funcList)
		sortList := make([]interface{}, 0)
		for _, v := range funcList {
			sortList = append(sortList, v.Sort)
		}

		// 数据绑定
		ctx.ViewData("info", info)
		ctx.ViewData("funcList", sortList)
	} else {
		// 添加
		pid := ctx.Params().GetIntDefault("pid", 0)
		var info model.Menu
		info.Pid = pid
		info.Status = 1
		info.Target = 1
		ctx.ViewData("info", info)
		ctx.ViewData("funcList", make([]interface{}, 0))
	}
	// 绑定常量
	ctx.ViewData("menuList", menuList)
	ctx.ViewData("typeList", constant.MENU_TYPE_LIST)
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("menu/edit.html")
}

func (c *MenuController) Add(ctx iris.Context) {
	// 添加对象
	var req dto.MenuAddReq
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
	rows, err := service.Menu.Add(req, utils.Uid(ctx))
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

func (c *MenuController) Update(ctx iris.Context) {
	// 更新对象
	var req dto.MenuUpdateReq
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
	rows, err := service.Menu.Update(req, utils.Uid(ctx))
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

func (c *MenuController) Delete(ctx iris.Context) {
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
	rows, err := service.Menu.Delete(ids)
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
