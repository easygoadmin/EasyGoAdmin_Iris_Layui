package controller

import (
	"easygoadmin/dto"
	"easygoadmin/service"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

var RoleMenu = new(RoleMenuController)

type RoleMenuController struct{}

func (c *RoleMenuController) Index(ctx iris.Context) {
	// 角色ID
	roleId := ctx.Params().GetIntDefault("roleId", 0)
	if roleId == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  "角色ID不能为空",
		})
		return
	}
	// 获取角色菜单权限列表
	list, err := service.RoleMenu.GetRoleMenuList(gconv.Int(roleId))
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
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *RoleMenuController) Save(ctx iris.Context) {
	// 角色菜单对象
	var req dto.RoleMenuSaveReq
	// 参数绑定
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 参数验证
	v := validate.Struct(req)
	if !v.Validate() {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
		return
	}
	// 调用保存方法
	err := service.RoleMenu.Save(req)
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
		Msg:  "保存成功",
	})
}
