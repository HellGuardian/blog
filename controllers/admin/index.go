package admin

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"

}