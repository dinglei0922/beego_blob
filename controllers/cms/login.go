package cms

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get(){
	c.TplName = "cms/login.html"
}

func (c *LoginController) Post(){

}