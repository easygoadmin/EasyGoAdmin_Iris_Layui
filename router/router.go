package router

import (
	"easygoadmin/controller"
	"easygoadmin/middleware"
	"easygoadmin/widget"
	"github.com/kataras/iris/v12"
)

// 注册路由
func RegisterRouter(app *iris.Application) {

	// 登录验证中间件
	app.Use(middleware.CheckLogin)

	//视图文件目录 每次请求时自动重载模板
	tmpl := iris.HTML("./view", ".html").Reload(true)

	// 注册自定义视图函数
	tmpl.AddFunc("safe", widget.Safe)
	tmpl.AddFunc("date", widget.Date)
	tmpl.AddFunc("widget", widget.Widget)
	tmpl.AddFunc("query", widget.Query)
	tmpl.AddFunc("add", widget.Add)
	tmpl.AddFunc("edit", widget.Edit)
	tmpl.AddFunc("delete", widget.Delete)
	tmpl.AddFunc("dall", widget.Dall)
	tmpl.AddFunc("expand", widget.Expand)
	tmpl.AddFunc("collapse", widget.Collapse)
	tmpl.AddFunc("addz", widget.Addz)
	tmpl.AddFunc("switch", widget.Switch)
	tmpl.AddFunc("select", widget.Select)
	tmpl.AddFunc("submit", widget.Submit)
	tmpl.AddFunc("icon", widget.Icon)
	tmpl.AddFunc("transfer", widget.Transfer)
	tmpl.AddFunc("upload_image", widget.UploadImage)
	tmpl.AddFunc("album", widget.Album)
	//tmpl.AddFunc("item", widget.Item)
	tmpl.AddFunc("kindeditor", widget.Kindeditor)
	tmpl.AddFunc("checkbox", widget.Checkbox)
	tmpl.AddFunc("radio", widget.Radio)
	//tmpl.AddFunc("city", widget.City)

	// 注册视图
	app.RegisterView(tmpl)

	// 静态文件
	app.HandleDir("/static", "./public/static")

	// 错误请求配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("error/404.html")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.View("error/500.html")
	})

	// 登录、主页
	index := app.Party("/")
	{
		index.Get("/", controller.Index.Index)
		index.Any("/login", controller.Login.Login)
		index.Get("/captcha", controller.Login.Captcha)
		index.Get("/index", controller.Index.Index)
		index.Get("/main", controller.Index.Main)
		index.Any("/userInfo", controller.Index.UserInfo)
		index.Any("/updatePwd", controller.Index.UpdatePwd)
		index.Get("/logout", controller.Index.Logout)
	}

	// 职级管理
	level := app.Party("/level")
	{
		level.Get("/index", controller.Level.Index)
		level.Post("/list", controller.Level.List)
		level.Get("/edit", controller.Level.Edit)
		level.Post("/add", controller.Level.Add)
		level.Post("/update", controller.Level.Update)
		level.Post("/delete/:ids", controller.Level.Delete)
		level.Post("/setStatus", controller.Level.Status)
	}

}
