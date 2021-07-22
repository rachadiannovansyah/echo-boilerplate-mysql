package main

import (
	"learn-go-echo/database"
	"learn-go-echo/routes"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1515"))
}
