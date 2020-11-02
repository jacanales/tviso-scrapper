package main

import (
	"log"
	"tviso-scrapper/pkg/cli"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)

		return
	}

	if err := cli.Execute(); err != nil {
		log.Fatal(err)

		return
	}
}
