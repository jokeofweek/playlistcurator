package main

import (
	"fmt"
    "log"
    "os/user"
	"path"
	"gopkg.in/fatih/set.v0"
	"github.com/jokeofweek/playlistcurator"
	"github.com/jokeofweek/playlistcurator/providers/banshee"
)

func main() {

    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }

	provider := banshee.NewBansheeProvider(path.Join(usr.HomeDir, ".config/banshee-1/banshee.db"))
	tracks, err := provider.ProvideTracks()

	if err != nil {
		log.Fatal (err)
	}

	artists := playlistcurator.GetArtists(tracks)
	similarArtists, err := playlistcurator.GetSimilarArtists("m83", artists, 3)

	if err != nil {
		log.Fatal (err)
	}
	for _, item := range set.StringSlice(similarArtists) {
		fmt.Println(item)
	}

	
}
