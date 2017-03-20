package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	dbname := beego.AppConfig.String("mysqldb")
	dburl := beego.AppConfig.String("mysqlurl")
	dbuser := beego.AppConfig.String("mysqluser")
	dbpasswd := beego.AppConfig.String("mysqlpass")

	dbMaxIdle := beego.AppConfig.int("mysqlmaxidle")
	dbMaxConn := beego.AppConfig.int("mysqlmaxconn")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpasswd+dburl+"@/"+dbname+"?charset=utf8", dbMaxIdle, dbMaxConn)
}
