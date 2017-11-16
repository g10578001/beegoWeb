package main

import (
	"beegoWeb/test/regis/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	o := orm.NewOrm()
	o.Using("default")

	user := new(models.User)
	user.Name = "JYW"
	user.Email = "jyw@mail.com"

	fmt.Println(o.Insert(user))
	beego.Run()
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "regis.db")
}
