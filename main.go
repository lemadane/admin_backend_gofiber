package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/routes"
)

func main() {
	db.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":5000")
}