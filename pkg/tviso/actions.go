package tviso

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return err
	}

	if len(collection) == 0 {
		return nil
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
