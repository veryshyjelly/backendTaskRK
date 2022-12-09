package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/veryshyvelly/task2-backend/database"
	"github.com/veryshyvelly/task2-backend/models"
)

// GetStudent API to get a student by roll_no (user_id)
func GetStudent() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		rollNo := c.Params("roll_no")
		// Check if the user is authorized
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		// Get the user from the database
		var user models.Student
		database.DB.WithContext(ctx).Model(&models.Student{}).First(&user, "roll_no = ?", rollNo)
		// Check if the user exists
		if user.RollNo == nil {
			c.Status(fiber.StatusNotFound)
			c.JSON(fiber.Map{"message": "user not found"})
			return
		}
		// Send the user as response
		c.Status(fiber.StatusOK)
		c.JSON(fiber.Map{"message": "success", "data": user})

		return
	}
}

func GetAllStudents() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var students []models.Student
		database.DB.WithContext(ctx).Model(&models.Student{}).Find(&students)

		if len(students) == 0 {
			c.Status(fiber.StatusNotFound)
			c.JSON(fiber.Map{"message": "no students found"})
			return
		}

		c.Status(fiber.StatusOK)
		c.JSON(fiber.Map{"message": "success", "data": students})
		return
	}
}

func UpdateStudent() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		return
	}
}
