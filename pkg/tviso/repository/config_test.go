package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tviso-scrapper/pkg/tviso/repository"
)

func TestNewConfig(t *testing.T) {
	cfg := repository.NewConfig()

	assert.IsType(t, repository.Config{}, cfg)
}
