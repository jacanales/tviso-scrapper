package tviso

import "fmt"

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return fmt.Errorf("could not get user collection: %w", err)
	}

	if len(collection) == 0 {
		return nil
	}

	var errors []error

	for i := range collection {
		media := collection[i]

		err := read.GetMediaInfo(&media)
		if err != nil {
			errors = append(errors, fmt.Errorf("media: %v, error: %w", media.ID, err))
		}

		collection[i] = media
	}

	if len(errors) > 0 {
		return NewErrGettingMediaInfo(errors)
	}

	return write.StoreCollection(collection)
}
