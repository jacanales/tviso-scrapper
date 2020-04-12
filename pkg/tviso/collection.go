package tviso

import "time"

type MediaType int32

const (
	SeriesMediaType MediaType = 1
	MoviesMediaType MediaType = 2
)

type ReadRepository interface {
	GetUserCollection() ([]Media, error)
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
	MediaType    MediaType    `json:"mediaType"`
	MediaStyle   string       `json:"mediaStyle"`
	IMDB         string       `json:"imdb"`
	Tags         []Tag        `json:"tags"`
	Rating       float64      `json:"rating"`
	Name         string       `json:"name"`
	Images       Images       `json:"images"`
	Artwork      Artworks     `json:"artwork"`
	Year         int          `json:"year"`
	Runtime      int          `json:"runtime"`
	AgeRating    int          `json:"ageRating"`
	MainGenre    string       `json:"mainGenre"`
	OriginalName string       `json:"originalName"`
	ShortPlot    string       `json:"shortPlot"`
	NumEpisodes  int          `json:"numEpisodes,omitempty"`
	NumSeasons   int          `json:"numSeasons,omitempty"`
	Availability Availability `json:"availability"`
	UserData     UserData     `json:"userData"`
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
	InTheaters      bool          `json:"inTheaters"`
	VODBestOffer    VODProvider   `json:"vodBestOffer"`
	VODUserOffers   []VODProvider `json:"vodUserOffers"`
}

type VODProvider struct {
	ProviderID                  string  `json:"providerId"`
	URL                         VODUrl  `json:"url"`
	SubscriptionRequiredForUser bool    `json:"subscriptionRequiredForUser"`
	RegisterRequiredForUser     bool    `json:"registerRequiredForUser"`
	PpvRequiredForUser          bool    `json:"ppvRequiredForUser"`
	Currency                    string  `json:"currency"`
	MinPrice                    float64 `json:"minPrice,omitempty"`
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
