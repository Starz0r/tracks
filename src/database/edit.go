package database

import (
	"time"

	"github.com/spidernest-go/logger"
)

func (t *Track) Update() error {
	// TODO: Check if all struct fields are present
	tracks := db.Collection("music")
	rs := tracks.Find(t.ID)
	// TODO: Check if DateCreated is not equal
	t.DateModified = time.Now()
	err := rs.Update(t)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated from the table.")
	}
	return err
}

func (t *Track) Edit() error {
	tracks := db.Collection("music")
	rs := tracks.Find(t.ID)
	// TODO: Check if DateCreated is not equal
	t.DateModified = time.Now()
	err := rs.Update(t)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated from the table.")
	}
	return err
}
