package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"tviso-scrapper/pkg/tviso"
)

const ListCollectionEndpoint = "/user/collection?mediaType=&status=&sortType=date&sortDirection=normal"

type TvisoAPI struct {
}

func NewTvisoAPI() tviso.ReadRepository {
	return TvisoAPI{}
}

func (t TvisoAPI) GetUserCollection() ([]tviso.Media, error) {
	cfg := NewConfig()
	page := 0
	hasMore := true
	collection := []tviso.Media{}

	for hasMore {
		cr, err := getCollectionForUserPage(cfg.APIAddr, page)
		if err != nil {
			return nil, err
		}

		hasMore = cr.Results.HasMore
		page++

		collection = append(collection, cr.Results.Collection...)
	}

	return collection, nil
}

func getCollectionForUserPage(serverURL string, page int) (tviso.Results, error) {
	url := fmt.Sprintf("%v%v&page=%v", serverURL, ListCollectionEndpoint, page)

	cr := tviso.Results{}

	r, err := http.Get(url)
	if err != nil {
		return cr, err
	}

	defer func() {
		if bErr := r.Body.Close(); bErr != nil {
			fmt.Printf("error closing body: %v", bErr)
		}
	}()

	if r.StatusCode != http.StatusOK {
		return cr, fmt.Errorf("error: %v", r.StatusCode)
	}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return cr, fmt.Errorf("cannot read body: %w", err)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	err = json.Unmarshal(contents, &cr)
	if err != nil {
		return tviso.Results{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return cr, nil
}
