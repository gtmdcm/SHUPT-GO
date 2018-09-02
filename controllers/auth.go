package controllers

import (
	"SHUPT-GO/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
)

func AuthHandler(ctx *context.Context) {
	userInfo := struct {
		CardId   string `json:"card_id"`
		Password string `json:"password"`
	}{"", ""}
	json.Unmarshal(ctx.Input.RequestBody, &userInfo)
	response, err := http.Post("https://www.shuhelper.cn/api/users/login/", "application/json",
		bytes.NewBuffer(ctx.Input.RequestBody))
	if err != nil {
		ctx.Output.Body([]byte("登录服务GG了……"))
		return
	}
	if response.StatusCode != 200 {
		ctx.Output.Body([]byte("您的登录信息有误"))
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	responseUserInfo := struct {
		Name string `json:"name"`
	}{}
	json.Unmarshal(body, &responseUserInfo)
	orm_ := orm.NewOrm()
	user := new(models.User)
	querySeter := orm_.QueryTable(user)
	err = querySeter.Filter("card_id", userInfo.CardId).One(user)
	if err == orm.ErrNoRows {
		user.CardId = userInfo.CardId
		user.NickName = responseUserInfo.Name
		orm_.Insert(user)
		ctx.Output.Body([]byte("注册成功"))
		return
	}
	ctx.Output.Body([]byte("登录成功"))
}
