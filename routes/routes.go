package routes

import (
	"github.com/lemadane/admin_backend_gofiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/ping", controllers.Ping)
}
