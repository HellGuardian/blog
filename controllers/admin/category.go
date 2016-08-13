package admin

import (

)

type CategoryController struct  {
	baseController
}

// 分类列表
func (this *CategoryController) List() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/category.html"
}

// 添加分类
func (this *CategoryController) Add() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/categoryadd.html"
}

// 编辑分类
func (this *CategoryController) Edit() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/categoryedit.html"
}
