package controllers

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"kube_web/common"
	"kube_web/common/jwt"
	"kube_web/dto"
	"kube_web/models"
	"time"
)

type Login struct {
	BaseController
}

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 设置自带的store
var store = base64Captcha.DefaultMemStore

// post /v1/login  form{Id:yc2ZdNgoNc4zKTQC0Vnh,Username:admin,Password:admin,Code:8}

func (this *Login) Login() {
	var authUser = new(dto.AuthUser)
	err := this.ParseForm(authUser)
	if err == nil {
		currentUser, err := models.GetUserByUsername(authUser.Username)
		if err != nil{
			this.Fail("用户或密码错误",401)
			return
		}
		//校验验证码  clear == true 验证完会删除验证码
		if !store.Verify(authUser.Id, authUser.Code, true) {
			this.Fail("验证码不对",400)
			return
		}
		if !common.ComparePwd(currentUser.Password,[]byte(authUser.Password)) {
			this.Fail("用户或密码错误",402)
			return
		}else{
			token,_ := jwt.GenerateToken(currentUser,time.Hour*24*100)
			var respUser = new(dto.LoginRes)
			respUser.Token = token
			respUser.User = currentUser
			this.Ok(respUser)
			return
		}
	}else {
		this.Fail(err.Error(),5004)
		return
	}
	this.Fail(err.Error(),400)
	return
}

func (c *Login) Captcha(){
	GenerateCaptcha(c.Ctx)
	c.ServeJSON()
}

// 生成图形化验证码  ctx *context.Context
func GenerateCaptcha(ctx *context.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverMath

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverMath{
		Height:          38,
		Width:           110,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 自定义配置，如果不需要自定义配置，则上面的结构体和下面这行代码不用写
	driverString = captchaConfig
	driver = driverString.ConvertFonts()

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		logs.Error(err.Error())
	}
	captchaResult := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
	}

	ctx.Output.JSON(SuccessData(captchaResult),true,true)
}