package printer

import (
	"fmt"
	"github.com/jokeofweek/playlistcurator/api"
)

func PrintTrack(t api.LibraryProvider) {
	tracks := t.ProvideTracks()
	for _, element := range tracks {
		fmt.Println(element)
	}
}
