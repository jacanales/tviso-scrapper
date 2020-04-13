package tviso

import `bytes`

type ErrGettingMediaInfo struct {
	error
	errors []error
}

func NewErrGettingMediaInfo(err []error) error {
	return ErrGettingMediaInfo{errors:err}
}

func (e ErrGettingMediaInfo) Error() string {
	var b bytes.Buffer

	for _, e := range e.errors {
		b.WriteString(e.Error() + "\n")
	}

	return b.String()
}

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return err
	}

	var errors []error
	for _, m := range collection {
		err := read.GetMediaInfo(&m)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return NewErrGettingMediaInfo(errors)
	}

	return write.StoreCollection(collection)
}
