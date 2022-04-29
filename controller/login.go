package controller

import (
	"easygoadmin/dto"
	"easygoadmin/service"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
	"github.com/mojocn/base64Captcha"
)

var Login = new(LoginController)

type LoginController struct{}

func (c *LoginController) Login(ctx iris.Context) {
	if ctx.Method() == "POST" {
		// 登录参数
		var req dto.LoginReq
		// 参数绑定
		if err := ctx.ReadForm(&req); err != nil {
			// 返回错误信息
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  "登录错误",
			})
			return
		}
		// 参数校验
		v := validate.Struct(req)
		if !v.Validate() {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  v.Errors.One(),
			})
			return
		}
		// 校验验证码
		verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		if !verifyRes {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  "验证码不正确",
			})
			return
		}
		// 系统登录
		err := service.Login.UserLogin(req.UserName, req.Password, ctx)
		if err != nil {
			// 登录错误
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 登录成功
		ctx.JSON(common.JsonResult{
			Code: 0,
			Msg:  "登录成功",
		})
		return
	}
	// 渲染登录界面
	ctx.View("login.html")
}

func (c *LoginController) Captcha(ctx iris.Context) {
	// 验证码参数配置：字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	///create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	// 返回结果集
	ctx.JSON(common.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}
