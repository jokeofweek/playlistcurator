package banshee

import (
	"database/sql"
	"github.com/jokeofweek/playlistcurator/api"
	_ "github.com/mattn/go-sqlite3"
)

type BansheeProvider struct {
	dbPath string
}

func (p BansheeProvider) ProvideTracks() ([]api.Track, error) {
	// Open the database
	db, err := sql.Open("sqlite3", p.dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStmt := `
	SELECT a.Name, Title, URI FROM CoreTracks t 
	JOIN CoreArtists a ON a.ArtistId = t.ArtistId 
	WHERE uri LIKE 'file://%'`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate though the query set, building up the tracks
	var tracks []api.Track
	for rows.Next() {
		var artist sql.NullString
		var name sql.NullString
		var path sql.NullString
		
		err = rows.Scan(&artist, &name, &path)
		if err != nil {
			return nil, err 
		}

		if artist.Valid && name.Valid && path.Valid {
			tracks = append(tracks, api.NewTrack(artist.String, name.String, path.String))
		}
	}

	return tracks, rows.Err()
}

func NewBansheeProvider(path string) BansheeProvider {
	return BansheeProvider{path}
}
