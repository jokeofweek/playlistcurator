package api

import "strings"

type Track struct {
	Artist string
	Name string
	Path string
}

func NewTrack(artist, name, path string) Track {
	return Track{strings.ToLower(artist), name, path}
}

type LibraryProvider interface {
	ProvideTracks() ([]Track, error)
}