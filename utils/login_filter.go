package utils

import "github.com/astaxie/beego/context"

func LoginFilter(ctx *context.Context){
	username:=ctx.Input.Session("username")
	if username==nil {
		ctx.Redirect(302,"/cms/login")
	}
}