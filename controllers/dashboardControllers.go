package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetDashboard(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "berhasil",
	})
}
