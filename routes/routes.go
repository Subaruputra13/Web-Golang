package routes

import (
	"Web-Golang/controllers"
	"Web-Golang/middlewares"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// Init Echo
	e := echo.New()

	// Middlewares
	middlewares.LogMiddlewares(e)

	//
	e.GET("/home", controllers.HomeController)
	e.GET("/kucing/:type", controllers.GetKucing)
	e.POST("/kucing", controllers.PostKucing)

	// Grouping Routes
	d := e.Group("/api")
	// Basic Auth
	d.Use(mid.BasicAuth(middlewares.BasicAuthMiddleware))
	d.Use(middlewares.ServerHeader)
	d.GET("/dashboard", controllers.GetDashboard)
	return e
}
