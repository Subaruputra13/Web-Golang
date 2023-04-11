package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
)

func CheckCooki(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session")
		if err != nil {
			return err
		}

		token, err := c.Cookie("JwtCookie")
		if err != nil {
			return err
		}
		if cookie.Value == "admin" && token.Name == "JwtCookie" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Authentikasi Gagal")
	}
}
