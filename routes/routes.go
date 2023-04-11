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
	e.Use(mid.Static("views"))

	// Routes
	e.Use(mid.StaticWithConfig(mid.StaticConfig{
		Root:  "view",
		Index: "about.html",
	}))
	e.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(constants.SCREAT_JWT),
		TokenLookup:   "cookie:JwtCookie",
	}))
	e.GET("/home", controllers.GetHome)

	l := e.Group("/login")
	l.GET("", controllers.Login)

	k := e.Group("/kucing")
	k.GET("/:type", controllers.GetKucing)
	k.POST("", controllers.PostKucing)

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
