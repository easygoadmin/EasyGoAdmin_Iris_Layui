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

	// 视图文件目录 每次请求时自动重载模板
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
	tmpl.AddFunc("item", widget.Item)
	tmpl.AddFunc("kindeditor", widget.Kindeditor)
	tmpl.AddFunc("checkbox", widget.Checkbox)
	tmpl.AddFunc("radio", widget.Radio)
	tmpl.AddFunc("city", widget.City)

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

	// 文件上传
	upload := app.Party("/upload")
	{
		upload.Post("/uploadImage", controller.Upload.UploadImage)
		upload.Post("/uploadEditImage", controller.Upload.UploadEditImage)
	}

	// 职级管理
	level := app.Party("/level")
	{
		level.Get("/index", controller.Level.Index)
		level.Post("/list", controller.Level.List)
		level.Get("/edit/{id:int}", controller.Level.Edit)
		level.Post("/add", controller.Level.Add)
		level.Post("/update", controller.Level.Update)
		level.Post("/delete/{id:int}", controller.Level.Delete)
		level.Post("/setStatus", controller.Level.Status)
	}

	// 岗位管理
	position := app.Party("/position")
	{
		position.Get("/index", controller.Position.Index)
		position.Post("/list", controller.Position.List)
		position.Get("/edit/{id:int}", controller.Position.Edit)
		position.Post("/add", controller.Position.Add)
		position.Post("/update", controller.Position.Update)
		position.Post("/delete/{id:int}", controller.Position.Delete)
		position.Post("/setStatus", controller.Position.Status)
	}

	/* 角色路由 */
	role := app.Party("role")
	{
		role.Get("/index", controller.Role.Index)
		role.Post("/list", controller.Role.List)
		role.Get("/edit/{id:int}", controller.Role.Edit)
		role.Post("/add", controller.Role.Add)
		role.Post("/update", controller.Role.Update)
		role.Post("/delete/{id:int}", controller.Role.Delete)
		role.Post("/setStatus", controller.Role.Status)
		role.Get("/getRoleList", controller.Role.GetRoleList)
	}

	/* 角色菜单权限 */
	roleMenu := app.Party("rolemenu")
	{
		roleMenu.Get("/index/{roleId:int}", controller.RoleMenu.Index)
		roleMenu.Post("/save", controller.RoleMenu.Save)
	}

	/* 部门管理 */
	dept := app.Party("dept")
	{
		dept.Get("/index", controller.Dept.Index)
		dept.Post("/list", controller.Dept.List)
		dept.Get("/edit/{id:int}", controller.Dept.Edit)
		dept.Post("/add", controller.Dept.Add)
		dept.Post("/update", controller.Dept.Update)
		dept.Post("/delete/{id:int}", controller.Dept.Delete)
		dept.Get("/getDeptList", controller.Dept.GetDeptList)
	}

	/* 用户管理 */
	user := app.Party("user")
	{
		user.Get("/index", controller.User.Index)
		user.Post("/list", controller.User.List)
		user.Get("/edit/{id:int}", controller.User.Edit)
		user.Post("/add", controller.User.Add)
		user.Post("/update", controller.User.Update)
		user.Post("/delete/{id:int}", controller.User.Delete)
		user.Post("/setStatus", controller.User.Status)
		user.Post("/resetPwd", controller.User.ResetPwd)
	}

	/* 菜单管理 */
	menu := app.Party("menu")
	{
		menu.Get("/index", controller.Menu.Index)
		menu.Post("/list", controller.Menu.List)
		menu.Get("/edit/{id:int}", controller.Menu.Edit)
		menu.Post("/add", controller.Menu.Add)
		menu.Post("/update", controller.Menu.Update)
		menu.Post("/delete/{id:int}", controller.Menu.Delete)
	}

	/* 友链管理 */
	link := app.Party("link")
	{
		link.Get("/index", controller.Link.Index)
		link.Post("/list", controller.Link.List)
		link.Get("/edit/{id:int}", controller.Link.Edit)
		link.Post("/add", controller.Link.Add)
		link.Post("/update", controller.Link.Update)
		link.Post("/delete/{id:int}", controller.Link.Delete)
		link.Post("/setStatus", controller.Link.Status)
	}

	/* 城市管理 */
	city := app.Party("city")
	{
		city.Get("/index", controller.City.Index)
		city.Post("/list", controller.City.List)
		city.Get("/edit/{id:int}", controller.City.Edit)
		city.Post("/add", controller.City.Add)
		city.Post("/update", controller.City.Update)
		city.Post("/delete/{id:int}", controller.City.Delete)
		city.Post("/getChilds", controller.City.GetChilds)
	}

	/* 通知管理 */
	notice := app.Party("notice")
	{
		notice.Get("/index", controller.Notice.Index)
		notice.Post("/list", controller.Notice.List)
		notice.Get("/edit/{id:int}", controller.Notice.Edit)
		notice.Post("/add", controller.Notice.Add)
		notice.Post("/update", controller.Notice.Update)
		notice.Post("/delete/{id:int}", controller.Notice.Delete)
		notice.Post("/setStatus", controller.Notice.Status)
	}

	/* 会员等级 */
	memberlevel := app.Party("memberlevel")
	{
		memberlevel.Get("/index", controller.MemberLevel.Index)
		memberlevel.Post("/list", controller.MemberLevel.List)
		memberlevel.Get("/edit/{id:int}", controller.MemberLevel.Edit)
		memberlevel.Post("/add", controller.MemberLevel.Add)
		memberlevel.Post("/update", controller.MemberLevel.Update)
		memberlevel.Post("/delete/{id:int}", controller.MemberLevel.Delete)
		memberlevel.Get("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	}

	/* 会员管理 */
	member := app.Party("member")
	{
		member.Get("/index", controller.Member.Index)
		member.Post("/list", controller.Member.List)
		member.Get("/edit/{id:int}", controller.Member.Edit)
		member.Post("/add", controller.Member.Add)
		member.Post("/update", controller.Member.Update)
		member.Post("/delete/{id:int}", controller.Member.Delete)
		member.Post("/setStatus", controller.Member.Status)
	}

	/* 字典管理 */
	dict := app.Party("dict")
	{
		dict.Get("/index", controller.Dict.Index)
		dict.Post("/list", controller.Dict.List)
		dict.Post("/add", controller.Dict.Add)
		dict.Post("/update", controller.Dict.Update)
		dict.Post("/delete/{id:int}", controller.Dict.Delete)
	}

	/* 字典项管理 */
	dictdata := app.Party("dictdata")
	{
		dictdata.Post("/list", controller.DictData.List)
		dictdata.Post("/add", controller.DictData.Add)
		dictdata.Post("/update", controller.DictData.Update)
		dictdata.Post("/delete/{id:int}", controller.DictData.Delete)
	}

	/* 配置管理 */
	config := app.Party("config")
	{
		config.Get("/index", controller.Config.Index)
		config.Post("/list", controller.Config.List)
		config.Post("/add", controller.Config.Add)
		config.Post("/update", controller.Config.Update)
		config.Post("/delete/{id:int}", controller.Config.Delete)
	}

	/* 配置项管理 */
	configdata := app.Party("configdata")
	{
		configdata.Post("/list", controller.ConfigData.List)
		configdata.Post("/add", controller.ConfigData.Add)
		configdata.Post("/update", controller.ConfigData.Update)
		configdata.Post("/delete/{id:int}", controller.ConfigData.Delete)
		configdata.Post("/setStatus", controller.ConfigData.Status)
	}

	/* 网站设置 */
	configweb := app.Party("configweb")
	{
		configweb.Any("/index", controller.ConfigWeb.Index)
	}

	/* 站点管理 */
	item := app.Party("item")
	{
		item.Get("/index", controller.Item.Index)
		item.Post("/list", controller.Item.List)
		item.Get("/edit/{id:int}", controller.Item.Edit)
		item.Post("/add", controller.Item.Add)
		item.Post("/update", controller.Item.Update)
		item.Post("/delete/{id:int}", controller.Item.Delete)
		item.Post("/setStatus", controller.Item.Status)
		item.Get("/getItemList", controller.Item.GetItemList)
	}

	/* 栏目管理 */
	itemcate := app.Party("itemcate")
	{
		itemcate.Get("/index", controller.ItemCate.Index)
		itemcate.Post("/list", controller.ItemCate.List)
		itemcate.Get("/edit/{id:int}", controller.ItemCate.Edit)
		itemcate.Post("/add", controller.ItemCate.Add)
		itemcate.Post("/update", controller.ItemCate.Update)
		itemcate.Post("/delete/{id:int}", controller.ItemCate.Delete)
		itemcate.Get("/getCateList", controller.ItemCate.GetCateList)
		itemcate.Get("/getCateTreeList", controller.ItemCate.GetCateTreeList)
	}

	/* 广告位管理 */
	adsort := app.Party("adsort")
	{
		adsort.Get("/index", controller.AdSort.Index)
		adsort.Post("/list", controller.AdSort.List)
		adsort.Get("/edit/{id:int}", controller.AdSort.Edit)
		adsort.Post("/add", controller.AdSort.Add)
		adsort.Post("/update", controller.AdSort.Update)
		adsort.Post("/delete/{id:int}", controller.AdSort.Delete)
		adsort.Get("/getAdSortList", controller.AdSort.GetAdSortList)
	}

	/* 广告管理 */
	ad := app.Party("ad")
	{
		ad.Get("/index", controller.Ad.Index)
		ad.Post("/list", controller.Ad.List)
		ad.Get("/edit/{id:int}", controller.Ad.Edit)
		ad.Post("/add", controller.Ad.Add)
		ad.Post("/update", controller.Ad.Update)
		ad.Post("/delete/{id:int}", controller.Ad.Delete)
		ad.Post("/setStatus", controller.Ad.Status)
	}

	/* 统计分析 */
	analysis := app.Party("analysis")
	{
		analysis.Get("/index", controller.Analysis.Index)
	}

	/* 代码生成器 */
	generate := app.Party("generate")
	{
		generate.Get("/index", controller.Generate.Index)
		generate.Post("/list", controller.Generate.List)
		generate.Post("/generate", controller.Generate.Generate)
		generate.Post("/batchGenerate", controller.Generate.BatchGenerate)
	}

	/* 演示一 */
	example := app.Party("example")
	{
		example.Get("/index", controller.Example.Index)
		example.Post("/list", controller.Example.List)
		example.Get("/edit/{id:int}", controller.Example.Edit)
		example.Post("/add", controller.Example.Add)
		example.Post("/update", controller.Example.Update)
		example.Post("/delete/{id:int}", controller.Example.Delete)
		example.Post("/setStatus", controller.Example.Status)
		example.Post("/setIsVip", controller.Example.IsVip)
	}

}
