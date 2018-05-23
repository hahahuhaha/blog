package routers

import (
	"github.com/astaxie/beego"
	"wshhz.com/blog/controllers"
)

func init() {
	beego.AutoRouter(&controllers.AdminController{})
	//beego.Router("/", &controllers.MainController{})
}
