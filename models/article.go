package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct  {
	Id int64
	Title string `orm:"type(text)"`
	Content string `orm:"type(text)"`
	Status int `orm:"index"`
	Category string
	Tag string
	AuthorId int64 `orm:"index"`
	AuthorName string
	ViewCount int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now"`
	ReplyCount int64
	ReplyLastUserId int64
}

func init() {
	orm.RegisterModel(new(Article))
}
