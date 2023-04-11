package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"Web-Golang/models"

	"github.com/labstack/echo"
)

func GetKucing(c echo.Context) error {
	// Query Parameter
	kucing := c.QueryParam("kucing")
	status := c.QueryParam("status")

	datatype := c.Param("type")

	if datatype == "string" {
		return c.String(http.StatusOK, "Kucing:"+kucing+", status:"+status)
	}

	if datatype == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"kucing": kucing,
			"status": status,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Tipe harus string dan json !",
	})
}

func PostKucing(c echo.Context) error {
	kucing := models.Kucing{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&kucing)

	if err != nil {
		log.Printf("Gagal melakukan decode %s", err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Gagal melakukan decode",
		})
	}

	log.Printf("Berhasil mendapatkan Kucing: %v\n", kucing)
	return c.JSON(http.StatusOK, map[string]string{
		"status": "berhasil",
	})

}