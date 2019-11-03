package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	resp "github.com/salihkemaloglu/gignoxqc-beta-001/helpers"
	serv "github.com/salihkemaloglu/gignoxqc-beta-001/services"
)

//LoginController ...
func LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("QC service is working for LoginController... \n")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	response, responseType := serv.LoginService(w, r)
	resp.RespondWithJSON(w, http.StatusOK, map[string]string{"result": response, "status": strconv.FormatBool(responseType)})
}
