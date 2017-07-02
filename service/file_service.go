package service

import (
	"fmt"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START READ")
	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		panic(er)
	}
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		panic(err)
	}
	fmt.Println("END READ")
	fmt.Println(handler.Filename)
	fmt.Println(file)
}
