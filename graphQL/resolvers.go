package graphQL

import (
	"SHUPT-GO/models"
	"github.com/astaxie/beego/orm"
	"golang.org/x/net/context"
)

type Resolver struct{}

func (r *Resolver) GetUser(ctx context.Context, args struct{ ID int32 }) (*models.User, error) {
	user := models.User{Id: uint64(args.ID)}
	orm_ := orm.NewOrm()
	orm_.Read(&user)
	return &user, nil
}

func (r *Resolver) GetPost(ctx context.Context, args struct{ ID int32 }) (*models.Post, error) {
	post := models.Post{Id: uint64(args.ID)}
	orm_ := orm.NewOrm()
	orm_.Read(&post)
	return &post, nil
}

func (r *Resolver) GetComment(ctx context.Context, args struct{ ID int32 }) (*models.Comment, error) {
	comment := models.Comment{Id: uint64(args.ID)}
	orm_ := orm.NewOrm()
	orm_.Read(comment)
	return &comment, nil
}

type UserInput struct {
	Nickname string
	CardId   string
}

func (r *Resolver) AddUser(ctx context.Context, args struct{ User UserInput }) (*models.User, error) {
	user := models.User{NickName: args.User.Nickname, CardId: args.User.CardId}
	orm_ := orm.NewOrm()
	orm_.Insert(&user)
	return &user, nil
}

type PostInput struct {
	Title     string
	Content   string
	CreatorId int32
}

func (r *Resolver) AddPost(ctx context.Context, args struct{ Post PostInput }) (*models.Post, error) {
	orm_ := orm.NewOrm()
	theCreator := models.User{Id: uint64(args.Post.CreatorId)}
	orm_.Read(&theCreator)
	post := models.Post{Title: args.Post.Title, Content: args.Post.Content, Creator: &theCreator}
	orm_.Insert(&post)
	return &post, nil
}

type CommentInput struct {
	Content        string
	CreatorId      int32
	BelongToPostId int32
}

func (r *Resolver) AddComment(ctx context.Context, args struct{ Comment CommentInput }) (*models.Comment, error) {
	orm_ := orm.NewOrm()
	theCreator := models.User{Id: uint64(args.Comment.CreatorId)}
	orm_.Read(&theCreator)
	thePost := models.Post{Id: uint64(args.Comment.BelongToPostId)}
	orm_.Read(&thePost)
	comment := models.Comment{
		Content:      args.Comment.Content,
		Creator:      &theCreator,
		BelongToPost: &thePost,
	}
	orm_.Insert(&comment)
	return &comment, nil
}
