package controller

import (
	"easygoadmin/constant"
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/vo"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

var User = new(UserController)

type UserController struct{}

func (c *UserController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("user/index.html")
}

func (c *UserController) List(ctx iris.Context) {
	// 参数
	var req dto.UserPageReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取列表方法
	lists, count, err := service.User.GetList(req)
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

func (c *UserController) Edit(ctx iris.Context) {
	// 获取职级
	levelAll := make([]model.Level, 0)
	utils.XormDb.Where("status=1 and mark=1").Find(&levelAll)
	levelList := make(map[int]string, 0)
	for _, v := range levelAll {
		levelList[v.Id] = v.Name
	}
	// 获取岗位
	positionAll := make([]model.Position, 0)
	utils.XormDb.Where("status=1 and mark=1").Find(&positionAll)
	positionList := make(map[int]string, 0)
	for _, v := range positionAll {
		positionList[v.Id] = v.Name
	}
	// 获取部门列表
	deptData, _ := service.Dept.GetDeptTreeList()
	deptList := service.Dept.MakeList(deptData)
	// 获取角色
	roleData := make([]model.Role, 0)
	utils.XormDb.Where("status=1 and mark=1").Find(&roleData)
	roleList := make(map[int]string)
	for _, v := range roleData {
		roleList[v.Id] = v.Name
	}

	// 查询记录
	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		info := &model.User{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		var userInfo = vo.UserInfoVo{}
		userInfo.User = *info
		// 头像
		userInfo.Avatar = utils.GetImageUrl(info.Avatar)

		// 角色ID
		var userRoleList []model.UserRole
		utils.XormDb.Where("user_id=?", utils.Uid(ctx)).Find(&userRoleList)
		roleIds := make([]interface{}, 0)
		for _, v := range userRoleList {
			roleIds = append(roleIds, v.RoleId)
		}
		userInfo.RoleIds = roleIds

		// 数据绑定
		ctx.ViewData("info", userInfo)
	}
	// 绑定参数
	ctx.ViewData("genderList", constant.GENDER_LIST)
	ctx.ViewData("levelList", levelList)
	ctx.ViewData("positionList", positionList)
	ctx.ViewData("deptList", deptList)
	ctx.ViewData("roleList", roleList)
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("user/edit.html")
}

func (c *UserController) Add(ctx iris.Context) {
	// 添加对象
	var req dto.UserAddReq
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
	rows, err := service.User.Add(req, utils.Uid(ctx))
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

func (c *UserController) Update(ctx iris.Context) {
	// 更新对象
	var req dto.UserUpdateReq
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
	rows, err := service.User.Update(req, utils.Uid(ctx))
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

func (c *UserController) Delete(ctx iris.Context) {
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
	rows, err := service.User.Delete(ids)
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

func (c *UserController) Status(ctx iris.Context) {
	// 设置对象
	var req dto.UserStatusReq
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
	rows, err := service.User.Status(req, utils.Uid(ctx))
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

func (c *UserController) ResetPwd(ctx iris.Context) {
	// 参数验证
	var req dto.UserResetPwdReq
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
	// 调用重置密码方法
	rows, err := service.User.ResetPwd(req.Id, utils.Uid(ctx))
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
		Msg:  "重置密码成功",
	})
}
