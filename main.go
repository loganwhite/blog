package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
