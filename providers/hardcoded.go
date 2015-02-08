package hardcoded

import (
	"github.com/jokeofweek/playlistcurator/api"
)

type HardcodedProvider struct {}

func (p HardcodedProvider) ProvideTracks() []api.Track {
	return []api.Track{
		api.Track{"Hooded Fang", "Wasteland", ""},
		api.Track{"Daft Punk", "One More Time", ""},
	}
}