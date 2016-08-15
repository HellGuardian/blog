package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct  {
	Id int
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

// 添加文章
func (a *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

// 获取文章
func (a *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

// 更新文章
func (a *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

// 删除文章
func (a *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(a); err != nil {
		return err
	}
	return nil
}

// 查询文章
func (a *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}