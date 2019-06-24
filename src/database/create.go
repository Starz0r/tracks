package database

import (
	"time"

	"github.com/spidernest-go/logger"
)

func (t *Track) New() error {
	t.DateCreated = time.Now()
	t.DateModified = time.Unix(0, 0)
	_, err := db.InsertInto("tracks").
		Values(t).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be inserted into the table.")
	}

	return err
}
