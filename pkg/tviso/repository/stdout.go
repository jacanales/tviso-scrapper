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
	for i := range collection {
		fmt.Println(collection[i].Name)
	}

	fmt.Printf("Total processed: %v\n", len(collection))

	return nil
}
