package main

import (
	_ "beegoWeb/test/regis/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	user := new(modles.User)
	user.Name = "JYW"
	user.Email = "jyw@mail.com"

	o := orm.NewOrm()
	o.Using("default")

	fmt.Println(o.Insert(user))
	beego.Run()
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRsqlite)
	orm.RegisterDataBase("user", "regis.db")
}
