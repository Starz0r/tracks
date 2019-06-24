package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/tracks/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func deleteTrack(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for deleting a track.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	err = database.Remove(id)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for getting a track.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}
	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Track deleted successfully."})
}
