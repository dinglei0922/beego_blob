package routers

import (
	"beego-blob/controllers/cms"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/cms/login",&cms.LoginController{},"get:Login")
    beego.Router("/cms/loginpost",&cms.LoginController{},"post:LoginPost")
    beego.Router("/cms/main/index",&cms.MainController{},"get:Index")
}
