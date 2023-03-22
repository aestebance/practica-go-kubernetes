package main

import (
	"fmt"
	"github.com/aestebance/practica-go-kubernetes/lib/client"
	"io"
	"log"
	"os"
	"time"
)

var (
	sleepTime string = "1s"
	serverURL string = "http://localhost:8000"
)

func main() {
	setupClient := client.SetupClient()

	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime = os.Getenv("SLEEP_TIME")
	}

	if len(os.Getenv("SERVER_URL")) != 0 {
		serverURL = os.Getenv("SERVER_URL")
	}

	endpoint := os.Getenv("ENDPOINT")

	for {
		accessPoint := fmt.Sprint(serverURL, endpoint)
		r, err := setupClient.Get(accessPoint)
		if err != nil {
			log.Panicln(err.Error())
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(r.Body)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Panicln(err.Error())
		}

		fmt.Printf("%s\n", body)

		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)
	}
}
