package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/veryshyvelly/task2-backend/database"
	"github.com/veryshyvelly/task2-backend/helpers"
	"github.com/veryshyvelly/task2-backend/models"
)

// Login API to login a user (admin or student)
func Login(givenUserType string) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Check if user is admin or student
		if givenUserType == "admin" {
			var user models.Admin
			var foundUser models.Admin
			// Parse the body of the request
			if err = c.BodyParser(&user); err != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": err.Error()})
				return
			}
			// Validate the body of the request
			if validationErr := validate.Struct(user); validationErr != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": validationErr.Error()})
				return
			}
			// Check if the user exists in the database
			database.DB.WithContext(ctx).First(&foundUser, "faculty_id = ?", user.FacultyId)
			if foundUser.FacultyId == nil { // If user doesn't exist
				c.Status(fiber.StatusNotAcceptable)
				c.JSON(fiber.Map{"message": "id or password is incorrect"})
				return
			}
			// Check if the password is correct
			passwordIsCorrect, err := CheckPasswordHash(*user.Password, *foundUser.Password)
			if !passwordIsCorrect {
				c.Status(fiber.StatusNotAcceptable)
				c.JSON(fiber.Map{"message": "id or password is incorrect"})
				return err
			}
			// Generate tokens
			token, refreshToken, err := helpers.GenerateAllTokens(*foundUser.Email, "admin", *foundUser.Name, *foundUser.FacultyId)
			if err != nil { // If error in generating tokens
				c.Status(fiber.StatusInternalServerError)
				c.JSON(fiber.Map{"message": err.Error()})
				return err
			}
			// Update tokens in the database
			helpers.UpdateAllTokens(token, refreshToken, foundUser.ID, "admin")
			// Get the updated user from the database
			database.DB.WithContext(ctx).First(&foundUser, "id = ?", foundUser.ID)
			// Send the user as response
			c.Status(fiber.StatusOK)
			c.JSON(fiber.Map{"message": "success", "data": foundUser})

		} else if givenUserType == "student" {
			var user models.Student
			var foundUser models.Student
			// Parse the body of the request
			if err = c.BodyParser(&user); err != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": err.Error()})
				return
			}
			// Check if the user exists in the database
			database.DB.WithContext(ctx).First(&foundUser, "roll_no = ?", user.RollNo)
			if foundUser.RollNo == nil {
				c.Status(fiber.StatusNotAcceptable)
				c.JSON(fiber.Map{"message": "id or password is incorrect"})
				return
			}
			// Check if the password is correct
			passwordIsCorrect, err := CheckPasswordHash(*user.Password, *foundUser.Password)
			if !passwordIsCorrect { // If password is incorrect
				c.Status(fiber.StatusNotAcceptable)
				c.JSON(fiber.Map{"message": "id or password is incorrect"})
				return err
			}
			// Generate tokens
			token, refreshToken, err := helpers.GenerateAllTokens(*foundUser.Email, "student", *foundUser.Name, *foundUser.RollNo)
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				c.JSON(fiber.Map{"message": err.Error()})
				return err
			}
			// Update tokens in the database
			helpers.UpdateAllTokens(token, refreshToken, foundUser.ID, "student")
			// Get the updated user from the database
			database.DB.WithContext(ctx).First(&foundUser, "id = ?", foundUser.ID)
			// Send the user as response
			c.Status(fiber.StatusOK)
			c.JSON(fiber.Map{"message": "success", "data": foundUser})

		} else {
			c.Status(fiber.StatusBadRequest)
			c.JSON(fiber.Map{"message": "invalid user type"})
			return
		}

		return
	}
}
