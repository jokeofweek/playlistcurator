package main

import (
	"fmt"
    "log"
    "os"
    "os/user"
	"path"

	"github.com/jokeofweek/playlistcurator"
	"github.com/jokeofweek/playlistcurator/creators/m3u"
	"github.com/jokeofweek/playlistcurator/providers/banshee"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Expected 1 argument - artist name.")
		return
	}

    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }

	provider := banshee.NewBansheeProvider(path.Join(usr.HomeDir, ".config/banshee-1/banshee.db"))

	result, err := playlistcurator.CreatePlaylist(provider, m3u.M3UPlaylistCreator{}, args[0], 3)

	if err != nil {
		log.Fatal (err)
	} else {
		fmt.Println(result)
	}
}
