package repository_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/repository"
)

func TestNewStdOut(t *testing.T) {
	repo := repository.NewStdOut()

	assert.Implements(t, (*tviso.WriteRepository)(nil), repo)
	assert.IsType(t, repository.StdOut{}, repo)
}

func TestStdOut_StoreCollection(t *testing.T) {
	repo := repository.NewStdOut()
	collection := []tviso.Media{
		{
			Name: "testMedia",
		},
	}

	output := captureOutput(func() {
		err := repo.StoreCollection(collection)
		assert.NoError(t, err)
	})

	assert.Equal(t, "testMedia\nTotal processed: 1\n", output)
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		var buf bytes.Buffer

		wg.Done()

		_, _ = io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()

	_ = writer.Close()

	return <-out
}
