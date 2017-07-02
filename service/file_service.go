package service

import (
	"fmt"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START READ")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		panic(err)
	}
	fmt.Println("END READ")
	fmt.Println(handler.Filename)
	fmt.Println(file)
}
