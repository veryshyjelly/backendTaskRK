package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veryshyvelly/task2-backend/helpers"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		clientToken := string(c.Request().Header.Peek("token"))
		if clientToken == "" {
			c.Status(fiber.StatusUnauthorized)
			c.JSON(fiber.Map{"message": "no authorization token provided"})
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			c.JSON(fiber.Map{"message": err.Error()})
			return
		}

		c.Locals("roll_no", claims.RollNo)
		c.Locals("user_type", claims.UserType)
		c.Locals("name", claims.Name)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
