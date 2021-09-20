package routes

import (
	"net/http"
	"time"

	"codedolphin.io/users-service/models"
	"codedolphin.io/users-service/services"
	"github.com/gofiber/fiber/v2"
)

func initAuthRoutes() {
	authRouter := app.Group("/v1/users/auth")
	authRouter.Post("/register", register)
	authRouter.Post("/login", login)
	authRouter.Get("/logout", logout)
}

func register(c *fiber.Ctx) error {
	var userDto models.UserDTO
	if err := c.BodyParser(&userDto); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userDto.HashPassword()
	user := userDto.ToUser()
	res := models.DB.Create(user)
	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}
	return c.JSON(user)
}

func login(c *fiber.Ctx) error {
	var uldto models.UserLoginDTO
	if err := c.BodyParser(&uldto); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var user models.User
	res := models.DB.Where("email = ?", uldto.Email).First(&user)
	if res.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "No user with that email address was found!",
		})
	}
	passCorrect := uldto.VerifyPassword(user.Password)
	if !passCorrect {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Password incorrect!",
		})
	}
	token, tokenExp := services.GenerateJWT(user)
	cookie := fiber.Cookie{
		Name:     "session",
		Value:    token,
		Expires:  time.Unix(tokenExp, 0),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"token":  token,
		"expiry": tokenExp,
	})
}

func logout(c *fiber.Ctx) error {
	cookie := c.Cookies("session")
	if cookie == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Not logged in!",
		})
	}
	invalidCookie := fiber.Cookie{
		Name:     "session",
		Value:    cookie,
		Expires:  time.Now().Add(time.Second * -1),
		HTTPOnly: true,
	}
	c.Cookie(&invalidCookie)
	return c.JSON(fiber.Map{
		"message": "Logged out!",
	})
}
