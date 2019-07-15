package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"
	repo "github.com/salihkemaloglu/gignoxqc-beta-001/repositories"
	serv "github.com/salihkemaloglu/gignoxqc-beta-001/services"
	"github.com/spf13/pflag"
	"goji.io"
	"goji.io/pat"
)

//SayHello ...
func SayHello(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("QC service is working for SayHello...Received rpc from client \n")
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "Hello QC service is working..."})
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("QC service is working for Login...Received rpc from client \n")
	resp, err := serv.LoginService(w, r)
	respondWithJSON(w, http.StatusOK, map[string]string{"result": resp, "status": strconv.FormatBool(err)})
}

//UploadFile ...
func UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("QC service is working for UploadFile...Received rpc from client \n")
	err := serv.UploadFileService(w, r)
	if err != nil {
		respondWithError(w, http.StatusUnsupportedMediaType, err.Error())
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
func main() {
	pflag.Parse()
	fmt.Println("QC Service is Starting...")
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello"), SayHello)
	mux.HandleFunc(pat.Post("/login"), Login)
	mux.HandleFunc(pat.Post("/uploadfile"), UploadFile)
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%v", 8904),
		Handler: cors.AllowAll().Handler(mux),
	}
	fmt.Println("Mongodb Service Started")
	if confErr := repo.LoadConfiguration("dev"); confErr != "ok" {
		fmt.Println(confErr)
	}
	fmt.Printf("server started as http and listen to port: %v \n", 8904)
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("failed starting http server: %v", err.Error())
	}
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
