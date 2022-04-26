package middleware

import "github.com/kataras/iris/v12"

// 登录验证中间件
func CheckLogin(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
