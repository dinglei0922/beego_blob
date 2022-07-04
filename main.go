package main

import (
	_ "beego-blob/routers"
	"beego-blob/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "beego-blob/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.InsertFilter("/cms/main/*",beego.BeforeRouter,utils.LoginFilter)
	mysql_username := beego.AppConfig.String("mysql::username")
	mysql_password := beego.AppConfig.String("mysql::password")
	mysql_port := beego.AppConfig.String("mysql::port")
	mysql_host := beego.AppConfig.String("mysql::host")
	mysql_database := beego.AppConfig.String("mysql::database")
	datasource:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",mysql_username,mysql_password,mysql_host,mysql_port,mysql_database)
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",datasource)
	beego.Run()
}

