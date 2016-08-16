package front

import (
	"github.com/astaxie/beego"
	"blog/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Layout = "front/layout.html"
	this.TplName = "front/index.html"
}

// 首页文章列表
func (this *IndexController) Index() {
	this.TplName = "front/index.html"
	this.Layout = "front/layout.html"

	var list []*models.Article
	var article models.Article
	count, _ := article.Query().Count()
	if count > 0 {
		article.Query().OrderBy("-id").All(&list)
	}
	this.Data["list"] = list
}

// 浏览文章
func (this *IndexController) View() {
	this.Layout = "front/layout.html"
	this.TplName = "front/viewsarticle.html"

	id, _ := this.GetInt("id")
	article := models.Article{Id: id}
	if err := article.Read(); err != nil {
		return
	}
	this.Data["article"] = article
}