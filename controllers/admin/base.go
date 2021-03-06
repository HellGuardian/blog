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
	return user, err
}

// 根据分隔符 | 来分隔字符串
func split(s rune) bool {
	if s == '|' {
		return true
	}
	return false
}
func Separate(s string) string {
	sm := strings.FieldsFunc(s, split)
	return sm[0]
}

// 获取分类名称
/*func (this *baseController) GetCategoryInfo(title string) error {
	o := orm.NewOrm()
	category := new(models.Category)

	qs := o.QueryTable("category")
	err := qs.Filter("title", title).One(category)
	if category.Title == title {
		this.ShowMsg("分类名称已经存在,请重新输入。")
		return err
	}
	return err
} */