package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/babenko/simple-go-service/service"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":80", router))
}
