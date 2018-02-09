package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := NewClient(
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_SECRET"),
	)

	for {
		err = Monitor(client, os.Getenv("USER_ID"))

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Duration(1) * time.Minute)
	}
}
