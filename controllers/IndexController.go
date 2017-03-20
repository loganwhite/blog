package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) RedirectIndex() {
	this.Redirect("/static/", 302)
}
