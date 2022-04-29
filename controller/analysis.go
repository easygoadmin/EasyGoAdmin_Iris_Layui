package controller

import "github.com/kataras/iris/v12"

var Analysis = new(AnalysisController)

type AnalysisController struct{}

func (c *AnalysisController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("analysis/index.html")
}
