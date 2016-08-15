package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id int
	Title	string
	CreateTime	time.Time	`orm:"index"`
	UpdateTime time.Time `orm:"auto_now"`
	Views	int64	`orm:"index"`
	ArticleCount	int64
}

func init() {
	orm.RegisterModel(new(Category))
}

// 插入文章
func (c *Category) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

// 读取某行数据
func (c *Category) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

// 更新数据
func (c *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

// 删除记录
func (c *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

// 查询记录
func (c *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(c)
}