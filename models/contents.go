package models

import (
	"github.com/astaxie/beego/orm"
)

type Contents struct {
	cid           uint64
	title         string
	slug          string
	created       uint64 //unix timestap
	modified      uint64 //unix timestap
	content       string
	author_id     uint64
	content_type  string
	status        string
	tags          string
	categories    string
	hits          uint
	comments_num  uint
	allow_comment uint
	allow_ping    uint
	allow_feed    uint
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Contents))
}
