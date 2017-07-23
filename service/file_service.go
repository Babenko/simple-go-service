package service

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

var MAX_FILE_SIZE int64 = 2<<20
var STORAGE_DIRECTORY = "/tmp"
var PARAM_NAME = "uploadfile"

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START READ")
	er := r.ParseMultipartForm(MAX_FILE_SIZE)
	if er != nil {
		panic(er)
	}
	file, handler, err := r.FormFile(PARAM_NAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	fmt.Print(handler.Filename)
	f, err := os.OpenFile(STORAGE_DIRECTORY + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	er := r.ParseMultipartForm(MAX_FILE_SIZE)
	if er != nil {
		panic(er)
	}
	w.Header().Set("Location", r.FormValue(PARAM_NAME + ".path"))
	w.Write([]byte(r.FormValue(PARAM_NAME + ".path")))
	w.WriteHeader(http.StatusCreated)
}
