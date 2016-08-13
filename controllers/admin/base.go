package admin

import (
	"github.com/astaxie/beego"
	"strings"
	"blog/models"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	userid int64
	username string
}

// 获取用户ip地址
func (this *baseController) GetClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 显示错误提示
func (this *baseController) ShowMsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = "admin/layout.html"
	this.TplName = "admin/showmsg.html"
	this.Render()
	this.StopRun()
}

// 显示错误提示
func (this *baseController) ShowSuccess(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = "admin/layout.html"
	this.TplName = "admin/showsuccess.html"
	this.Render()
	this.StopRun()
}

// 获取用户信息
func (this *baseController) GetUserInfo(username string) (*models.User, error) {
	o := orm.NewOrm()
	user := new(models.User)

	qs := o.QueryTable("user")
	err := qs.Filter("username", username).One(user)
	if err != nil {
		return nil, err
	}
	//err = qs.Filter("password", password).One(user)
	//if err != nil {
	//	return nil, err
	//}
	return user, err
}


