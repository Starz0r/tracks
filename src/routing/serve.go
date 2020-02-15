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
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyNbZz3Ig6VWUTxsBt5d4
Co+9VKIHm4BvQjG4ynh2v3a5an+gE7V6wY5ExBvIPNqOJnJWnvvEk22wYPB3to1T
6KMlpTmWmuO9aqBaLBwDY42UctS30B18bcOpz8wZy5gL1BkheTExfg09yOj0igW1
gMNyVCVYuhh5ye8NAinMCNxc9QgLz6ODxGXIfVlNN96C0iGhxAto7x9cMYTaT2FS
9GN6ZuOlbV4RnlmaiI+avbga6sy4m0WEiRFcx5Je7GZhsmtuQ65PaeUiOM/MpWNB
doBgwAWghhHc4WSTqyGbsVgl82qHvV+7Z9MmGq1k9fUk5zNtnP7Ou+gv2FBEMu9p
QQIDAQAB
-----END PUBLIC KEY-----`

func ListenAndServe() {
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

	r.Start(":5001")
}
