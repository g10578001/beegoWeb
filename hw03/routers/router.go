package routers

import (
	"beegoWeb/hw03/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.MainController{}, "get,post:Signup")
}
