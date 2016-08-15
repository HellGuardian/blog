package admin

import (
	"github.com/astaxie/beego"
	"os"
	"runtime"
	"blog/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"

	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["hostname"], _ = os.Hostname()
	this.Data["GoVersion"] = runtime.Version()
	this.Data["arch"] = runtime.GOARCH
	this.Data["os"] = runtime.GOOS
	this.Data["CpuNum"] = runtime.NumCPU()

	this.Data["UserNum"], _ = new(models.User).Query().Count()

}