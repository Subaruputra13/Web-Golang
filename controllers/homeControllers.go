package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeController(c echo.Context) error {
	return c.String(http.StatusOK, "Nama Saya Alfian Subaru Putra")
}
