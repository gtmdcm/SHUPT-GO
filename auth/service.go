package auth

import (
	"encoding/json"
	"net/http"
)

func authHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form["username"][0]
	password := request.Form["password"][0]
	user, err := validateUser(username, password)
	if user == nil || err != nil {
		return
	}
	token, err := makeTokenJson(*user)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(token)
}

type CreateResultJson struct {
	Result bool `json:"result"`
}

func createUserHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form["username"][0]
	password := request.Form["password"][0]
	mail := request.Form["mail"][0]
	// todo: 增加密码、邮箱验证逻辑
	_, err := createUser(username, password, mail)
	result := CreateResultJson{Result: err == nil}
	writer.Header().Set("Content-Type", "application/json")
	resultAsByte, _ := json.Marshal(result)
	writer.Write(resultAsByte)
}

func StartService() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/create-user", createUserHandler)
	http.ListenAndServe(":8080", nil)
}
