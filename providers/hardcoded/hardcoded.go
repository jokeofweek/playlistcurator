package hardcoded

import (
	"github.com/jokeofweek/playlistcurator/api"
)

type HardcodedProvider struct{}

func (p HardcodedProvider) ProvideTracks() ([]api.Track, error) {
	return []api.Track{
		api.NewTrack("Hooded Fang", "Wasteland", ""),
		api.NewTrack("Daft Punk", "One More Time", ""),
	}, nil
}
