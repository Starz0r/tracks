package routing

import (
	"net/http"

	"github.com/orchestrafm/music/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createTrack(c echo.Context) error {
	t := new(database.Track)
	if err := c.Bind(t); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed music track data.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music track data was invalid or malformed."})
	}

	err := t.New()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music track data did not get submitted to the database."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "OK."})
}
