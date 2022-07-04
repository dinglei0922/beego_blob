package main

import (
	_ "beego-blob/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	mysql_username := beego.AppConfig.String("username")
	mysql_password := beego.AppConfig.String("password")
	mysql_port := beego.AppConfig.String("port")
	mysql_host := beego.AppConfig.String("host")
	mysql_database := beego.AppConfig.String("database")
	datasource:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",mysql_username,mysql_password,mysql_host,mysql_port,mysql_database)
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",datasource)
	beego.Run()
}

