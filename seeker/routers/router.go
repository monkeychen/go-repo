package routers

import (
	"github.com/astaxie/beego"
	"seeker/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.MainController{})
}
