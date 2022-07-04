package routers

import (
	"beego-blob/controllers/cms"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/cms",&cms.LoginController{})
}
