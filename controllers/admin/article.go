package admin

import (
	"blog/models"
	"strings"
	"time"
)

type ArticleController struct  {
	baseController
}

// 文章列表
func (this *ArticleController) List() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article.html"

	var list []*models.Article
	var article models.Article
	count, _ := article.Query().Count()
	if count > 0 {
		article.Query().OrderBy("-id").All(&list)
	}
	this.Data["list"] = list
}

// 添加文章
func (this *ArticleController) Add() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/articleadd.html"

	var list []*models.Category
	var category models.Category
	count, _ := category.Query().Count()
	if count > 0 {
		category.Query().All(&list)
	}
	this.Data["list"] = list

	if this.Ctx.Request.Method == "POST" {
		title := strings.TrimSpace(this.GetString("title"))
		cName := strings.TrimSpace(this.GetString("category"))
		content := strings.TrimSpace(this.GetString("content"))
		tag := strings.TrimSpace(this.GetString("tag"))
		status, _ := this.GetInt("status")

		//sessionID := this.Ctx.GetCookie("auth")

		var article models.Article
		article.Title = title
		article.Category = cName
		article.Content = content
		article.Tag = tag
		article.Status = status
		article.ViewCount = 1
		article.CreateTime = time.Now()

	}
}

// 编辑文章
func (this *ArticleController) Edit() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/articlerdit.html"
}

// 删除文章
func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	article := models.Article{Id: id}
	if article.Read() == nil {
		article.Delete()
		this.ShowSuccess("文章删除成功。")
	}
	this.Redirect("/admin/Article/list", 302)
}