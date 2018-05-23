package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"wshhz.com/blog/models"
	_ "wshhz.com/blog/routers"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
