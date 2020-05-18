package tviso_test

import (
	`testing`

	`github.com/magiconair/properties/assert`

	`tviso-scrapper/pkg/tviso`
)

func TestMediaType_String(t *testing.T) {
	assert.Equal(t, tviso.SeriesMediaType.String(), "series")
	assert.Equal(t, tviso.MoviesMediaType.String(), "movie")
	assert.Equal(t, tviso.TVShowMediaType.String(), "tv-show")
	assert.Equal(t, tviso.EpisodeMediaType.String(), "episode")
}

func TestMediaStatus_String(t *testing.T) {
	assert.Equal(t, tviso.Watched.String(), "watched")
	assert.Equal(t, tviso.Pending.String(), "pending")
	assert.Equal(t, tviso.Following.String(), "following")
}