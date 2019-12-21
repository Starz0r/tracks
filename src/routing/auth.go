package routing

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spidernest-go/mux"
)

type jwtExtendedClaims struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	Scope             string `json:"scope"`
	PreferredUsername string `json:"preferred_username"`
	jwt.StandardClaims
}

func FullAuthCheck(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtExtendedClaims)

	auth := strings.Contains(claims.Scope, "track:write") || strings.Contains(claims.Scope, "track:admin")
	admin := strings.Contains(claims.Scope, "track:admin")
	if !admin || !auth {
		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Insufficient Permissions."})
	}
	return nil
}

func AuthorizationCheck(c echo.Context) (bool, bool) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtExtendedClaims)

	auth := strings.Contains(claims.Scope, "track:write") || strings.Contains(claims.Scope, "track:admin")
	admin := strings.Contains(claims.Scope, "track:admin")

	return admin, auth
}

func SelfAuthCheck(c echo.Context) *jwtExtendedClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*jwtExtendedClaims)
}
