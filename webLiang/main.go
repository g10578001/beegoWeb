package main

import (
	"beegoWeb/webLiang/models"
	_ "beegoWeb/webLiang/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	o := orm.NewOrm()
	o.Using("default")

	user := new(models.User)
	user.Username = "g02"
	user.Email = "g02@cycu.edu.tw"
	user.Password = "g02g02"
	//o.Insert(user)
	fmt.Println(o.Insert(user))

	beego.Run()
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "regis.db")
}
