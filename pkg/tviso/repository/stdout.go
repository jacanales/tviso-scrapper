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
	for _, c := range collection {
		fmt.Println(c.Name)
	}

	return nil
}
