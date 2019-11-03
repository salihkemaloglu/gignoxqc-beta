package controllers

import (
	"fmt"
	"net/http"

	resp "github.com/salihkemaloglu/gignoxqc-beta-001/helpers"
)

//SayHelloController ...
func SayHelloController(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("QC service is working for SayHello...Received rpc from client \n")
	resp.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "Hello QC service is working..."})
}
