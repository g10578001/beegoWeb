package routers

import (
	"beegoWeb/webLiang/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.MainController{}, "get,post:Signup")
}
