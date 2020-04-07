// +build wireinject

package cmd

import "fmt"

func Start() error {
	fmt.Print("app started")

	return nil
}