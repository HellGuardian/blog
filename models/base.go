package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	/*dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dbcon := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	*/
	//orm.RegisterDataBase("default", "mysql", "blog:linux@tcp(192.168.31.165:3306)/blog?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "blog:linux@tcp(192.168.191.2:3306)/blog?charset=utf8")
	orm.RunSyncdb("default", false, true)
}