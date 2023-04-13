package main

import (
	"Web-Golang/config"
	"Web-Golang/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
