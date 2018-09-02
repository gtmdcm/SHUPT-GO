package models

import "github.com/astaxie/beego/orm"

type Comment struct {
	Id           int
	Content      string
	Creator      *User `orm:"rel(fk)"`
	BelongToPost *Post `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}
