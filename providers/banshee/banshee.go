package banshee

import (
	"github.com/jokeofweek/playlistcurator/api"
	"github.com/mattn/go-sqlite3"
)

type BansheeProvider struct {
	dbPath string
}

func (p BansheeProvider) ProvideTracks() ([]api.Track, error) {
	return []api.Track{
		api.Track{"Hooded Fang", "Wasteland", ""},
		api.Track{"Daft Punk", "One More Time", ""},
	}
}

func NewBansheeProvider(path string) BansheeProvider {
	return BansheeProvider{path}
}