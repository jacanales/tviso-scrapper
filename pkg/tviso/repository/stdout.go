package repository

import (
	"fmt"

	"tviso-scrapper/pkg/tviso"
)

type StdOut struct{}

func NewStdOut() tviso.WriteRepository {
	return StdOut{}
}

func (r StdOut) StoreCollection(collection []tviso.Media) error {
	i := 0

	for _, c := range collection {
		i++

		fmt.Println(c.Name)
	}

	fmt.Printf("Total processed: %v\n", i)

	return nil
}
