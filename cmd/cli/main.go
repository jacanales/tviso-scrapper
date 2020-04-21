package main

import (
	"log"

	"github.com/joho/godotenv"

	"tviso-scrapper/pkg/cli"
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
