package playlistcurator

import (
	"gopkg.in/fatih/set.v0"
	"github.com/jokeofweek/playlistcurator/api"
	"github.com/jokeofweek/playlistcurator/printer"
	"github.com/jokeofweek/playlistcurator/providers/hardcoded"
	"io/ioutil"
	"net/http"
	"strings"
)

func PrintTrack() {
	t := hardcoded.HardcodedProvider{}
	printer.PrintTrack(t)
}

func GetArtists(tracks []api.Track) *set.Set {
	result := set.New()
	for _, track := range tracks {
		result.Add(track.Artist)
	}
	return result
}

func getSimilarTo(artist string) ([]string, error) {
	resp, err := http.Get("http://www.gnoosic.com/api/top?name=" + strings.Replace(artist, " ", "+", -1))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	artists := strings.Split(string(body[:]), "\n")

	// Convert all to lower case
	for index, item := range artists {
		artists[index] = strings.ToLower(item)
	}

	// The API returns 10 elements, but in reality we have an extra in the slice due to 
	// a trailing new line, so we have to remove this.
	return artists[0:10], nil
}

func GetSimilarArtists(seedArtist string, availableArtists *set.Set, depth int) (*set.Set, error) {
	result := set.New()

	currentSet := set.New()
	currentSet.Add(seedArtist)
	var awaitingSet (*set.Set)

	for depth > 0 {
		awaitingSet = set.New()

		// Go through everyone in the current set, finding similar artists if they haven't
		// already been used in the result
		current := set.StringSlice(currentSet)
		for _, item := range current {
			if !result.Has(item) {
				result.Add(item)

				// Fetch the similar artists 
				similarArtists, err := getSimilarTo(item)
				if err != nil {
					return nil, err
				}

				// Go through the similar artists, adding the ones which are in the set of available.
				for _, similarArtist := range similarArtists {
					if availableArtists.Has(similarArtist) {
						awaitingSet.Add(similarArtist)
					}
				}
			}
		}

		currentSet = awaitingSet
		depth--
	}

	return result, nil
}