package models

import (
	"SHUPT-GO/auth"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(auth.User))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=test password=test dbname=test host=127.0.0.1 sslmode=disable")
	orm.RunSyncdb("default", true, false)
}
