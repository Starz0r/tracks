package database

import (
	"database/sql"

	"github.com/spidernest-go/logger"
)

func SelectID(id uint64) (*Track, error) {
	tracks := db.Collection("music")
	rs := tracks.Find(id)
	t := *new(Track)
	err := rs.One(&t)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return &t, nil
	}
}
