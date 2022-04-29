package service

import (
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gstr"
	"errors"
	"github.com/kataras/iris/v12"
	"time"
)

var Login = new(loginService)

type loginService struct{}

// 系统登录
func (s *loginService) UserLogin(username, password string, ctx iris.Context) error {
	// 查询用户
	var user model.User
	has, err := utils.XormDb.Where("username=? and mark=1", username).Get(&user)
	if err != nil && !has {
		return errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.Md5(password + user.Username)
	if user.Password != pwd {
		return errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		return errors.New("您的账号已被禁用,请联系管理员")
	}

	// 更新登录时间、登录IP
	utils.XormDb.Id(user.Id).Update(&model.User{LoginTime: time.Now().Unix(), LoginIp: "", UpdateTime: time.Now().Unix()})

	// 初始化session对象
	session := common.Session.Start(ctx)
	// 写入SESSION
	session.Set(common.USER_ID, user.Id)
	// 返回token
	return nil
}

// 获取个人信息
func (s *loginService) GetProfile(userId int) (user *model.User) {
	user = &model.User{Id: userId}
	has, err := user.Get()
	if err != nil || !has {
		return nil
	}
	// 头像
	if user.Avatar != "" && !gstr.Contains(user.Avatar, utils.ImageUrl()) {
		user.Avatar = utils.GetImageUrl(user.Avatar)
	}
	return
}
