package models

import (
	"github.com/astaxie/beego/orm"
)

type Options struct {
	name        string
	value       string
	description string
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Options))
}
