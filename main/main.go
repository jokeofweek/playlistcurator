package main

import (
	"fmt"
    "log"
    "os/user"
	"path"
	"github.com/jokeofweek/playlistcurator"
	"github.com/jokeofweek/playlistcurator/creators/m3u"
	"github.com/jokeofweek/playlistcurator/providers/banshee"
)

func main() {

    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }

	provider := banshee.NewBansheeProvider(path.Join(usr.HomeDir, ".config/banshee-1/banshee.db"))
	creator := m3u.NewM3UPlaylistCreator(path.Join(usr.HomeDir, "test.m3u"))

	err = playlistcurator.CreatePlaylist(provider, creator, "flying lotus", 3)

	if err != nil {
		log.Fatal (err)
	}

	fmt.Println("Created!")	
}
