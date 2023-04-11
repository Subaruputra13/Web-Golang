package main

import "Web-Golang/routes"

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
