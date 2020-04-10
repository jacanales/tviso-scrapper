package main

import (
	`fmt`

	"github.com/joho/godotenv"

	"tviso-scrapper/pkg/cli"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)

		return
	}

	if err := cli.Execute(); err != nil {
		fmt.Println(err)

		return
	}
}
