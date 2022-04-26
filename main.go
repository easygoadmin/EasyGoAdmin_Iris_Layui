package main

import (
	"easygoadmin/router"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// 创建app结构体对象
	app := iris.New()
	// 设置调试模式
	app.Logger().SetLevel("debug")
	// 可选项添加两个内置的句柄（handlers）
	// 捕获相对于http产生的异常行为
	app.Use(recover.New())
	//记录请求日志
	app.Use(logger.New())
	// 初始化配置
	config := iris.WithConfiguration(iris.YAML("./config/config.yml"))
	// 路由注册
	router.RegisterRouter(app)

	//// 谓词:   GET
	//// 资源: http://localhost:8080
	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	ctx.HTML("<h1>Welcome</h1>")
	//})
	//
	//// 等价于 app.Handle("GET", "/ping", [...])
	//// 谓词:   GET
	//// 资源: http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("ping")
	//})
	//
	//// 谓词:   GET
	//// 资源: http://localhost:8080/hello
	//app.Get("/hello", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	//})
	//
	//// http://localhost:8080
	//// http://localhost:8080/ping
	//// http://localhost:8080/hello
	//// Run 方法第二个参数为应用的配置参数
	//app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

	// Listens and serves incoming http requests
	// on http://localhost:8081.
	app.Run(iris.Addr(":9087"), config)
}
