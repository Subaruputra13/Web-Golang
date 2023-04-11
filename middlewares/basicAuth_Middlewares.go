package middlewares

import (
	"crypto/subtle"

	"github.com/labstack/echo"
)

func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte("subaru")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("123")) == 1 {
		return true, nil
	}
	return false, nil
}
