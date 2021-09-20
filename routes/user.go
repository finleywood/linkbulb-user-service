package routes

import (
	"net/http"

	"codedolphin.io/users-service/models"
	"codedolphin.io/users-service/services"
	"github.com/gofiber/fiber/v2"
)

func initUserRoutes() {
	userRouter := app.Group("/v1/users")
	userRouter.Get("/user", getLoggedInUser)
}

func getLoggedInUser(c *fiber.Ctx) error {
	token := c.Cookies("session")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not logged in!",
		})
	}

	valid, uid := services.ValidateJWT(token)
	if !valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid session token!",
		})
	}

	var user models.User
	models.DB.First(&user, uid)
	return c.JSON(user)
}
