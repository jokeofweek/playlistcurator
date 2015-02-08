package playlistcurator

import (
	"github.com/jokeofweek/playlistcurator/printer"
	"github.com/jokeofweek/playlistcurator/providers/hardcoded"
)

func PrintTrack() {
	t := hardcoded.HardcodedProvider{}
	printer.PrintTrack(t)
}
