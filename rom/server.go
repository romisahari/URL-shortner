package main

import (
	"rom/api"
	"rom/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.Connect()
	e.GET("/romina/server/1/:url", api.SaveURL)
	e.GET("/romina/server/2/:url", api.GetURL)
	e.Logger.Fatal(e.Start(":1323"))
}
