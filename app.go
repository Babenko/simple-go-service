package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/babenko/simple-go-service/service"
	"os"
	"strconv"
)

var PORT = "8080"

func main() {
	log.Println("Start")

	applyArgs()

	router := mux.NewRouter()

	router.
	Path("/upload").
		Methods("POST").
		HandlerFunc(service.UploadFile)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

func applyArgs() {
	log.Print(os.Args[0] + " | " + os.Args[1] + " | " + strconv.Itoa(len(os.Args)))
	if len(os.Args) == 2 && os.Args[1] != "" {
		service.STORAGE_DIRECTORY = os.Args[1]
	}
	log.Print("Files store at " + service.STORAGE_DIRECTORY)

	if len(os.Args) == 3 && os.Args[2] != "" {
		size, err := strconv.ParseInt(os.Args[2], 10, 0)
		if err != nil {
			log.Panic(err)
		} else {
			service.MAX_FILE_SIZE = size
		}
	}

	log.Print("Max file size is " + strconv.FormatInt(service.MAX_FILE_SIZE / 1000000, 10) + " mb.")

	if len(os.Args) == 4 && os.Args[3] != "" {
		PORT = os.Args[3]
	}
	log.Print("Listen on port " + PORT)

	if len(os.Args) == 5 && os.Args[4] != "" {
		service.PARAM_NAME = os.Args[4]
	}
	log.Print("Param name is " + service.PARAM_NAME)
}
