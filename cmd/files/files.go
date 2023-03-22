package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	folderPath string = "/srv/files"
	sleepTime  int    = 5
)

func main() {
	for {
		hostname, err := os.Hostname()
		if err != nil {
			log.Panicln(err.Error())
		}

		files, err := os.ReadDir(folderPath)
		if err != nil {
			log.Panicln(err.Error())
		}

		countFiles := len(files)

		fmt.Printf("Hostname: %s, Total ficheros: %d\n", hostname, countFiles)

		name := fmt.Sprintf("%s/%s-%d", folderPath, hostname, time.Now().Unix())
		file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)

		if err != nil {
			log.Panicln(err)
		}
		file.Close()

		if len(os.Getenv("SLEEP_TIME")) != 0 {
			sleepTime, err = strconv.Atoi(os.Getenv("SLEEP_TIME"))
			if err != nil {
				log.Panicln(err)
			}
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
