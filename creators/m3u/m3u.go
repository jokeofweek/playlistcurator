package m3u

import (
	"github.com/jokeofweek/playlistcurator/api"
)

type M3UPlaylistCreator struct {}


func (c M3UPlaylistCreator) CreatePlaylist(tracks []api.Track) string {
	result := ""
	for _, item := range tracks {
		result += item.Path + "\n"
	}

	return result
}