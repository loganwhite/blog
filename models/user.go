package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	uid         uint64
	username    string
	password    string
	email       string
	home_url    string
	screen_name string
	created     uint64 //unix timestap
	activated   bool
	logged      uint64 //unix timestap
	group_name  string
}

func init() {
	//orm.RegisterModel(new(User))
	orm.RegisterModelWithPrefix("t_", new(User))
}
