//go:generate mockgen -destination=mocks/media_mock.go -package=mocks . ReadRepository,WriteRepository

package tviso

import "time"

type (
	MediaType   int32
	MediaStatus int
)

const (
	SeriesMediaType  MediaType = 1
	MoviesMediaType  MediaType = 2
	TVShowMediaType  MediaType = 4
	EpisodeMediaType MediaType = 5
)

const (
	Watched MediaStatus = iota
	Pending
	Following
)

func (s MediaType) String() string {
	return [...]string{"unknown", "series", "movie", "unknown", "tv-show", "episode"}[s]
}

func (s MediaType) Val() int32 {
	return int32(s)
}

func (s MediaStatus) String() string {
	return [...]string{"watched", "pending", "following"}[s]
}

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
	ID           int          `json:"idm" bson:"tviso_id"`
	Name         string       `json:"name" bson:"name"`
	MediaType    MediaType    `json:"mediaType" bson:"media_type"`
	MediaStyle   string       `json:"mediaStyle" bson:"media_style"`
	IMDB         string       `json:"imdb" bson:"imdb"`
	Tags         []Tag        `json:"tags" bson:"tags"`
	Rating       float64      `json:"rating" bson:"rating"`
	Images       Images       `json:"images" bson:"images"`
	Artwork      Artworks     `json:"artwork" bson:"artwork"`
	Year         int          `json:"year" bson:"year"`
	Runtime      int          `json:"runtime" bson:"runtime"`
	AgeRating    int          `json:"ageRating" bson:"age_rating"`
	MainGenre    string       `json:"mainGenre" bson:"main_genre"`
	OriginalName string       `json:"originalName" bson:"original_name"`
	ShortPlot    string       `json:"shortPlot" bson:"short_plot"`
	Availability Availability `json:"availability" bson:"availability"`
	UserData     UserData     `json:"userData" bson:"user_data"`

	// Only for Series
	NumEpisodes int `json:"numEpisodes,omitempty" bson:"num_episodes,omitempty"`
	NumSeasons  int `json:"numSeasons,omitempty" bson:"num_seasons,omitempty"`

	// Only Full info
	Plot        string      `json:"plot,omitempty" bson:"plot,omitempty"`
	Cast        []Cast      `json:"cast,omitempty" bson:"cast,omitempty"`
	Genres      []string    `json:"genres,omitempty" bson:"genres,omitempty"`
	Distributor string      `json:"distributor,omitempty" bson:"distributor,omitempty"`
	Producers   []Producers `json:"producers,omitempty" bson:"producers,omitempty"`
	Writers     []Writers   `json:"writers,omitempty" bson:"writers,omitempty"`
	Directors   []Directors `json:"directors,omitempty" bson:"directors,omitempty"`
	Composer    Person      `json:"composer" bson:"composer"`
	Countries   []string    `json:"countries,omitempty" bson:"countries,omitempty"`
	Status      int         `json:"status" bson:"status"`
	StatusMedia string      `json:"statusMedia" bson:"status_media"`

	// Only for Series Full info
	SeasonsBlocked []int    `json:"seasonsBlocked,omitempty" bson:"seasons_blocked,omitempty"`
	Seasons        []Season `json:"seasons" bson:"seasons"`
}

type Season struct {
	Limit        int       `json:"limit,omitempty" bson:"limit,omitempty"`
	TotalResults int       `json:"totalResults,omitempty" bson:"total_results,omitempty"`
	HasMore      bool      `json:"hasMore,omitempty" bson:"has_more,omitempty"`
	Page         int       `json:"page,omitempty" bson:"page,omitempty"`
	SeasonName   string    `json:"seasonName,omitempty" bson:"season_name,omitempty"`
	SeasonNum    int       `json:"seasonNum" bson:"season_num"`
	Episodes     []Episode `json:"episodes" bson:"episodes"`
}

