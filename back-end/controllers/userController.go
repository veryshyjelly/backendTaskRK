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
		// Check if the same student is asking for the data
		if c.Locals("roll_no") != rollNo {
			c.Status(fiber.StatusUnauthorized)
			c.JSON(fiber.Map{"message": "unauthorized"})
			return
		}
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

type student struct {
	Name    *string `json:"name"`
	RollNo  *string `json:"roll_no"`
	BlockNo *string `json:"block_no"`
	RoomNo  *string `json:"room_no"`
	Contact *string `json:"contact"`
	Email   *string `json:"email"`
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

		var studentsList []student = make([]student, len(students))
		for i, stud := range students {
			studentsList[i] = student{
				Name:    stud.Name,
				RollNo:  stud.RollNo,
				BlockNo: stud.BlockNo,
				RoomNo:  stud.RoomNo,
				Contact: stud.Contact,
				Email:   stud.Email,
			}
		}

		c.Status(fiber.StatusOK)
		c.JSON(fiber.Map{"message": "success", "data": studentsList})
		return
	}
}

func UpdateStudent() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userType := c.Locals("user_type")
		if userType == "student" {
			var student models.Student
			if err = c.BodyParser(&student); err != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": "invalid request"})
				return
			}
			var foundStudent models.Student
			database.DB.WithContext(ctx).Model(&models.Student{}).First(&foundStudent, "roll_no = ?", student.RollNo)
			foundStudent.Contact = student.Contact

			database.DB.WithContext(ctx).Model(&models.Student{}).Where("roll_no = ?", student.RollNo).Updates(foundStudent)

			c.Status(fiber.StatusOK)
			c.JSON(fiber.Map{"message": "success", "data": foundStudent})
		} else if userType == "admin" {
			var student models.Student
			if err = c.BodyParser(&student); err != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": "invalid request"})
				return
			}
			database.DB.WithContext(ctx).Model(&models.Student{}).Where("roll_no = ?", student.RollNo).Updates(student)

			c.Status(fiber.StatusOK)
			c.JSON(fiber.Map{"message": "success", "data": student})
		} else {
			c.Status(fiber.StatusUnauthorized)
			c.JSON(fiber.Map{"message": "unauthorized"})
		}
		return
	}
}
