package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int
	Name  string `orm:size(256);unique` //引號``裡的東西是給orm看的
	Email string `orm:size(256;unique`
}

func init() {
	orm.RegisterModel(new(User)) //註冊model
}
