package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/babenko/simple-go-service/service"
	"os"
	"strconv"
)

var PORT = "8080"

func main() {
	router := mux.NewRouter()

	fmt.Println("Starting")
	router.
	Path("/upload").
		Methods("POST").
		HandlerFunc(service.UploadFile)
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(PORT, router))
}

func applyArgs() {
	if os.Args[0] != "" {
		service.STORAGE_DIRECTORY = os.Args[1]
	}
	log.Print("Files store at " + service.STORAGE_DIRECTORY)

	if os.Args[1] != "" {
		size, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err != nil {
			log.Panic(err)
		} else {
			service.MAX_FILE_SIZE = size
		}
	}

	log.Print("Max file size is " + strconv.FormatInt(service.MAX_FILE_SIZE / 1000000, 10) + " mb.")

	if os.Args[2] != "" {
		PORT = os.Args[0]
	}
	log.Print("Listen on port " + PORT)
}
