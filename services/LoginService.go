package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	helper "github.com/salihkemaloglu/gignoxqc-beta-001/helpers"
	inter "github.com/salihkemaloglu/gignoxqc-beta-001/interfaces"
	repo "github.com/salihkemaloglu/gignoxqc-beta-001/repositories"
)

//LoginService ...
func LoginService(w http.ResponseWriter, r *http.Request) (string, bool) {
	defer r.Body.Close()
	var user repo.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return "Invalid request payload", false
	}

	var op inter.IUserRepository = user
	_, err := op.Login()
	if err != nil {
		return "Invalid user information", false
	}
	tokenRes, tokenErr := helper.CreateTokenEndpointService(user)
	if tokenErr != nil {
		return fmt.Sprintf("Token creation error "+": %v", tokenErr.Error()), false
	}

	return tokenRes, true

}
