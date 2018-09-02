package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=test password=test dbname=test host=127.0.0.1 sslmode=disable")
}

func SyncDB() {
	orm.RunSyncdb("default", false, false)
}
