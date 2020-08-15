package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/orchestrafm/tracks/src/database"
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

func getTracksByName(c echo.Context) error {
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

func getTracksByRecent(c echo.Context) error {
	limit, err := strconv.ParseInt(c.Param("limit"), 10, 64)
	l := *new(int)
	if limit >= 50 {
		l = 50
	} else {
		l = int(limit)
	}

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for getting a track.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	ts, err := database.SelectRecent(l)
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
