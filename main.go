package untweet

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Load and Set settings variables from .env file
// and Monitor given user ID profile for changes
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

	rt, err := strconv.Atoi(os.Getenv("REFRESH_TIME"))

	if err != nil {
		log.Fatal("Can't convert string to integer")
	}

	for {
		err = Monitor(client, os.Getenv("USER_ID"))

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Duration(rt) * time.Minute)
	}
}
