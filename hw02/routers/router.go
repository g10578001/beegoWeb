package routers

import (
	"beegoWeb/hw02/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{}, "get,post:Home")
	beego.Router("/about", &controllers.MainController{}, "get,post:About")
	beego.Router("/samplePost", &controllers.MainController{}, "get,post:SamplePost")
	beego.Router("/contact", &controllers.MainController{}, "get,post:Contact")
}
