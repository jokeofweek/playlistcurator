package api

type Track struct {
	Artist string
	Name string
	Path string
}

type LibraryProvider interface {
	ProvideTracks() ([]Track, error)
}