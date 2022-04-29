package controller

import (
	"easygoadmin/dto"
	"easygoadmin/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

var Index = new(IndexController)

type IndexController struct{}

func (c *IndexController) Index(ctx iris.Context) {
	// 获取用户信息
	userInfo := service.Login.GetProfile(utils.Uid(ctx))
	// 获取菜单列表
	menuList := service.Menu.GetPermissionMenuList(userInfo.Id)
	ctx.ViewData("userInfo", userInfo)
	ctx.ViewData("menuList", menuList)
	// 渲染模板
	ctx.View("index.html")
}

func (c *IndexController) Main(ctx iris.Context) {
	// 渲染模板
	ctx.View("welcome.html")
}

func (c *IndexController) UserInfo(ctx iris.Context) {
	if ctx.Method() == "POST" {
		// 参数验证
		var req dto.UserInfoReq
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
		// 更新信息
		_, err := service.User.UpdateUserInfo(req, utils.Uid(ctx))
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
			Msg:  "更新成功",
		})
		return
	}
	// 获取用户信息
	userInfo := service.Login.GetProfile(utils.Uid(ctx))
	// 绑定数据
	ctx.ViewData("userInfo", userInfo)
	// 渲染模板
	ctx.View("user_info/index.html")
}

func (c *IndexController) UpdatePwd(ctx iris.Context) {
	// 参数验证
	var req dto.UpdatePwd
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
	// 调用更新密码方法
	rows, err := service.User.UpdatePwd(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新密码成功",
	})
}

func (c *IndexController) Logout(ctx iris.Context) {
	// 清除SESSION
	common.Session.Start(ctx).Clear()
	// 跳转登录页
	ctx.Redirect("/login")
}
