package repository_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"tviso-scrapper/pkg/tviso/repository"
)

func TestNewHTTPClient(t *testing.T) {
	cli := repository.NewHTTPClient()

	assert.Implements(t, (*repository.HTTPClient)(nil), cli)
	assert.IsType(t, &http.Client{}, cli)
}
