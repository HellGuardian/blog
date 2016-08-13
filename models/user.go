package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	UserName string	`orm:"unique;size(25)"`
	PassWord string `orm:"size(32)"`
	Sex string
	Email string `orm:"size(50)"`
	Status int
	LoginCount int64
	LastIp string
	Active int
	ImgUrl string
	LastLogin time.Time `orm:"type(datetime)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

/*func GetOneUserInfo(id int) (*User, error) {
	o := orm.NewOrm()
	user := new(User)

	qs := o.QueryTable("user")
	err := qs.Filter("id", 1).One(&user)
	if err != nil {
		return nil, err
	}
	return user, err
} */