package main

import (
	"SHUPT-GO/auth"
	_ "SHUPT-GO/auth"
	_ "SHUPT-GO/models"
	_ "SHUPT-GO/routers"
	"github.com/astaxie/beego"
)

func main() {
	auth.StartService()
	beego.Run()
}
