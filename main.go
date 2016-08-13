package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"blog/models"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}

