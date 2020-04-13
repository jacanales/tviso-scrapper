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
	FullInfoEndpoint = "/media/full_info?liveAvailability=true"
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
		cr, err := getCollectionForUserPage(t.Config.APIAddr, t.Config.Cookie, page)
		if err != nil {
			return nil, err
		}

		hasMore = cr.Results.HasMore
		page++

		collection = append(collection, cr.Results.Collection...)
	}

	return collection, nil
}

func getCollectionForUserPage(serverURL string, cookie string, page int) (tviso.Results, error) {
	url := fmt.Sprintf("%v%v&page=%v", serverURL, ListCollectionEndpoint, page)

	contents, err := readURL(url, cookie)
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
	url := fmt.Sprintf("%v%v&idm=%v&mediaType=%v", t.Config.APIAddr, FullInfoEndpoint, m.ID, m.MediaType)

	content, err := readURL(url, t.Config.Cookie)
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

func readURL(url string, cookie string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Cookie", cookie)

	client := http.DefaultClient

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if bErr := r.Body.Close(); bErr != nil {
			fmt.Printf("error closing body: %v", bErr)
		}
	}()

	if err := checkStatusCode(r); err != nil {
		return readURL(url, cookie)
	}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	return contents, nil
}

func checkStatusCode(r *http.Response) error {
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %v, message: %v", r.StatusCode, r.Body)
	}

	return nil
}
