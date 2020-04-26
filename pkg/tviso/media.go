//go:generate mockgen -destination=mocks/media_mock.go -package=mocks . ReadRepository,WriteRepository

package tviso

import "time"

type MediaType int32

const (
	SeriesMediaType  MediaType = 1
	MoviesMediaType  MediaType = 2
	TVShowMediaType  MediaType = 4
	EpisodeMediaType MediaType = 5
)

type ReadRepository interface {
	GetUserCollection() ([]Media, error)
	GetMediaInfo(*Media) error
}

type WriteRepository interface {
	StoreCollection([]Media) error
}

type Results struct {
	Results CollectionResult `json:"results"`
}

type CollectionResult struct {
	HasMore    bool      `json:"hasMore"`
	Count      CountInfo `json:"count"`
	Collection []Media   `json:"collection"`
}

type CountInfo struct {
	All       int            `json:"all"`
	MediaType MediaTypeCount `json:"mediaType"`
	Status    StatusCount    `json:"status"`
}

type MediaTypeCount struct {
	Series  int `json:"series"`
	Movies  int `json:"movies"`
	TVShows int `json:"tv-shows"`
}

type StatusCount struct {
	Pending   int `json:"pending"`
	Following int `json:"following"`
	Watched   int `json:"watched"`
}

type Media struct {
	ID           int          `json:"idm"`
	Name         string       `json:"name"`
	MediaType    MediaType    `json:"mediaType"`
	MediaStyle   string       `json:"mediaStyle"`
	IMDB         string       `json:"imdb"`
	Tags         []Tag        `json:"tags"`
	Rating       float64      `json:"rating"`
	Images       Images       `json:"images"`
	Artwork      Artworks     `json:"artwork"`
	Year         int          `json:"year"`
	Runtime      int          `json:"runtime"`
	AgeRating    int          `json:"ageRating"`
	MainGenre    string       `json:"mainGenre"`
	OriginalName string       `json:"originalName"`
	ShortPlot    string       `json:"shortPlot"`
	Availability Availability `json:"availability"`
	UserData     UserData     `json:"userData"`

	// Only for Series
	NumEpisodes int `json:"numEpisodes,omitempty"`
	NumSeasons  int `json:"numSeasons,omitempty"`

	// Only Full info
	Plot        string      `json:"plot,omitempty"`
	Cast        []Cast      `json:"cast,omitempty"`
	Genres      []string    `json:"genres,omitempty"`
	Distributor string      `json:"distributor,omitempty"`
	Producers   []Producers `json:"producers,omitempty"`
	Writers     []Writers   `json:"writers,omitempty"`
	Directors   []Directors `json:"directors,omitempty"`
	Composer    Person      `json:"composer"`
	Countries   []string    `json:"countries,omitempty"`
	Status      int         `json:"status"`
	StatusMedia string      `json:"statusMedia"`

	// Only for Series Full info
	SeasonsBlocked []int    `json:"seasonsBlocked,omitempty"`
	Seasons        []Season `json:"seasons"`
}

type Season struct {
	Limit        int       `json:"limit,omitempty"`
	TotalResults int       `json:"totalResults,omitempty"`
	HasMore      bool      `json:"hasMore,omitempty"`
	Page         int       `json:"page,omitempty"`
	SeasonName   string    `json:"seasonName,omitempty"`
	SeasonNum    int       `json:"seasonNum"`
	Episodes     []Episode `json:"episodes"`
}

type Episode struct {
	ID           interface{}   `json:"idm"`
	MediaType    MediaType     `json:"mediaType"`
	Name         string        `json:"name"`
	Plot         string        `json:"plot,omitempty"`
	Images       Images        `json:"images"`
	Runtime      int           `json:"runtime"`
	Artwork      Artworks      `json:"artwork"`
	SeasonNum    int           `json:"seasonNum"`
	EpisodeNum   int           `json:"episodeNum"`
	ReleaseDates []ReleaseDate `json:"releaseDates"`
	Availability Availability  `json:"availability"`
	SeasonName   string        `json:"seasonName,omitempty"`
}

type ReleaseDate struct {
	Country string    `json:"country"`
	Date    time.Time `json:"date"`
}

type Person struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Images map[string]string `json:"images"`
}

type Cast struct {
	Person
	Role string `json:"role"`
}

type Producers struct {
	Person
}

type Writers struct {
	Person
}

type Directors struct {
	Person
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Images struct {
	Poster   string `json:"poster"`
	Backdrop string `json:"backdrop"`
	Country  string `json:"country"`
}

type Artworks struct {
	Backdrops Backdrops `json:"backdrops"`
	Posters   Posters   `json:"posters"`
}
type Artwork struct {
	Original string `json:"original"`
	Large    string `json:"large"`
	Small    string `json:"small"`
}

type Backdrops struct {
	Artwork
}
type Posters struct {
	Artwork
}

type Availability struct {
	VODSummary      []string      `json:"vodSummary"`
	ScheduleSummary string        `json:"scheduleSummary,omitempty"`
	Schedule        []string      `json:"schedule"`
	InTheaters      bool          `json:"inTheaters"`
	VOD             []VODProvider `json:"vod,omitempty"`
	VODBestOffer    VODProvider   `json:"vodBestOffer"`
	VODUserOffers   []VODProvider `json:"vodUserOffers"`
}

type VODProvider struct {
	ProviderID                  string   `json:"providerId"`
	Languages                   []string `json:"languages,omitempty"`
	Quality                     []string `json:"quality,omitempty"`
	URL                         VODUrl   `json:"url"`
	SubscriptionRequiredForUser bool     `json:"subscriptionRequiredForUser"`
	RegisterRequiredForUser     bool     `json:"registerRequiredForUser"`
	PpvRequiredForUser          bool     `json:"ppvRequiredForUser"`
	Currency                    string   `json:"currency"`
	MinPrice                    float64  `json:"minPrice,omitempty"`
}

type VODPricing struct {
	Free              bool               `json:"free"`
	FreeForRegistered bool               `json:"freeForRegistered"`
	Subscription      bool               `json:"subscription"`
	PayPerView        bool               `json:"payPerView"`
	Prices            map[string]float64 `json:"prices"`
	Currency          string             `json:"currency"`
}

type VODUrl struct {
	Embed  string `json:"embed"`
	Native string `json:"native"`
	Web    string `json:"web"`
}

type UserData struct {
	Status  string    `json:"status"`
	AddedAt time.Time `json:"addedAt"`
}
