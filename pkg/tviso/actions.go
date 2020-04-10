package tviso

func GetUserCollection(read ReadRepository, write WriteRepository) error {
	collection, err := read.GetUserCollection()
	if err != nil {
		return err
	}

	return write.StoreCollection(collection)
}
