package middlewares

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/models"
	"github.com/lemadane/admin_backend_gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

// IsAuthorized checks if the user is authorized to access a specific page.
// It takes a `context` object of type `*fiber.Ctx` and a `page` string as parameters.
// It returns an error if the user is unauthorized, otherwise it returns nil.
// The function first checks if the user has a valid JWT token in the cookie.
// If the token is valid, it retrieves the user ID from the token and fetches the user from the database.
// Then it retrieves the user's role and its associated permissions from the database.
// If the HTTP method is GET, it checks if the user has either "view"+page or "edit"+page permission.
// If the HTTP method is not GET, it only checks if the user has "edit"+page permission.
// If the user has the required permission, it returns nil indicating authorization.
// If the user is unauthorized, it sets the response status to 401 (Unauthorized) and returns an error.
func IsAuthorized(context *fiber.Ctx, page string) error {
	cookie := context.Cookies("jwt")
	id, err := utils.ParseJwt(cookie)
	if err != nil {
		context.Status(fiber.StatusUnauthorized)
		return context.JSON(fiber.Map{
			"message": "Not authorized",
		})
	}
	userId, _ := strconv.Atoi(*id)
	user := models.User{
		Id: uint(userId),
	}
	db.Session().Preload("Role").Find(&user)
	role := models.Role{
		Id: user.RoleId,
	}
	db.Session().Preload("Permissions").Find(&role)
	fmt.Println(role.Permissions)
	if context.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view"+page || permission.Name == "edit"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit"+page {
				return nil
			}
		}
	}
	context.Status(fiber.StatusUnauthorized)
	return errors.New("Not authorized")
}
