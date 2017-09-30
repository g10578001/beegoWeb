package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "classroom.google.com"
	//c.Data["Email"] = "jcliang@cycu.org.tw"
	//c.TplName = "index.tpl"
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = "index.tpl"

}

func (c *MainController) About() {
	//c.Data["Website"] = "classroom.google.com"
	//c.Data["Email"] = "jcliang@cycu.org.tw"
	//c.TplName = "about.tpl"
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = "about.tpl"

}

func (c *MainController) SamplePost() {
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = "samplePost.tpl"

}

func (c *MainController) Contact() {
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = "contact.tpl"

}
