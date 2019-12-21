package routing

import (
	"net/http"

	"github.com/orchestrafm/tracks/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func updateTrack(c echo.Context) error {
	// auth check
	admin, auth := AuthorizationCheck(c)
	if auth != true {
		logger.Info().
			Msg("user intent to create a update a track, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Insufficient Permissions."})
	}

	// data binding
	t := new(database.Track)
	if err := c.Bind(t); err != nil {
		logger.Error().
			Err(err).
			Msg("Request Data could not be binded to datastructure.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	// update resource
	if !admin && t.Publisher != SelfAuthCheck(c).Subject {
		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Missing Ownership."})
	}
	err := t.Update()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Track edited successfully."})
}
