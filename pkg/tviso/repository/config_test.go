package repository_test

import (
	"os"
	"testing"
	"tviso-scrapper/pkg/tviso/repository"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	envs := setUp()
	retCode := m.Run()
	tearDown(envs)

	os.Exit(retCode)
}

func setUp() map[string]string {
	original := map[string]string{}
	mockedVars := map[string]string{
		"TVISO_ENDPOINT": "localhost:30011",
		"TVISO_COOKIE":   "test_cookie",
	}

	for varName, mockedValue := range mockedVars {
		original[varName] = os.Getenv(varName)
		_ = os.Setenv(varName, mockedValue)
	}

	return original
}

func tearDown(original map[string]string) {
	for name, value := range original {
		_ = os.Setenv(name, value)
	}
}

func TestNewConfig(t *testing.T) {
	cfg := repository.NewConfig()
	a := assert.New(t)

	a.IsType(repository.Config{}, cfg)
	a.Equal("localhost:30011", cfg.APIAddr)
	a.Equal("test_cookie", cfg.Cookie)
}
