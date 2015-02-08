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

	result, err := playlistcurator.CreatePlaylist(provider, m3u.M3UPlaylistCreator{}, "flying lotus", 3)

	if err != nil {
		log.Fatal (err)
	} else {
		fmt.Println(result)
	}
}
