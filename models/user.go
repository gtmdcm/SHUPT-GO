package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
)

type User struct {
	Id       uint64
	CardId   string
	NickName string
	Post     []*Post `orm:"reverse(many)"`
}

func (user *User) ID(ctx context.Context) graphql.ID {
	id := graphql.ID(fmt.Sprint(user.Id))
	return id
}

func (user *User) CARDID(ctx context.Context) string {
	return user.CardId
}

func (user *User) NICKNAME(ctx context.Context) string {
	return user.NickName
}

func (user *User) POSTS(ctx context.Context) *[]*Post {
	return &user.Post
}

func init() {
	orm.RegisterModel(new(User))
}
