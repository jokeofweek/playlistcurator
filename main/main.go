package main

import (
    "os/user"
    "log"
	"path"
	"github.com/jokeofweek/playlistcurator/printer"
	"github.com/jokeofweek/playlistcurator/providers/banshee"
)

func main() {

    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }

	printer.PrintTrack(
		banshee.NewBansheeProvider(path.Join(usr.HomeDir, ".config/banshee-1/banshee.db")))
}
