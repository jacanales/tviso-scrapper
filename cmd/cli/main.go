package main

import (
	"tviso-scrapper/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
