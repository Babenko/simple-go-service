package service

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

var MAX_FILE_SIZE int64 = 2<<20
var STORAGE_DIRECTORY = "/tmp"

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START READ")
	er := r.ParseMultipartForm(MAX_FILE_SIZE)
	if er != nil {
		panic(er)
	}
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(STORAGE_DIRECTORY + "/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
