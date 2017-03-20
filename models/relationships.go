package models

import (
	"github.com/astaxie/beego/orm"
)

type Relationships struct {
	cid uint64
	mid uint64
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Relationships))
}
