package front

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
	this.Layout = "front/layout.html"
	this.TplName = "front/index.html"
}
