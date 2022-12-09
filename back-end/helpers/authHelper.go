package helpers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CheckUserType(ctx *fiber.Ctx, role string) (err error) {
	userType := ctx.Locals("user_type").(string)
	if userType != role {
		err = errors.New("unauthorized access")
		return
	}

	return
}

func CheckUserTypeAndID(ctx *fiber.Ctx, role string, id uint) (err error) {
	userType := ctx.Locals("user_type").(string)
	userID := ctx.Locals("user_id").(uint)

	if userType != role || userID != id {
		err = errors.New("unauthorized access")
		return
	}

	return
}
