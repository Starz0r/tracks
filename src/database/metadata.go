package database

import (
	"github.com/spidernest-go/logger"
)

func (t *Track) New() error {
	_, err := db.InsertInto("music").
		Values(t).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be inserted into the table.")
	}

	return err
}
