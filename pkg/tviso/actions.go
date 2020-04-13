package tviso

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return err
	}

	for _, m := range collection {
		_ = read.GetMediaInfo(&m)
	}

	return write.StoreCollection(collection)
}
