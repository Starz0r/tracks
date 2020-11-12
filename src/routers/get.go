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

	// offset is an optional parameter
	o, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil && c.QueryParam("offset") != "" {
		logger.Error().
			Err(err).
			Msgf("Passed offset parameter (%s) was not a valid number", c.QueryParam("offset"))

		return c.JSON(http.StatusBadRequest, nil)
	}
	// set offset to 0 if it wasn't present in the request
	if c.QueryParam("offset") == "" {
		o = 0
	}

	ts, err := database.SelectRecent(l, o)
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

func getCount(c echo.Context) error {
	amt, err := database.SelectCount()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Counting current amount of submitted tracks failed.")

		return c.JSON(http.StatusInternalServerError, &struct {
			Message string
		}{
			Message: "Database had an problem getting the current amount of tracks."})
	}

	return c.JSON(http.StatusOK, &struct {
		Count uint64 `json:"count"`
	}{
		Count: amt})

}
