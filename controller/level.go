package controller

import (
	"github.com/kataras/iris/v12"
)

var Level = new(LevelController)

type LevelController struct{}

func (c *LevelController) Index(ctx iris.Context) {
	ctx.View("level/index.html")
}

func (c *LevelController) List(ctx iris.Context) {

}

func (c *LevelController) Edit(ctx iris.Context) {

}

func (c *LevelController) Add(ctx iris.Context) {

}

func (c *LevelController) Update(ctx iris.Context) {

}

func (c *LevelController) Delete(ctx iris.Context) {

}

func (c *LevelController) Status(ctx iris.Context) {

}
