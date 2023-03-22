package main

import (
	"github.com/aestebance/practica-go-kubernetes/lib/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	addr string = ":8000"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", server.HealthResponse).Methods("GET")
	router.HandleFunc("/echo", server.EchoResponse).Methods("GET")

	//starting server
	log.Println("Server started")
	err := http.ListenAndServe(addr, router)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", addr, err)
	}
}
