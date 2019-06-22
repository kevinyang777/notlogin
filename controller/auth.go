package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kevinyang777/loginer/model"
	"github.com/kevinyang777/loginer/utils/response"
)

var CreateAccount = func(data *gin.Context) {
	account := &model.Account{}
	err := data.BindJSON(&account)
	fmt.Println("error", err)
	if err != nil {
		response.FailResponse(data)
		return
	}
	fmt.Println("name", account.Name)
	response.SuccessResponse(data, `{"asu":"gblk"}`)
	return
}
