package models

import (
	"github.com/astaxie/beego/orm"
)

type Comments struct {
	coid         uint64
	cid          uint64
	created      uint64 //unix timestap
	author       string
	author_id    uint64
	owner_id     uint64
	mail         string
	url          string
	ip           string
	agent        string
	content      string
	comment_type string
	status       string
	parent       uint64
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Comments))
}
