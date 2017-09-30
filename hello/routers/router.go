package routers

import (
	"beegoWeb/hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.MainController{}, "get,post:About")
	beego.Router("/samplePost", &controllers.MainController{}, "get,post:SamplePost")
	beego.Router("/contact", &controllers.MainController{}, "get,post:Contact")
}
