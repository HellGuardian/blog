package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id int64
	Tid int64
	Name string
	Email string
	Context string `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
}

func init() {
	orm.RegisterModel(new(Comment))
}
