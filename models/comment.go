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

func (c *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(c)
}