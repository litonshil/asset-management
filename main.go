package main

import (
	"asset-management/connection"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	connection.ConnectDB()
	dbClient := connection.GetDB()

	InitRoutes(e, dbClient)

	e.Logger.Fatal(e.Start(":8080"))
}
