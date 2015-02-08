package m3u

import (
	"bufio"
	"os"
	"github.com/jokeofweek/playlistcurator/api"
)

type M3UPlaylistCreator struct {
	path string
}

func NewM3UPlaylistCreator(path string) M3UPlaylistCreator {
	return M3UPlaylistCreator{path}
}


func (c M3UPlaylistCreator) CreatePlaylist(tracks []api.Track) error {
	out, err := os.Create(c.path)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	for _, item := range tracks {
		writer.WriteString(item.Path)
		writer.WriteString("\n")
	}
	writer.Flush()

	return nil
}