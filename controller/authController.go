package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kevinyang777/loginer/utils/response"
)

var CreateAccount = func(data *gin.Context) {
	// account := &model.Account{}
	var recievedData map[string]interface{}
	err := data.BindJSON(&recievedData)
	if err != nil {
		response.FailResponse(data)
		panic(err)
	}
	finalData, _ := json.Marshal(recievedData)
	response.SuccessResponse(data, finalData)
	return
}
