package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"tviso-scrapper/pkg/tviso"
)

const (
	ListCollectionEndpoint = "/user/collection?mediaType=&status=&sortType=date&sortDirection=normal"
	FullInfoEndpoint = "/media/full_info?mediaType=1&liveAvailability=true"
)


type TvisoAPI struct {
	Config Config
}

func NewTvisoAPI() tviso.ReadRepository {
	return TvisoAPI{
		Config:NewConfig(),
	}
}

func (t TvisoAPI) GetUserCollection() ([]tviso.Media, error) {
	page := 0
	hasMore := true
	collection := []tviso.Media{}

	for hasMore {
		cr, err := getCollectionForUserPage(t.Config.APIAddr, page)
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

	contents, err := readURL(url)
	if err != nil {
		return tviso.Results{}, err
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	cr := tviso.Results{}
	err = json.Unmarshal(contents, &cr)
	if err != nil {
		return tviso.Results{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return cr, nil
}

func (t TvisoAPI) GetMediaInfo(m *tviso.Media) error {
	url := fmt.Sprintf("%v%v&idm=%v", t.Config.APIAddr, FullInfoEndpoint, m.ID)

	content, err := readURL(url)
	if err != nil {
		return err
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	err = json.Unmarshal(content, m)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	return nil
}

func readURL(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		if bErr := r.Body.Close(); bErr != nil {
			fmt.Printf("error closing body: %v", bErr)
		}
	}()

	if err := checkStatusCode(r); err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	return contents, nil
}

func checkStatusCode(r *http.Response) error {
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %v", r.StatusCode)
	}

	return nil
}
