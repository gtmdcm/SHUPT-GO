package models

import "github.com/astaxie/beego/orm"

type Post struct {
	Id       int
	Title    string
	Content  string
	Creator  *User      `orm:"rel(fk)"`
	Comments []*Comment `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Post))
}
