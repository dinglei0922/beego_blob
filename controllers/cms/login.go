package cms

import (
	"beego-blob/models"
	"beego-blob/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ReturnData struct {
	Code int
	Msg string
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login(){
	c.TplName = "cms/login.html"
}

func (c *LoginController) LoginPost(){
	username:=c.GetString("username","password")
	password:=c.GetString("password")
	var returnData ReturnData
	if username=="" || password=="" {
		returnData.Code = -200
		returnData.Msg = "参数错误"
		c.Data["json"] = returnData
		c.ServeJSON()
	}
	password = utils.GetMd5(password)
	ormobj := orm.NewOrm()
	if !ormobj.QueryTable(new(models.User)).Filter("username",username).Filter("password",password).Exist(){
		returnData.Code = -200
		returnData.Msg = "用户数据不存在"
		c.Data["json"] = returnData
		c.ServeJSON()
	}else{
		c.SetSession("username",username)
		c.Redirect("/cms/main/index",302)
	}
}