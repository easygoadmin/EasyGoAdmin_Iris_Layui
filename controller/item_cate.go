package controller

import (
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

var ItemCate = new(ItemCateController)

type ItemCateController struct{}

func (c *ItemCateController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("item_cate/index.html")
}

func (c *ItemCateController) List(ctx iris.Context) {
	// 参数
	var req dto.ItemCateQueryReq
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用获取列表方法
	lists := service.ItemCate.GetList(req)
	// 返回结果集
	ctx.JSON(common.JsonResult{
		Code: 0,
		Data: lists,
		Msg:  "操作成功",
	})
}

func (c *ItemCateController) Edit(ctx iris.Context) {
	// 站点列表
	result := make([]model.Item, 0)
	utils.XormDb.Where("mark=1").Find(&result)
	var itemList = map[int]string{}
	for _, v := range result {
		itemList[v.Id] = v.Name
	}
	// 查询记录
	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		info := &model.ItemCate{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 封面
		if info.IsCover == 1 && info.Cover != "" {
			info.Cover = utils.GetImageUrl(info.Cover)
		}
		// 数据绑定
		ctx.ViewData("info", info)
	}
	// 绑定参数
	ctx.ViewData("itemList", itemList)
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("item_cate/edit.html")
}

func (c *ItemCateController) Add(ctx iris.Context) {
	// 添加对象
	var req dto.ItemCateAddReq
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
	rows, err := service.ItemCate.Add(req, utils.Uid(ctx))
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

func (c *ItemCateController) Update(ctx iris.Context) {
	// 更新对象
	var req dto.ItemCateUpdateReq
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
	rows, err := service.ItemCate.Update(req, utils.Uid(ctx))
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

func (c *ItemCateController) Delete(ctx iris.Context) {
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
	rows, err := service.ItemCate.Delete(ids)
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

func (c *ItemCateController) GetCateList(ctx iris.Context) {
	list := make([]model.ItemCate, 0)
	utils.XormDb.Where("status=1 and mark=1").OrderBy("sort asc").Find(&list)
	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}

func (c *ItemCateController) GetCateTreeList(ctx iris.Context) {
	itemId := ctx.Params().GetIntDefault("itemId", 0)
	list, err := service.ItemCate.GetCateTreeList(itemId, 0)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 数据源转换
	result := service.ItemCate.MakeList(list)
	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: result,
	})
}
