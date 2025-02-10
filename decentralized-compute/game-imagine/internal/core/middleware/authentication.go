package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUserEmail(ctx *fiber.Ctx, defaultIfNil ...string) string {
	userAddress := ctx.Locals(LocalAuthUserAddress)
	if userAddress == nil {
		if len(defaultIfNil) == 1 {
			return defaultIfNil[0]
		}
		return ""
	}
	return fmt.Sprintf("%s", userAddress)
}
