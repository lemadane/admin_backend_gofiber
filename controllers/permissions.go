package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/models"
)

// AllPermissions retrieves all permissions from the database and returns them as JSON.
func AllPermissions(context *fiber.Ctx) error {
	permissions := make([]models.Permission, 0)
	db.Session().Find(&permissions)
	return context.JSON(permissions)
}
