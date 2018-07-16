package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func authHandler(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	var authInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(body, &authInfo)
	user, err := validateUser(authInfo.Username, authInfo.Password)
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
	body, err := ioutil.ReadAll(request.Body)
	if err == nil {
		var createInfo struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Mail     string `json:"mail"`
		}
		err = json.Unmarshal(body, &createInfo)
		if err == nil {
			// todo: 增加密码、邮箱验证逻辑
			_, err = createUser(createInfo.Username, createInfo.Password, createInfo.Mail)
		}
	}
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
