package rest

import (
	"github.com/nvo-liat/platform-usergroup/src/service"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	auth "github.com/nvo-liat/platform-auth/entity"
)

func RestSessionRequest(ctx echo.Context, action string) (sd *auth.SessionData, e error) {
	if u := ctx.Get("user"); u != nil {
		c := u.(*jwt.Token).Claims.(jwt.MapClaims)
		sd, e = service.NewAuthService().Session(c["user_id"].(string), action)
	}

	if e != nil {
		if e.Error() == echo.ErrServiceUnavailable.Error() {
			e = echo.NewHTTPError(echo.ErrServiceUnavailable.Code)
		} else {
			e = echo.ErrUnauthorized
		}
	}

	return
}
