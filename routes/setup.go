package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func Run() {
	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setupRoutes()

	app.Listen(":8000")
}

func setupRoutes() {
	initAuthRoutes()
	initUserRoutes()
}
