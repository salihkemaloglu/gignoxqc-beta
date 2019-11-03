package services

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

//UploadFileService ...
func UploadFileService(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// copy example
	absPath, _ := filepath.Abs(handler.Filename)
	f, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)
}
