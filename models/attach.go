package models

import (
	"github.com/astaxie/beego/orm"
)

type Attach struct {
	uid       uint64
	fname     string
	ftype     string
	fkey      string
	author_id uint64
	created   uint64 //unix timestap
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Attach))
}
