package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veryshyvelly/task2-backend/controllers"
)

func AuthRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.Post("/student/login", controllers.Login("student"))
	incomingRoutes.Post("/student/signup", controllers.SignUp("student"))
	incomingRoutes.Post("/admin/login", controllers.Login("admin"))
}
