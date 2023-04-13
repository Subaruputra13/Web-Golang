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
	e.Use(m.ServerHeader)
	e.Use(mid.StaticWithConfig(mid.StaticConfig{
		Root:  "views",
		Index: "about.html",
	}))

	// Routes
	e.GET("/home", controllers.GetHome, mid.JWTWithConfig(mid.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(constants.SCREAT_JWT),
		TokenLookup:   "cookie:JwtCookie",
	}))

	l := e.Group("/login")
	l.GET("", controllers.Login)

	k := e.Group("/kucing")
	k.GET("", controllers.GetKucing)
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
}
