package models

import "github.com/astaxie/beego/orm"

type Comment struct {
	Id           uint64
	Content      string
	Creator      *User `orm:"rel(fk);on_delete(cascade)"`
	BelongToPost *Post `orm:"rel(fk);on_delete(cascade)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}
