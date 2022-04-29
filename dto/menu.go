package dto

import (
	"easygoadmin/model"
	"github.com/gookit/validate"
)

// 列表查询条件
type MenuPageReq struct {
	Title string `form:"name"` // 菜单标题
}

// 添加菜单
type MenuAddReq struct {
	Id         string `form:"id"`
	Name       string `form:"name" validate:"required"` // 菜单标题
	Icon       string `form:"icon"`                     // 图标
	Url        string `form:"url" validate:"required"`  // URL地址
	Param      string `form:"param"`                    // 参数
	Pid        int    `form:"pid" validate:"int"`       // 上级ID
	Type       int    `form:"type"`                     // 类型：1模块 2导航 3菜单 4节点
	Permission string `form:"permission"`               // 权限标识
	Status     int    `form:"status" validate:"int"`    // 状态：1正常 2禁用
	Target     string `form:"target"`                   // 打开方式：1内部打开 2外部打开
	Note       string `form:"note"`                     // 菜单备注
	Sort       int    `form:"sort" validate:"int"`      // 显示顺序
	Func       string `form:"func"`                     // 权限节点
}

// 添加菜单表单验证
func (v MenuAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "菜单名称不能为空.",
		"Pid.int":       "请选择上级菜单.",
		"Type.int":      "请选择菜单类型.",
		"Status.int":    "请选择菜单状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新菜单
type MenuUpdateReq struct {
	Id         int    `form:"id" validate:"int"`
	Name       string `form:"name" validate:"required"` // 菜单标题
	Icon       string `form:"icon"`                     // 图标
	Url        string `form:"url" validate:"required"`  // URL地址
	Param      string `form:"param"`                    // 参数
	Pid        int    `form:"pid" validate:"int"`       // 上级ID
	Type       int    `form:"type"`                     // 类型：1模块 2导航 3菜单 4节点
	Permission string `form:"permission"`               // 权限标识
	Status     int    `form:"status" validate:"int"`    // 状态：1正常 2禁用
	Target     string `form:"target"`                   // 打开方式：1内部打开 2外部打开
	Note       string `form:"note"`                     // 菜单备注
	Sort       int    `form:"sort" validate:"int"`      // 显示顺序
	Func       string `form:"func"`                     // 权限节点
}

// 添加菜单表单验证
func (v MenuUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "菜单ID不能为空.",
		"Name.required": "菜单名称不能为空.",
		"Pid.int":       "请选择上级菜单.",
		"Type.int":      "请选择菜单类型.",
		"Status.int":    "请选择菜单状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 菜单信息
type MenuInfoVo struct {
	model.Menu
	CheckedList []int `form:"checkedList"` // 权限节点列表
}
