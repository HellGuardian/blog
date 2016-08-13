package admin

import (
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

// 登录
func (this *AccountController) Login() {
	this.TplName = "admin/login.html"
}
