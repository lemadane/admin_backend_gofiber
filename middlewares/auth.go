package middlewares

import (
	"github.com/lemadane/admin_backend_gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

// IsAuthenticated is a middleware function that checks if the user is authenticated.
// It retrieves the JWT token from the cookie and verifies its validity.
// If the token is invalid or missing, it returns an unauthorized status and a JSON response.
// Otherwise, it allows the request to proceed to the next middleware or route handler.
func IsAuthenticated(context *fiber.Ctx) error {
	cookie := context.Cookies("jwt")
	if _, err := utils.ParseJwt(cookie); err != nil {
		context.Status(fiber.StatusUnauthorized)
		return context.JSON(fiber.Map{
			"message": "Not authenticated",
		})
	}
	return context.Next()
}
