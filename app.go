package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/babenko/simple-go-service/service"
	"os"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Starting")
	router.
	Path("/upload").
		Methods("POST").
		HandlerFunc(service.UploadFile)
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func applyArgs() {
	port := os.Args[0]

}
