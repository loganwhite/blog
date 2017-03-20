package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/login", &controllers.AdminController{}, "get:AdminLogin")
}
