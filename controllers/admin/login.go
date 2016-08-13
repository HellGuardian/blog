package admin

import (
	"strings"
	"blog/util"
	"strconv"
)

type LoginController struct  {
	baseController
}

// 登录
func (this *LoginController) Login() {
	this.TplName = "admin/login.html"
	if this.GetString("dosubmit") == "yes" {
		username := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")

		user, err := this.GetUserInfo(username)
		if err != nil {
			this.Data["ErrMsg"] = "用户不存在"
			return
		}

		if user.Active == 0 {
			this.Data["ErrMsg"] = "帐号未激活,请联系管理员"
			return
		} else if user.UserName != username {
			this.Data["ErrMsg"] = "用户名错误"
			return
		} else if user.PassWord != util.Md5([]byte(password)) {
			this.Data["ErrMsg"] = "密码错误"
			return
		} else {
			user.LoginCount += 1
			user.LastIp = this.GetClientIp()
			user.Status = 1
			user.Update()
			AuthKey := util.Md5([]byte(this.GetClientIp() + "|" + user.PassWord))
			if remember == "yes" {
				this.Ctx.SetCookie("auth", strconv.Itoa(user.Id) + "|" + AuthKey, 7*86400)
			} else {
				this.Ctx.SetCookie("auth", strconv.Itoa(user.Id) + "|" + AuthKey)
			}
		}
		this.Redirect("/admin", 302)
	}
}

// 退出登录
func (this *LoginController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/admin/login", 302)
}