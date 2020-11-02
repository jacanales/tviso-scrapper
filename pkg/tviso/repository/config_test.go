package repository_test

import (
	"testing"
	"tviso-scrapper/pkg/tviso/repository"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := repository.NewConfig()

	assert.IsType(t, repository.Config{}, cfg)
}
