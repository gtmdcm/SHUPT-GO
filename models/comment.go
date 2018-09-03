package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
)

type Comment struct {
	Id           uint64
	Content      string
	Creator      *User `orm:"rel(fk);on_delete(cascade)"`
	BelongToPost *Post `orm:"rel(fk);on_delete(cascade)"`
}

func (comment Comment) ID(ctx context.Context) graphql.ID {
	id := graphql.ID(comment.Id)
	return id
}
func (comment Comment) CONTENT(ctx context.Context) string {
	return comment.Content
}
func (comment Comment) CREATOR(ctx context.Context) *User {
	return comment.Creator
}
func (comment Comment) BELONGTOPOST(ctx context.Context) *Post {
	return comment.BelongToPost
}
func init() {
	orm.RegisterModel(new(Comment))
}
