package routes

import "github.com/gofiber/fiber/v2"

var app *fiber.App

func Run() {
	app = fiber.New()

	setupRoutes()

	app.Listen(":8000")
}

func setupRoutes() {
	initAuthRoutes()
	initUserRoutes()
}
