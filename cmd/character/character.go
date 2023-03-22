package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	character string = "..."
	sleepTime string = "1s"
)

func main() {
	for {
		hostname, err := os.Hostname()
		if err != nil {
			log.Panicln(err.Error())
		}

		if len(os.Getenv("CHARACTER")) != 0 {
			character = os.Getenv("CHARACTER")
		}
		fmt.Printf("Hostname: %s - érase una vez %s\n", hostname, character)

		if len(os.Getenv("SLEEP_TIME")) != 0 {
			sleepTime = os.Getenv("SLEEP_TIME")
		}

		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)
	}
}
