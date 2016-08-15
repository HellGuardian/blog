package admin

import (
	"github.com/astaxie/beego"
)

type TestController struct  {
	baseController
}

type MainController struct  {
	beego.Controller
}

func (this *TestController) Get() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/test.html"

	user, err := this.GetUserInfo("MaChao")
	if err != nil {
		beego.Error(err)
		return
	}
	//sessionID := this.GetSession("auth")
	s := this.Ctx.GetCookie("auth")

	this.Data["auth"] = s
	this.Data["user"] = user
}
