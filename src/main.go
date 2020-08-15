package main

import (
	"github.com/orchestrafm/tracks/src/database"
	"github.com/orchestrafm/tracks/src/routers"
	"github.com/spidernest-go/logger"
)

func main() {
	err := database.Connect()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("MySQL Database could not be attached to.")
	}
	database.Synchronize()

	err = routers.ListenAndServe()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("REST Server failed to start.")
	}
}
