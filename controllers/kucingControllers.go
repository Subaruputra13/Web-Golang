package controllers

import (
	"Web-Golang/config"
	"Web-Golang/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetKucing(c echo.Context) error {
	kucings := []models.Kucing{}

	err := config.DB.Find(&kucings).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "gagal mendapatkan data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Mendapatkan data",
		"data":    kucings,
	})

}

func PostKucing(c echo.Context) error {
	kucings := models.Kucing{}

	c.Bind(&kucings)

	err := config.DB.Create(&kucings).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Gagal melakukan decode",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil",
		"data":   kucings,
	})
}
