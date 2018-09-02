package main

import (
	_ "SHUPT-GO/models"
	_ "SHUPT-GO/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
