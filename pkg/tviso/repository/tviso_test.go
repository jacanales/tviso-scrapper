package repository_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/repository"
	"tviso-scrapper/pkg/tviso/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ErrClientRequest = errors.New("client error")

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

func TestTvisoAPI_GetUserCollection_UnmarshalErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	json, err := ioutil.ReadFile("stubs/user_collection_invalid.json")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(json)
	}))

	cfg.APIAddr = server.URL

	repo := repository.NewTvisoAPI(http.DefaultClient, cfg)

	collection, err := repo.GetUserCollection()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unmarshal error")
	assert.Empty(t, collection)
}

func TestTvisoAPI_GetUserCollection_RetryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	cfg.APIAddr = server.URL

	repo := repository.NewTvisoAPI(http.DefaultClient, cfg)

	collection, err := repo.GetUserCollection()

	assert.EqualError(t, err, tviso.ErrRequestError.Error())
	assert.Empty(t, collection)
}

func TestTvisoAPI_GetUserCollection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	json, err := ioutil.ReadFile("stubs/user_collection_empty.json")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(json)
	}))

	cfg.APIAddr = server.URL

	repo := repository.NewTvisoAPI(http.DefaultClient, cfg)

	collection, err := repo.GetUserCollection()

	assert.NoError(t, err)
	assert.Empty(t, collection)
}

func TestTvisoAPI_GetMediaInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	json, err := ioutil.ReadFile("stubs/media_file_movie.json")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(json)
	}))

	cfg.APIAddr = server.URL

	repo := repository.NewTvisoAPI(http.DefaultClient, cfg)

	media := tviso.Media{
		ID:        3864,
		MediaType: tviso.MoviesMediaType,
	}

	err = repo.GetMediaInfo(&media)

	assert.NoError(t, err)
	assert.Equal(t, media.IMDB, "tt0808510")
	assert.Equal(t, tviso.Pending.String(), media.StatusMedia)
}

func TestTvisoAPI_GetMediaInfo_DoRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	cli := mocks.NewMockHTTPClient(ctrl)
	cli.EXPECT().Do(gomock.Any()).Return(nil, ErrClientRequest)

	repo := repository.NewTvisoAPI(cli, cfg)

	media := tviso.Media{
		ID:        3864,
		MediaType: tviso.MoviesMediaType,
	}

	err := repo.GetMediaInfo(&media)

	assert.EqualError(t, err, fmt.Sprintf("request error: %v", ErrClientRequest.Error()))
}

func TestTvisoAPI_GetMediaInfo_UnmarshalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := repository.Config{}

	json, err := ioutil.ReadFile("stubs/media_file_invalid.json")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(json)
	}))

	cfg.APIAddr = server.URL

	repo := repository.NewTvisoAPI(http.DefaultClient, cfg)

	media := tviso.Media{
		ID:        3864,
		MediaType: tviso.MoviesMediaType,
	}

	err = repo.GetMediaInfo(&media)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unmarshal error")
}
