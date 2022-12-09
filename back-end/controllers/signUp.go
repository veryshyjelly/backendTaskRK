package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/veryshyvelly/task2-backend/database"
	"github.com/veryshyvelly/task2-backend/helpers"
	"github.com/veryshyvelly/task2-backend/models"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, errors.New("email or password is incoorect")
	}
	return true, nil
}

func SignUp(givenUserType string) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		// Check if the user is authorized
		if givenUserType == "admin" {
			var user models.Admin
			// Parse the body
			if err = c.BodyParser(&user); err != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": err.Error()})
				return
			}
			// Validate the user
			if validationErr := validate.Struct(user); validationErr != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": validationErr.Error()})
				return
			}
			// Check if the user already exists
			var count1, count2, count3 int64
			database.DB.WithContext(ctx).Model(&models.Admin{}).Where("contact = ?", user.Contact).Count(&count3)
			database.DB.WithContext(ctx).Model(&models.Admin{}).Where("email = ?", user.Email).Count(&count1)
			database.DB.WithContext(ctx).Model(&models.Admin{}).Where("faculty_id = ?", user.FacultyId).Count(&count2)
			if count1+count2+count3 > 0 {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": "email or phone number already exists"})
				return
			}
			// Hash the password
			var password = HashPassword(*user.Password)
			user.UpdatedAt = time.Now().Local()
			user.CreatedAt = time.Now().Local()
			user.Password = &password
			// Generate the tokens
			token, refreshToken, err := helpers.GenerateAllTokens(*user.Email, *user.Name, "admin", *user.FacultyId)
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				c.JSON(fiber.Map{"message": err.Error()})
				return err
			}
			user.Token = &token
			user.RefreshToken = &refreshToken
			// Save the user
			database.DB.WithContext(ctx).Create(&user)

			// Send the user as response
			c.Status(fiber.StatusOK)
			c.JSON(fiber.Map{"message": "success", "user": user})

		} else if givenUserType == "student" {
			var user models.Student
			// Parse the body
			if err = c.BodyParser(&user); err != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": err.Error()})
				return
			}
			// Validate the user
			if validationErr := validate.Struct(user); validationErr != nil {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": validationErr.Error()})
				return
			}
			// Check if the user already exists
			var count1, count2 int64
			database.DB.WithContext(ctx).Model(&models.Student{}).Where("email = ?", user.Email).Count(&count1)
			database.DB.WithContext(ctx).Model(&models.Student{}).Where("roll_no = ?", user.RollNo).Count(&count2)
			if count1+count2 > 0 {
				c.Status(fiber.StatusBadRequest)
				c.JSON(fiber.Map{"message": "email or roll number already exists"})
				return
			}
			// Hash the password
			var password = HashPassword(*user.Password)
			user.UpdatedAt = time.Now().Local()
			user.CreatedAt = time.Now().Local()
			user.Password = &password
			// Generate the tokens
			token, refreshToken, err := helpers.GenerateAllTokens(*user.Email, *user.Name, "student", *user.RollNo)
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				c.JSON(fiber.Map{"message": err.Error()})
				return err
			}
			user.Token = &token
			user.RefreshToken = &refreshToken
			// Save the user
			database.DB.WithContext(ctx).Create(&user)
			// Send the user as response
			c.Status(fiber.StatusOK)
			c.JSON(fiber.Map{"message": "success", "user": user})

		} else {
			// If the user type is invalid
			c.Status(fiber.StatusBadRequest)
			c.JSON(fiber.Map{"message": "invalid user type"})
		}

		return
	}
}
