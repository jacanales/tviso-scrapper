package repository_test

import (
	"testing"
	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/repository"

	"github.com/stretchr/testify/assert"
)

func TestNewMongoDBRepository(t *testing.T) {
	rp := repository.NewMongoDBRepository()

	collection := []tviso.Media{
		{
			Name: "testMedia",
		},
	}

	err := rp.StoreCollection(collection)

	assert.NoError(t, err)
}
