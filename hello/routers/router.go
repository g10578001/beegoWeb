package routers

import (
	"test/hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/root", &controllers.MainController{})
}
