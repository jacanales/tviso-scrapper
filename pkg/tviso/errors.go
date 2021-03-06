package tviso

import (
	"bytes"
	"errors"
)

var ErrRequestError = errors.New("request error")

type ErrGettingMediaInfo struct {
	error
	errors []error
}

func NewErrGettingMediaInfo(err []error) error {
	return ErrGettingMediaInfo{errors: err}
}

func (e ErrGettingMediaInfo) Error() string {
	var b bytes.Buffer

	for _, e := range e.errors {
		b.WriteString(e.Error() + "\n")
	}

	return b.String()
}
