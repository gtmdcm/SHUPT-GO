package main

import (
	_ "SHUPT-GO/models"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.RunSyncdb("default", false, false)
}
