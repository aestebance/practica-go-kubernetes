package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aestebance/practica-go-kubernetes/lib/utils"
)

var (
	animal string = "cat"
)

func HealthResponse(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, r, map[string]string{"status": "OK"})
	return
}

func EchoResponse(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Panicln("Error al obtener el hostname.")
	}

	if len(os.Getenv("ANIMAL")) != 0 {
		animal = os.Getenv("ANIMAL")
	}
	message := fmt.Sprintf("Esto es un %s", animal)
	utils.JSONResponse(w, r, map[string]string{"hostname": hostname, "message": message})
	return
}