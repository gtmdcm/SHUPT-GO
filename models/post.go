package models

import "github.com/astaxie/beego/orm"

type Post struct {
	Id       uint64
	Title    string
	Content  string
	Creator  *User      `orm:"rel(fk);on_delete(cascade)"`
	Comments []*Comment `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Post))
}
