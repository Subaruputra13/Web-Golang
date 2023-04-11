package middlewares

import "github.com/labstack/echo"

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Ini Server saya !")
		c.Response().Header().Set("Developer", "Alfian Subaru Putra")
		return next(c)
	}
}
