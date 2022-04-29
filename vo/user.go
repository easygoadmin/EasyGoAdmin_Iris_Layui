package vo

import "easygoadmin/model"

// 用户信息Vo
type UserInfoVo struct {
	model.User
	GenderName   string      `json:"genderName"`   // 性别
	LevelName    string      `json:"levelName"`    // 职级
	PositionName string      `json:"positionName"` // 岗位
	DeptName     string      `json:"deptName"`     // 部门
	RoleIds      interface{} `json:"roleIds"`      // 角色ID
	RoleList     interface{} `json:"roleList"`     // 角色列表
	City         interface{} `json:"city"`         // 省市区
}

// 个人信息Vo
type ProfileInfoVo struct {
	Realname       string        `json:"realname"`       // 真实姓名
	Nickname       string        `json:"nickname"`       // 昵称
	Gender         int           `json:"gender"`         // 性别:1男 2女 3保密
	Avatar         string        `json:"avatar"`         // 头像
	Mobile         string        `json:"mobile"`         // 手机号码
	Email          string        `json:"email"`          // 邮箱地址
	City           []string      `json:"city"`           // 省市区
	Address        string        `json:"address"`        // 详细地址
	Intro          string        `json:"intro"`          // 个人简介
	Roles          []interface{} `json:"roles"`          // 用户角色
	Authorities    []interface{} `json:"authorities"`    // 用户权限
	PermissionList []string      `json:"permissionList"` // 权限列表
}
