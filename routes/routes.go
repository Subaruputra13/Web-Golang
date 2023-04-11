package routes

import (
	"Web-Golang/constants"
	"Web-Golang/controllers"
	m "Web-Golang/middlewares"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// Init Echo
	e := echo.New()

	// Middlewares
	m.LogMiddlewares(e)

	// Routes
	e.GET("/home", controllers.GetHome, mid.JWT([]byte(constants.SCREAT_JWT)))
	e.GET("/login", controllers.Login)
	e.GET("/kucing/:type", controllers.GetKucing)
	e.POST("/kucing", controllers.PostKucing)

	// Grouping Routes
	d := e.Group("/api")
	// Basic Auth
	d.Use(m.ServerHeader)
	d.Use(mid.BasicAuth(m.BasicAuthMiddleware))
	d.GET("/dashboard", controllers.GetDashboard)

	// Cookie
	co := e.Group("/cookie")
	co.Use(m.CheckCooki)
	co.GET("/main", controllers.GetCookie)
	return e

	// JWT
}
