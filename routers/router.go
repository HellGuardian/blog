package routers

import (
	"github.com/astaxie/beego"
	"blog/controllers/admin"
	"blog/controllers/front"
)

func init() {
	// 后台管理
    beego.Router("/admin", &admin.IndexController{})
	beego.Router("/admin/login", &admin.LoginController{}, "*:Login")

	// 用户管理
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")

	beego.Router("/admin/category/list", &admin.CategoryController{}, "*:List")
	beego.Router("/admin/category/add", &admin.CategoryController{}, "*:Add")
	beego.Router("/admin/category/edit", &admin.CategoryController{}, "*:Edit")


	beego.Router("/admin/test", &admin.TestController{})

	// 前台管理
	beego.Router("/", &front.IndexController{})
}
