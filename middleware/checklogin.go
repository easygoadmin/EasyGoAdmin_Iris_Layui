package middleware

import (
	"easygoadmin/utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
)

// 登录验证中间件
func CheckLogin(ctx iris.Context) {
	fmt.Println("登录验证中间件")
	// 放行设置
	urlItem := []string{"/captcha", "/login"}
	if !utils.InStringArray(ctx.Path(), urlItem) && !strings.Contains(ctx.Path(), "static") {
		if !utils.IsLogin(ctx) {
			// 跳转登录页,方式：301(永久移动),308(永久重定向),307(临时重定向)
			ctx.Redirect("/login", http.StatusTemporaryRedirect)
			return
		}
	}
	// 前置中间件
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
