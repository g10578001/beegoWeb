package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "classroom.google.com"
	c.Data["Email"] = "jcliang@cycu.org.tw"
	c.TplName = "index.tpl"
}
