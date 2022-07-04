package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int
	Username string
	Password string
	IsAdmin int
	CreateTime time.Time
	Cover string
}

func (u *User) TableName()string{
	return "sys_user"
}

func init(){
	orm.RegisterModel(new(User))
}
