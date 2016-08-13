package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id int64
	Title	string
	CreateTime	time.Time	`orm:"index"`
	Views	int64	`orm:"index"`
	ArticleCount	int64
}

func init() {
	orm.RegisterModel(new(Category))
}
