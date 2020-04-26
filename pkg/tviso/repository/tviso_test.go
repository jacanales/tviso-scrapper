package repository_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/repository"
	"tviso-scrapper/pkg/tviso/repository/mocks"
)

var (
	ErrClientRequest = errors.New("client error")
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

	repo := repository.NewTvisoAPI(cli, repository.Config{})

	assert.Implements(t, (*tviso.ReadRepository)(nil), repo)
	assert.IsType(t, repository.TvisoAPI{}, repo)
}

func TestTvisoAPI_GetUserCollection_DoRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cli := mocks.NewMockHTTPClient(ctrl)
	cli.EXPECT().Do(gomock.Any()).Return(nil, ErrClientRequest)

	repo := repository.NewTvisoAPI(cli, repository.Config{})

	collection, err := repo.GetUserCollection()

	assert.EqualError(t, err, fmt.Sprintf("request error: %v", ErrClientRequest.Error()))
	assert.Empty(t, collection)
}

func TestTvisoAPI_GetUserCollection(t *testing.T) {
	t.Skip("pending")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	cli := mocks.NewMockHTTPClient(ctrl)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	cfg.APIAddr = server.URL

	repo := repository.NewTvisoAPI(cli, repository.Config{})

	collection, err := repo.GetUserCollection()

	assert.NoError(t, err)
	assert.Empty(t, collection)
}