package controllers

import (
	"Web-Golang/middlewares"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	// Login Cookie
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "subaru" && password == "123" {
		cookie := new(http.Cookie)
		cookie.Name = "session"
		cookie.Value = "admin"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		// Create Jwt
		token, err := middlewares.CreateToken()

		if err != nil {
			return c.String(http.StatusInternalServerError, "Login Gagal")
		}

		JwtCookie := new(http.Cookie)
		JwtCookie.Name = "JwtCookie"
		JwtCookie.Value = token
		JwtCookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(JwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"token":  token,
			"status": "Login Berhasil",
		})
	}

	return c.String(http.StatusInternalServerError, "Login Gagal")

}
