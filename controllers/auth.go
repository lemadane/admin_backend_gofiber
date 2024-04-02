package controllers

import (
	"strconv"
	"time"

	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/models"
	"github.com/lemadane/admin_backend_gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

// Ping is a handler function that sends a "pong" response.
func Ping(context *fiber.Ctx) error {
	return context.SendString("pong")
}

// Register is a function that handles the registration of a new user.
// It receives a context object from the Fiber framework and returns an error.
// The function parses the request body and validates the password.
// If the passwords match, it creates a new user in the database and returns the user object as JSON.
// If the passwords do not match, it returns a JSON response with an error message.
func Register(context *fiber.Ctx) error {
	data := make(map[string]string)

	if err := context.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		context.Status(400)
		return context.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	user := models.User{
		Firstname: data["firstname"],
		Lastname:  data["lastname"],
		Email:     data["email"],
		PhoneNo:   data["phone_no"],
		Password:  data["password"],
		RoleId:    1,
	}
	user.SetPassword(data["password"])
	db.Session().Create(&user)
	return context.JSON(user)
}

// Login handles the login functionality.
// It receives a request context and attempts to authenticate the user.
// If the user is found and the password is correct, it returns the user details as JSON.
// If the user is not found or the password is incorrect, it returns an appropriate error message.
func Login(context *fiber.Ctx) error {
	data := make(map[string]string)
	if err := context.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	db.Session().Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		context.Status(400)
		return context.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if !user.IsCorrectPassword(data["password"]) {
		context.Status(400)
		return context.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	return context.JSON(user)
}

// Logout is a handler function that logs out the user by clearing the JWT cookie.
// It sets the "jwt" cookie value to an empty string and expires it.
// Returns a JSON response with a success message.
func Logout(context *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Duration(time.Now().Local().Day())),
		HTTPOnly: true,
	}
	context.Cookie(&cookie)
	return context.Status(fiber.StatusNoContent).Send(nil)
}

// UpdateInfo updates the user information based on the provided data.
// It parses the request body, retrieves the user ID from the JWT cookie,
// and updates the corresponding user record in the database.
// Finally, it returns the updated user information as a JSON response.
func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJwt(cookie)
	userId, _ := strconv.Atoi(*id)
	user := models.User{
		Id:        uint(userId),
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		PhoneNo:   data["phone_no"],
		Email:     data["email"],
	}
	db.Session().Model(&user).Updates(user)
	return c.JSON(user)
}

// UpdatePassword updates the password of a user.
// It expects a JSON object containing the new password and password confirmation in the request body.
// If the passwords do not match, it returns a JSON response with a status code of 422 and an error message.
// It retrieves the user ID from the JWT cookie and updates the password for the corresponding user in the database.
// Finally, it returns a JSON response with the updated user object.
func UpdatePassword(c *fiber.Ctx) error {
	var data = make(map[string]string)
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJwt(cookie)
	userId, _ := strconv.Atoi(*id)
	user := models.User{
		Id: uint(userId),
	}
	user.SetPassword(data["password"])
	db.Session().Model(&user).Updates(user)
	return c.JSON(user)
}
