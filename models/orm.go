package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"os"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", os.Getenv("CONNSTR"))
}
