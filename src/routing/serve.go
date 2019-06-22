package routing

import (
	"github.com/spidernest-go/mux"
)

var r *echo.Echo

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`

func ListenAndServe() {
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.POST("/music", createTrack)
	v0.PUT("/music", updateTrack)
	v0.PATCH("/music/:id", editTrack)
	v0.GET("/music/:id", getTrackByID)
	v0.DELETE("/music/:id", deleteTrack)
	v0.GET("/music/:title", getTracksByName)

	r.Start(":5000")
}
