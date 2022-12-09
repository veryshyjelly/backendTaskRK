package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veryshyvelly/task2-backend/controllers"
	"github.com/veryshyvelly/task2-backend/middleware"
)

func UserRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.Get("/student/:roll_no", controllers.GetStudent())
	incomingRoutes.Get("/students", controllers.GetAllStudents())
	incomingRoutes.Put("/student/:roll_no", controllers.UpdateStudent())
}
