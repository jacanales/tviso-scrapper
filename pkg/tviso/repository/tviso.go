//go:generate mockgen -destination=mocks/tviso_mock.go -package=mocks . HTTPClient

package repository

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"

	"tviso-scrapper/pkg/tviso"
)

const (
	ListCollectionEndpoint = "/user/collection?mediaType=&status=&sortType=date&sortDirection=normal"
	FullInfoEndpoint       = "/media/full_info?liveAvailability=true"
	ClientKeepAlive        = 5
	ConnectionTimeout      = 3
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHTTPClient() HTTPClient {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				KeepAlive: ClientKeepAlive * time.Second,
			}).DialContext,
			IdleConnTimeout: ConnectionTimeout * time.Second,
		},
	}

	return httpClient
}

type TvisoAPI struct {
	Config Config
	Client HTTPClient
}

func NewTvisoAPI(cli HTTPClient, cfg Config) tviso.ReadRepository {
	return TvisoAPI{
		Config: cfg,
		Client: cli,
	}
}

func (t TvisoAPI) GetUserCollection() ([]tviso.Media, error) {
	page := 0
	hasMore := true

	var collection []tviso.Media

	for hasMore {
		cr, err := t.getCollectionForUserPage(t.Config.APIAddr, t.Config.Cookie, page)
		if err != nil {
			return nil, err
		}

		hasMore = cr.Results.HasMore
		page++

		collection = append(collection, cr.Results.Collection...)
	}

	return collection, nil
}

func (t TvisoAPI) GetMediaInfo(m *tviso.Media) error {
	url := fmt.Sprintf("%v%v&idm=%v&mediaType=%v", t.Config.APIAddr, FullInfoEndpoint, m.ID, m.MediaType)

	content, err := t.readURL(url, t.Config.Cookie)
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

func (t TvisoAPI) getCollectionForUserPage(serverURL, cookie string, page int) (tviso.Results, error) {
	url := fmt.Sprintf("%v%v&page=%v", serverURL, ListCollectionEndpoint, page)

	contents, err := t.readURL(url, cookie)
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

func (t TvisoAPI) readURL(url, cookie string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Cookie", cookie)

	r, err := t.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	defer func() {
		_ = r.Body.Close()
	}()

	if err = checkStatusCode(r); err != nil {
		return t.readURL(url, cookie)
	}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	return contents, nil
}

func checkStatusCode(r *http.Response) error {
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %v, message: %v, error: %w", r.StatusCode, r.Body, tviso.ErrRequestError)
	}

	return nil
}
