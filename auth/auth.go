package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

func createUser(username string, password string, mail string) (id int64, err error) {
	ormObject := orm.NewOrm()
	theUser := new(User)
	theUser.Username = username
	theUser.PasswordHash = makePassword(password)
	theUser.Mail = mail
	id, err = ormObject.Insert(theUser)
	return id, err
}

func validateUser(username string, password string) (user *User, err error) {
	ormObject := orm.NewOrm()
	theUser := new(User)
	theUser.Username = username
	err = ormObject.Read(theUser, "Username")
	if err != nil {
		return nil, err
	} else {
		if checkPassword(password, *theUser) {
			return theUser, err
		}
		return nil, errors.New("invalid password for user" + username)
	}
}
