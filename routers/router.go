package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:RedirectIndex")
	beego.Router("/index.html", &controllers.IndexController{}, "get:RedirectIndex")
	beego.Router("/admin/login", &controllers.AdminController{}, "get:AdminLogin")
}
