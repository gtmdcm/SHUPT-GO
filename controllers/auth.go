package controllers

import (
	"SHUPT-GO/models"
	"SHUPT-GO/tools/jwt"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
)

const loginBackend = "https://www.shuhelper.cn/api/users/login/"

type message struct {
	Success bool `json:"success"`
}

type failMessage struct {
	message
	Reason string `json:"reason"`
}

func newFailMessage(reason string) failMessage {
	return failMessage{
		message{false},
		reason,
	}
}

type successMessage struct {
	message
	NewUser bool   `json:"new_user"`
	Token   string `json:"token"`
}

func newSuccessMessage(newUser bool, userId uint64) successMessage {
	return successMessage{
		message{true},
		newUser,
		jwt.MakeToken(userId),
	}
}

type userInfo struct {
	CardId   string `json:"card_id"`
	Password string `json:"password"`
}

type simulateLoginResponse struct {
	Name string `json:"name"`
}

func AuthHandler(context *context.Context) {
	context.Output.ContentType("json")
	var userInfo userInfo
	json.Unmarshal(context.Input.RequestBody, &userInfo)
	response, err := http.Post(loginBackend, "application/json",
		bytes.NewBuffer(context.Input.RequestBody))
	retryCount := 0
	for err != nil && retryCount < 5 {
		response, err = http.Post(loginBackend, "application/json",
			bytes.NewBuffer(context.Input.RequestBody))
		retryCount++
	}
	if err != nil {
		msg, _ := json.Marshal(newFailMessage("登录服务GG了……"))
		context.Output.Body(msg)
		return
	}
	if response.StatusCode != 200 {
		msg, _ := json.Marshal(newFailMessage("您的信息有误"))
		context.Output.Body(msg)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var simulateLoginResponse simulateLoginResponse
	json.Unmarshal(body, &simulateLoginResponse)
	orm_ := orm.NewOrm()
	user := models.User{CardId: userInfo.CardId, NickName: simulateLoginResponse.Name}
	if created, id, _ := orm_.ReadOrCreate(&user, "card_id"); err == nil {
		msg, _ := json.Marshal(newSuccessMessage(created, uint64(id)))
		context.Output.Body(msg)
		return
	}
	panic("Should never reach this")
}