type Episode struct {
	ID           interface{}   `json:"idm" bson:"idm"`
	MediaType    MediaType     `json:"mediaType" bson:"media_type"`
	Name         string        `json:"name" bson:"name"`
	Plot         string        `json:"plot,omitempty" bson:"plot,omitempty"`
	Images       Images        `json:"images" bson:"images"`
	Runtime      int           `json:"runtime" bson:"runtime"`
	Artwork      Artworks      `json:"artwork" bson:"artwork"`
	SeasonNum    int           `json:"seasonNum" bson:"season_num"`
	EpisodeNum   int           `json:"episodeNum" bson:"episode_num"`
	ReleaseDates []ReleaseDate `json:"releaseDates" bson:"release_dates"`
	Availability Availability  `json:"availability" bson:"availability"`
	SeasonName   string        `json:"seasonName,omitempty" bson:"season_name,omitempty"`
}

type ReleaseDate struct {
	Country string    `json:"country" bson:"country"`
	Date    time.Time `json:"date" bson:"date"`
}

type Person struct {
	ID     int               `json:"id" bson:"id"`
	Name   string            `json:"name" bson:"name"`
	Images map[string]string `json:"images" bson:"images"`
}

type Cast struct {
	Person
	Role string `json:"role" bson:"role"`
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
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

type Images struct {
	Poster   string `json:"poster" bson:"poster"`
	Backdrop string `json:"backdrop" bson:"backdrop"`
	Country  string `json:"country" bson:"country"`
}

type Artworks struct {
	Backdrops Backdrops `json:"backdrops" bson:"backdrops"`
	Posters   Posters   `json:"posters" bson:"posters"`
}

type Artwork struct {
	Original string `json:"original" bson:"original"`
	Large    string `json:"large" bson:"large"`
	Small    string `json:"small" bson:"small"`
}

type Backdrops struct {
	Artwork
}

type Posters struct {
	Artwork
}

type Availability struct {
	VODSummary      []string      `json:"vodSummary" bson:"vod_summary"`
	ScheduleSummary string        `json:"scheduleSummary,omitempty" bson:"schedule_summary,omitempty"`
	Schedule        []string      `json:"schedule" bson:"schedule"`
	InTheaters      bool          `json:"inTheaters" bson:"in_theaters"`
	VOD             []VODProvider `json:"vod,omitempty" bson:"vod,omitempty"`
	VODBestOffer    VODProvider   `json:"vodBestOffer" bson:"vod_best_offer"`
	VODUserOffers   []VODProvider `json:"vodUserOffers" bson:"vod_user_offers"`
}

type VODProvider struct {
	ProviderID                  string   `json:"providerId" bson:"provider_id"`
	Languages                   []string `json:"languages,omitempty" bson:"languages,omitempty"`
	Quality                     []string `json:"quality,omitempty" bson:"quality,omitempty"`
	URL                         VODUrl   `json:"url" bson:"url"`
	SubscriptionRequiredForUser bool     `json:"subscriptionRequiredForUser" bson:"subscription_required_for_user"`
	RegisterRequiredForUser     bool     `json:"registerRequiredForUser" bson:"register_required_for_user"`
	PpvRequiredForUser          bool     `json:"ppvRequiredForUser" bson:"ppv_required_for_user"`
	Currency                    string   `json:"currency" bson:"currency"`
	MinPrice                    float64  `json:"minPrice,omitempty" bson:"min_price,omitempty"`
}

type VODPricing struct {
	Free              bool               `json:"free" bson:"free"`
	FreeForRegistered bool               `json:"freeForRegistered" bson:"free_for_registered"`
	Subscription      bool               `json:"subscription" bson:"subscription"`
	PayPerView        bool               `json:"payPerView" bson:"pay_per_view"`
	Prices            map[string]float64 `json:"prices" bson:"prices"`
	Currency          string             `json:"currency" bson:"currency"`
}

type VODUrl struct {
	Embed  string `json:"embed" bson:"embed"`
	Native string `json:"native" bson:"native"`
	Web    string `json:"web" bson:"web"`
}

type UserData struct {
	Status  string    `json:"status" bson:"status"`
	AddedAt time.Time `json:"addedAt" bson:"added_at"`
}
