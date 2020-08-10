package routing

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"

	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
	"github.com/spidernest-go/mux/middleware"
)

var r *echo.Echo

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`
const rsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl3Mw0lnzWr+KrhhP1/jnKHblCM/DqIhvUHgsOYZWrE3+fEvHjc6wrUT9RtC3eRZfRxtdyxa9CPuSnPEt/Jmu2YPVRWxOVUJfUxgZQg0OPXurMy0h6O1Yal4s9yNq0+OmCSIE3DFVNTs5hlYNI7TNkjPp/UJx8Xc+J+g/gUPrIVQo+XWNGoKv+udiQhi9LrYZuQOy9MZPKgUKSfJwmwWRBb7CZmvWSwprQ3/619+2vf1gS/K3vqenlZfCRFadPuxebmQ595LKAn0tgnw2R0c4aAU/G1LsJsBFfY0kvhE/asFvNSoAoJA3jnQMYmMekqgVdVNV2FLrLWve5520EjTeMQIDAQAB
-----END PUBLIC KEY-----`

func ListenAndServe() error {
	// decode pem block into rsa public key
	block, _ := pem.Decode([]byte(rsaPublicKey))
	if block == nil {
		logger.Fatal().
			Msg("PEM RSA public key block was invalid and failed to decode.")
	}
	pkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Decoded PEM block failed to parse.")
	}
	rsaKey, ok := pkey.(*rsa.PublicKey)
	if !ok {
		logger.Fatal().
			Msgf("got unexpected key type: %T", pkey)
	}

	// route apis and start http multiplexer
	r = echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	v0 := r.Group("/api/v0")
	v0AuthReq := v0.Group("", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "RS256",
		SigningKey:    rsaKey,
	}))
	v0AuthReq.POST("/track", createTrack)
	v0AuthReq.PUT("/track", updateTrack)
	v0AuthReq.PATCH("/track/:id", editTrack)
	v0.GET("/track/:id", getTrackByID)
	v0AuthReq.DELETE("/track/:id", deleteTrack)
	v0.GET("/track/:title", getTracksByName)
	v0.GET("/track/recent/:limit", getTracksByRecent)

	return r.Start(":5001")
}
