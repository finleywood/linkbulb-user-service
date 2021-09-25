package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"linkbulb.io/users-service/models"
	"linkbulb.io/users-service/services"
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
