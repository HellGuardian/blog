package admin

import (
	"github.com/astaxie/beego"
)

type TestController struct  {
	baseController
}

func (this *TestController) Get() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/test.html"

	user, err := this.GetUserInfo("MaChao")
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["user"] = user
}
