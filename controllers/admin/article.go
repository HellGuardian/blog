package admin

import (

)

type ArticleController struct  {
	baseController
}

// 文章列表
func (this *ArticleController) List() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article.html"
}

// 添加文章
func (this *ArticleController) Add() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/articleadd.html"
}

// 编辑文章
func (this *ArticleController) Edit() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/articlerdit.html"
}