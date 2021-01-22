package main

import (
	"github.com/BryanKMorrow/aqua-exec-detection/src/webhooksrv"
	"log"
	"os"
)

func main() {
	log.Println("Using the environment variables")
	fatal := checkEnv()
	if fatal {
		log.Fatalln("Environment variables not set, stopping aqua-events-go")
	}
	s := webhooksrv.NewServer()
	s.Start()
}

func checkEnv() bool {
	fatal := false

	url := os.Getenv("AQUA_URL")
	if url == "" {
		log.Println("Please set the AQUA_URL environment variable")
		fatal = true
	}
	user := os.Getenv("AQUA_USER")
	if user == "" {
		log.Println("Please set the AQUA_USER environment variable")
	}
	password := os.Getenv("AQUA_PASSWORD")
	if password == "" {
		log.Println("Please set the AQUA_PASSWORD environment variable")
	}
	return fatal
}
