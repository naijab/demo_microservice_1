package main

import (
	"user_service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user", user.GetUser)

	// Start server
	e.Logger.Fatal(e.Start(":1331"))
}
