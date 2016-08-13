package admin

import (
	"blog/models"
	"strings"
	"github.com/astaxie/beego/validation"
	"blog/util"
	"time"
)

type UserController struct {
	baseController
}

// 用户列表
func (this *UserController) List() {
	this.TplName = "admin/userlist.html"
	this.Layout = "admin/layout.html"
	var list []*models.User
	var details models.User
	id, _ := this.GetInt("details")
	var user models.User
	count, _ := user.Query().Count()
	if count > 0 {
		user.Query().OrderBy("-id").All(&list)
		user.Query().Filter("id", id).One(&details)
	}

	this.Data["list"] = list
	this.Data["count"] = count
	this.Data["details"] = details

}

// 添加用户
func (this *UserController) Add() {
	this.TplName = "admin/useradd.html"
	this.Layout = "admin/layout.html"

	user := make(map[string]string)
	errmsg := make(map[string]string)
	if this.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		password1 := strings.TrimSpace(this.GetString("password1"))
		sex := strings.TrimSpace(this.GetString("sex"))
		email := strings.TrimSpace(this.GetString("email"))
		active, _ := this.GetInt("active")

		user["username"] = username
		user["password"] = password
		user["password1"] = password1
		user["email"] = email
		user["sex"] = sex

		valid := validation.Validation{}

		if v := valid.Required(username, "username"); !v.Ok {
			errmsg["username"] = "请输入用户名。"
		} else if v := valid.MaxSize(username, 25, "username"); !v.Ok {
			errmsg["username"] = "用户名长度不能超过25个字符!"
		}

		if v := valid.Required(password, "password"); !v.Ok {
			errmsg["password"] = "请输入密码"
		}

		if v:= valid.Required(password1, "password1"); !v.Ok {
			errmsg["password1"] = "请再次输入密码"
		} else if password != password1 {
			errmsg["password1"] = "两次输入密码不一致"
		}

		if active > 0 {
			active = 1
		} else {
			active = 0
		}

		if len(errmsg) == 0 {
			var user models.User
			user.UserName = username
			user.PassWord = util.Md5([]byte(password))
			user.Sex = sex
			user.Email = email
			user.Active = int(active)
			user.LoginCount = 1
			user.LastIp = this.GetClientIp()
			user.LastLogin = time.Now()
			user.CreateTime = time.Now()
			user.UpdateTime = time.Now()

			if err := user.Insert(); err != nil {
				this.ShowMsg(err.Error())
			}
			this.Redirect("/admin/user/list", 302)
		}

	}

	this.Data["user"] = user
	this.Data["errmsg"] = errmsg
}

// 编辑用户
func (this *UserController) Edit() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/useredit.html"
	id, _ := this.GetInt("id")
	user := models.User{Id: id}
	if err := user.Read(); err != nil {
		this.ShowMsg("用户不存在")
	}

	ErrMsg := make(map[string]string)

	if this.Ctx.Request.Method == "POST" {
		password := strings.TrimSpace(this.GetString("password"))
		password1 := strings.TrimSpace(this.GetString("password1"))
		sex := strings.TrimSpace(this.GetString("sex"))
		email := strings.TrimSpace(this.GetString("email"))
		active, _ := this.GetInt("active")
		valid := validation.Validation{}

		if password != ""  {
			if v := valid.Required(password1, "password1"); !v.Ok {
				ErrMsg["password1"] = "请再次输入密码"
			} else if password != password1 {
				ErrMsg["password1"] = "两次输入密码不一致"
			} else {
				user.PassWord = util.Md5([]byte(password))
			}
		}
		user.Sex = sex
		user.Email = email
		user.UpdateTime = time.Now()

		if active > 0 {
			user.Active = 1
		} else {
			user.Active = 0
		}

		if len(ErrMsg) == 0 {
			user.Update()
			this.Redirect("/admin/user/list", 302)
		}
	}

	this.Data["user"] = user
	this.Data["errmsg"] = ErrMsg
}

// 删除用户
func (this *UserController) Delete() {
	id, _ := this.GetInt("id")
	if id == 1 {
		this.ShowMsg("不能删除 ID 为1的用户")
	}
	user := models.User{Id: id}
	if user.Read() == nil {
		user.Delete()
		this.ShowSuccess("删除成功!")
	}

	this.Redirect("/admin/user/list", 302)
}

// 获取用户详情
/*func (this *UserController) Details() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/userlist.html"
	var details models.User
	id, _ := this.GetInt("details")
	var user models.User
	count, _ := user.Query().Count()
	if count > 0 {
		user.Query().Filter("id", id).One(&details)
	}

	this.Data["details"] = details
}*/