package auth

type User struct {
	Id           int64  // 用户ID
	Username     string `orm:"unique"` // 用户名
	PasswordHash string // 用户密码的哈希值
	Mail         string `orm:"unique"` // 用户的邮箱
}
