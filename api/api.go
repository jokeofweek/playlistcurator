package api

import "strings"

type Track struct {
	Artist string // Must be in lowercase.
	Name string
	Path string
}

func NewTrack(artist, name, path string) Track {
	return Track{strings.ToLower(artist), name, path}
}

type LibraryProvider interface {
	ProvideTracks() ([]Track, error)
}

type PlaylistCreator interface{
	CreatePlaylist([]Track) error
}