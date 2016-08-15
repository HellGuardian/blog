package admin

import (
	"strings"
	"time"
	"blog/models"
	"github.com/astaxie/beego/orm"
)

type CategoryController struct  {
	baseController
}

// 分类列表
func (this *CategoryController) List() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/category.html"
	var list []*models.Category
	var category models.Category
	count, _ := category.Query().Count()
	if count > 0 {
		category.Query().OrderBy("-id").All(&list)
	}
	this.Data["list"] = list
}

// 添加分类
func (this *CategoryController) Add() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/categoryadd.html"

	if this.Ctx.Request.Method == "POST" {
		title := strings.TrimSpace(this.GetString("cName"))
		o := orm.NewOrm()
		category := new(models.Category)

		qs := o.QueryTable("category")
		err := qs.Filter("title", title).One(category)
		if err == nil {
			this.ShowMsg("分类名称已经存在,请重新输入。")
			return
		} else {
			cate := &models.Category{Title: title, CreateTime: time.Now(), Views: 1}
			if _, err := o.Insert(cate); err != nil {
				this.ShowMsg(err.Error())
			}
		}
		this.Redirect("/admin/category/list", 302)
	}
}

// 编辑分类
func (this *CategoryController) Edit() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/categoryedit.html"

	id, _ := this.GetInt("id")
	cate := models.Category{Id: id}
	if err := cate.Read(); err != nil {
		return
	}
	this.Data["cate"] = cate

	if this.Ctx.Request.Method == "POST" {
		title := strings.TrimSpace(this.GetString("cName"))


		o := orm.NewOrm()
		category := new(models.Category)

		qs := o.QueryTable("category")
		err := qs.Filter("title", title).One(category)
		if err == nil {
			this.ShowMsg("分类名称已经存在,请重新输入。")
			return
		} else if cate.Title == title {
			this.ShowMsg("分类名称已经存在,请重新输入。")
			return
		} else {
			cate.Title = title
			cate.UpdateTime = time.Now()
			cate.Views += 1
			cate.Update()
			this.Redirect("/admin/category/list", 302)
		}
	}
}

func (this *CategoryController) Delete() {
	id, _ := this.GetInt("id")
	category := models.Category{Id: id}
	if category.Read() == nil {
		category.Delete()
		this.ShowSuccess("分类删除成功。")
	}
	this.Redirect("/admin/category/list", 302)
}
