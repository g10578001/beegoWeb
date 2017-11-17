package main

import (
	"beegoWeb/test/b02/models"
	_ "beegoWeb/test/b02/routers"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	o := orm.NewOrm()  //宣告一個新的orm
	o.Using("default") //選擇使用那一個資料庫，使用資料庫別名來作選擇

	user := new(models.User) //宣告一個新的User結構
	user.Name = "g03"
	user.Email = "g03@cycu.edu.tw"
	fmt.Println(o.Insert(user)) //Insert資料並印出訊息
	beego.Run()
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)            //註冊使用哪一個資料庫
	orm.RegisterDataBase("default", "sqlite3", "regis.db") //註冊table以及資料庫
}
