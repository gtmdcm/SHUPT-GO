package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
)

type Post struct {
	Id       uint64
	Title    string
	Content  string
	Creator  *User      `orm:"rel(fk);on_delete(cascade)"`
	Comments []*Comment `orm:"reverse(many)"`
}

func (post Post) ID(ctx context.Context) graphql.ID {
	id := graphql.ID(post.Id)
	return id
}

func (post Post) TITLE(ctx context.Context) string {
	return post.Title
}

func (post Post) CONTENT(ctx context.Context) string {
	return post.Content
}

func (post Post) CREATOR(ctx context.Context) *User {
	return post.Creator
}

func (post Post) COMMENTS(ctx context.Context) *[]*Comment {
	return &post.Comments
}

func init() {
	orm.RegisterModel(new(Post))
}
