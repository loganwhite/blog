package models

import (
	"github.com/astaxie/beego/orm"
)

type Logs struct {
	id        uint64
	action    string
	data      string
	author_id uint64
	ip        string
	created   uint64 //unix timestap
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Logs))
}
