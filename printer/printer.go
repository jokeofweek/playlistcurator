package printer

import (
	"fmt"
	"log"
	"github.com/jokeofweek/playlistcurator/api"
)

func PrintTrack(t api.LibraryProvider) {
	tracks, err := t.ProvideTracks()
	if err != nil {
		log.Fatal(err)
	}
	for _, element := range tracks {
		fmt.Println(element)
	}
}
