package tviso

import `time`

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return err
	}

	for _, m := range collection {
		_ = read.GetMediaInfo(&m)

		time.Sleep(200*time.Millisecond)
	}

	return write.StoreCollection(collection)
}
