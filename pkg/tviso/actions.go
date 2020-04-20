package tviso

import `fmt`

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return fmt.Errorf("could not get user collection: %w", err)
	}

	if len(collection) == 0 {
		return nil
	}

	var errors []error
	for _, m := range collection {
		err := read.GetMediaInfo(&m)
		if err != nil {
			errors = append(errors, fmt.Errorf("media: %v, error: %w", m.ID, err))
		}
	}

	if len(errors) > 0 {
		return NewErrGettingMediaInfo(errors)
	}

	return write.StoreCollection(collection)
}
