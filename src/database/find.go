package database

import (
	"database/sql"

	"github.com/spidernest-go/logger"
)

func SelectID(id uint64) (*Track, error) {
	tracks := db.Collection("tracks")
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

func SelectName(name string) ([]*Track, error) {
	var ts []*Track
	tracks := db.Collection("tracks")
	rs := tracks.Find().Where("title LIKE", name+"%")

	err := rs.All(&ts)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return ts, nil
	}
}

func SelectRecent(limit int) ([]*Track, error) {
	var ts []*Track
	tracks := db.Collection("tracks")
	rs := tracks.Find().OrderBy("id DESC").Limit(limit)

	err := rs.All(&ts)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return ts, nil
	}
}
