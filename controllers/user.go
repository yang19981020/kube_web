package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/validation"
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/common"
	"kube_web/common/jwt"
	"kube_web/dto"
	"kube_web/models"
)

// 用户 API
type UserController struct {
	BaseController
}

func (c *UserController) GetAll() {
	deptId, _ := c.GetInt("deptId",-1)
	enabled, _ := c.GetInt("enabled",-1)
	_,list := models.GetAllUser("sys_user",deptId,enabled,"")
	c.Ok(list)
}

func (c *UserController) Post()  {
	var model models.SysUser
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message,5001)
		}
	}
	e := models.AddUser(&model)
	if e != nil {
		c.Fail(e.Error(),5002)
	}
	c.Ok("操作成功")
}

func (c *UserController) Put()  {
	var model models.SysUser
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message,400)
		}
	}
	e := models.UpdateByUser(&model)
	if e != nil {
		c.Fail(e.Error(),400)
	}
	c.Ok("操作成功")
}

func (c *UserController) Delete() {
	var id int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &id)
	e := models.DelByUser(id)
	if e != nil {
		c.Fail(e.Error(),400)
	}
	c.Ok("操作成功")
}

// 用户上传图像

func (c *UserController) Avatar()  {
	f, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	defer f.Close()
	var path = "static/upload/" + h.Filename
	e := c.SaveToFile("file", path) // 保存位置在 static/upload, 没有文件夹要先创建
	logs.Error(e)
	apiUrl, _ := beego.AppConfig.String("api_url")
	avatarUrl := apiUrl + "/" +path

	//save user
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	user, _ := models.GetUserById(uid)
	if user == nil {
		c.Fail("非法操作",5006)
	}else {
		user.Avatar = avatarUrl
		models.UpdateCurrentUser(user)
		c.Ok("操作成功")
	}
}

func (c *UserController) Pass()  {
	var model dto.UserPass
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message,5007)
		}
	}
	//save user
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	user, _ := models.GetUserById(uid)
	if user == nil {
		c.Fail("非法操作",5008)
	}else {
		if !common.ParsPwd(user.Password,[]byte(model.OldPass)) {
			c.Fail("旧密码错误密码错误",5009)
		}
		user.Password = common.HashAndSalt([]byte(model.NewPass))
		models.UpdateCurrentUser(user)
		c.Ok("ok")
	}
}

// 用户修改个人信息

func (c *UserController) Center()  {
	var model dto.UserPost
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message,50010)
		}
	}
	//save user
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	user, _ := models.GetUserById(uid)
	if user == nil {
		c.Fail("非法操作",50011)
	}else {
		user.Phone = model.Phone
		user.Sex = model.Sex
		user.NickName = model.NickName
		models.UpdateCurrentUser(user)
		c.Ok("ok")
	}
}