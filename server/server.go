package server

import (
	"gin-ory-stack-demo/config"
	"log"
	"net/http"
)

func Init() {
	config := config.GetConfig()
	router := NewRouter()

	serverPort := config.GetString("server.port")
	address := "0.0.0.0:" + serverPort

	log.Printf("Server listening on http://%s/", address)

	if error := http.ListenAndServe(address, router); error != nil {
		log.Fatalf("There was an error with the http server: %v", error)
	}
}
