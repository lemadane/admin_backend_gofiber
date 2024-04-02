package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/middlewares"
	"github.com/lemadane/admin_backend_gofiber/models"
	"github.com/lemadane/admin_backend_gofiber/utils"
)

// AllUsers returns a list of all users.
// It checks if the user is authorized to access the "users" resource,
// retrieves the page number from the query parameters, and returns
// a JSON response with the paginated list of users.
func AllUsers(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "users"); err != nil {
		return err
	}
	pageNum, _ := strconv.Atoi(context.Query("page", "1"))
	return context.JSON(utils.Paginate(db.Session(), &models.User{}, pageNum))
}

// GetUser retrieves a user by ID and returns it as JSON.
// It first checks if the user is authorized to access this endpoint.
// If the user is authorized, it fetches the user from the database and returns it as JSON.
// If there is an error during authorization or fetching the user, it returns the error.
func GetUser(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(context.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	db.Session().Find(&user)
	return context.JSON(user)
}

// CreateUser creates a new user.
// It first checks if the user is authorized to create users.
// Then it parses the request body into a new User object.
// It sets the password for the user and creates the user in the database.
// Finally, it returns the created user as a JSON response.
func CreateUser(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "users"); err != nil {
		return err
	}
	user := new(models.User)
	if err := context.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword(user.Password)
	db.Session().Create(user)
	return context.JSON(user)
}

// UpdateUser updates a user's information based on the provided ID.
// It first checks if the user is authorized to perform the update.
// Then, it parses the request body into a User struct and updates the corresponding record in the database.
// Finally, it returns the updated user information as a JSON response.
func UpdateUser(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(context.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	if err := context.BodyParser(&user); err != nil {
		return err
	}
	db.Session().Model(&user).Updates(user)
	return context.JSON(user)
}

// DeleteUser deletes a user from the database.
// It first checks if the user is authorized to perform this action.
// Then it retrieves the user ID from the request parameters and converts it to an integer.
// It creates a new User struct with the retrieved ID.
// Finally, it deletes the user from the database and returns a JSON response with a success message.
func DeleteUser(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(context.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	db.Session().Delete(&user)
	return context.Status(fiber.StatusNoContent).Send(nil)
}
