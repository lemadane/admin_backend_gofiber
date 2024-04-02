package controllers

import (
	"strconv"

	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/middlewares"
	"github.com/lemadane/admin_backend_gofiber/models"

	"github.com/gofiber/fiber/v2"
)

// AllRoles is a handler function that returns all roles.
// It checks if the user is authorized to access the "roles" resource,
// and then retrieves all roles from the database and returns them as JSON.
func AllRoles(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "roles"); err != nil {
		return err
	}
	roles := make([]models.Role, 0)
	db.Session().Find(&roles)
	return context.JSON(roles)
}

// CreateRole creates a new role.
// It checks if the user is authorized to perform this action using the "roles" permission.
// It parses the request body into a fiber.Map object representing the role.
// It creates the role in the database using the db.Session() method.
// Finally, it returns the created role as a JSON response.
func CreateRole(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "roles"); err != nil {
		return err
	}
	var role fiber.Map
	if err := context.BodyParser(&role); err != nil {
		return err
	}
	db.Session().Create(role)
	return context.JSON(role)
}

// UpdateRole updates a role in the system.
// It first checks if the user is authorized to perform the update operation.
// Then, it parses the request body into a roleDto variable.
// Next, it extracts the permissions from the roleDto and converts them into models.Permission objects.
// After that, it deletes the existing role permissions from the database for the given role ID.
// Finally, it updates the role in the database with the new information and returns the updated role as JSON.
// If any error occurs during the process, it is returned as an error response.
func UpdateRole(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "roles"); err != nil {
		return err
	}
	var roleDto fiber.Map
	if err := context.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))
	for i, permissionId := range list {
		id, _ := permissionId.(uint)
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	id, _ := strconv.Atoi(context.Params("id"))
	var result interface{}
	db.Session().Table("role_permissions").Where("role_id", id).Delete(&result)
	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}
	db.Session().Model(&role).Updates(role)
	return context.JSON(role)
}

// DeleteRole deletes a role based on the provided ID.
// It first checks if the user is authorized to perform this action using the "roles" permission.
// If the user is authorized, it retrieves the ID from the request parameters and converts it to an integer.
// Then, it creates a Role struct with the retrieved ID.
// Finally, it deletes the role from the database and returns a response with a status code of 204 (No Content).
func DeleteRole(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "roles"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(context.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	db.Session().Delete(&role)
	return context.Status(fiber.StatusNoContent).Send(nil)
}
