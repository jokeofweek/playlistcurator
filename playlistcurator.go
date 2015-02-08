package playlistcurator

import (
	"github.com/jokeofweek/playlistcurator/printer"
	"github.com/jokeofweek/playlistcurator/providers"
)

func PrintTrack() {
	t := hardcoded.HardcodedProvider{}
	printer.PrintTrack(t)
}
