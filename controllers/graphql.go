package controllers

import (
	"SHUPT-GO/graphQL"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/context"
)

func GraphQLController(context *context.Context) {
	context.Output.ContentType("json")
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(bytes.NewBuffer(context.Input.RequestBody)).Decode(&params); err != nil {
		return
	}
	response := graphQL.Schema.Exec(context.Request.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return
	}
	context.Output.Body(responseJSON)
}
