package controllers

import (
	"fmt"

	"beegoWeb/webLiang/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type MainController struct {
	beego.Controller
}

var frm_data models.User

func (c *MainController) Get() {
	c.Data["Website"] = "classroom.google.com"
	c.Data["Email"] = "jcliang@cycu.org.tw"
	c.TplName = "index.tpl"
}

func (c *MainController) Signup() {
	c.TplName = "signup.tpl"

	if c.Ctx.Input.Method() == "POST" {
		fmt.Println("PPPPPPPost!")

		c.ParseForm(&frm_data)
		o := orm.NewOrm()
		o.Using("default")
		user := new(models.User)
		user.Username = c.GetString("User_name")
		user.Email = c.GetString("email")
		user.Password = c.GetString("password")
		fmt.Println(user.Username + ",\n" + user.Email + ",\n" + user.Password)
		o.Insert(user)
		fmt.Println(o.Insert(user))
	}
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "regis.db")
}
