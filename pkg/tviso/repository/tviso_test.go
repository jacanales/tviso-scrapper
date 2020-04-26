package repository_test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/repository"
	"tviso-scrapper/pkg/tviso/repository/mocks"
)

func TestNewHTTPClient(t *testing.T) {
	cli := repository.NewHTTPClient()

	assert.Implements(t, (*repository.HTTPClient)(nil), cli)
	assert.IsType(t, &http.Client{}, cli)
}


func TestNewTvisoAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cli := mocks.NewMockHTTPClient(ctrl)

	repo := repository.NewTvisoAPI(cli)

	assert.Implements(t, (*tviso.ReadRepository)(nil), repo)
	assert.IsType(t, repository.TvisoAPI{}, repo)
}