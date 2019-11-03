package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	ctrl "github.com/salihkemaloglu/gignoxqc-beta-001/controllers"
	repo "github.com/salihkemaloglu/gignoxqc-beta-001/repositories"
	"github.com/spf13/pflag"
	"goji.io"
	"goji.io/pat"
)

func main() {
	pflag.Parse()
	fmt.Println("QC Service is Starting...")
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello"), ctrl.SayHelloController)
	mux.HandleFunc(pat.Post("/login"), ctrl.LoginController)
	mux.HandleFunc(pat.Post("/uploadfile"), ctrl.UploadFileController)
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%v", 8904),
		Handler: cors.AllowAll().Handler(mux),
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	fmt.Println("Mongodb Service Started")
	if confErr := repo.LoadConfiguration("dev"); confErr != "ok" {
		fmt.Println(confErr)
	}
	fmt.Printf("server started as http and listen to port: %v \n", 8904)
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("failed starting http server: %v", err.Error())
	}
}
