package routing

import (
	"github.com/spidernest-go/mux"
)

var r *echo.Echo

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`

func ListenAndServe() {
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.POST("/track", createTrack)
	v0.PUT("/track", updateTrack)
	v0.PATCH("/track/:id", editTrack)
	v0.GET("/track/:id", getTrackByID)
	v0.DELETE("/track/:id", deleteTrack)
	v0.GET("/track/:title", getTracksByName)

	r.Start(":5000")
}
