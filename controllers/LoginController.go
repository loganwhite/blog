package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) AdminLogin() {
	this.Data["title"] = "Logan's blog"
	this.TplName = "admin/login.html"
	this.Render()
}

func (this *AdminController) SubmitLogin() {
	username := this.GetString("u_name")
	password := this.GetString("u_passwd")
}
