package routing

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/orchestrafm/music/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func getTrackByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for getting a track.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	t, err := database.SelectID(id)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("")
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	return c.JSON(http.StatusOK, t)
}

func getTrackByName(c echo.Context) error {
	title := c.Param("title")

	ts, err := database.SelectName(title)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("")
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	return c.JSON(http.StatusOK, ts)
}
