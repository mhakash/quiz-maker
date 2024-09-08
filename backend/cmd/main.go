package main

import (
	"quiz/internal/config"
	"quiz/internal/models"
	"quiz/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db, err := config.InitializeDb()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db.AutoMigrate(&models.Quiz{}, &models.Tag{})

	routes.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start("127.0.0.1:1234"))
}
