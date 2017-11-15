package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int64
	Name  string `orm:"size(256);unique"`
	Email string `orm:"size(256);unique"`
}

func init() {
	orm.RegisterModel(new(User))
}
