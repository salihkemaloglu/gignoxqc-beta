package controllers

import (
	"fmt"
	"net/http"

	resp "github.com/salihkemaloglu/gignoxqc-beta-001/helpers"
	serv "github.com/salihkemaloglu/gignoxqc-beta-001/services"
)

//UploadFileController ...
func UploadFileController(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("QC service is working for UploadFileController... \n")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	serv.UploadFileService(w, r)
	resp.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
