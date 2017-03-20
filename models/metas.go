package models

import (
	"github.com/astaxie/beego/orm"
)

type Metas struct {
	mid         uint64
	name        string
	slug        string
	meta_type   string
	description string
	sort        int
	parent      uint64
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Metas))
}
