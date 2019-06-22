package controller

import (
	"net/http"

	base "github.com/kevinyang777/loginer/utils/response"
)

var CreateAccount = func(req http.ResponseWriter, res *http.Request) {
	// account := &model.Account{}
	// accData := `{"id":"123,"name":"jancuk"}`
	// resp := json.NewDecoder(accDat).Decode(account)
	base.Respond(req, base.Message(false, "Invalid request"))
}
