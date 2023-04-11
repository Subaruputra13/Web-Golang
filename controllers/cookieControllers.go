package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetCookie(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Cookie Berhasil",
		"page":    "Cookie",
	})
}
