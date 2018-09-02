package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int
	CardId   string
	NickName string
	Post     []*Post `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User))
}