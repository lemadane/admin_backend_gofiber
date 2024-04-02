package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/models"
)

// Chart generates a chart of sales data.
// It retrieves sales data from the database and returns it as JSON.
func Chart(context *fiber.Ctx) error {
	var sales []models.Sales
	println(context.Method())

	db.Session().Raw(`
		SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d') as date, SUM(oi.price * oi.quantity) as sum
		FROM orders o
		JOIN order_items oi on o.id = oi.order_id
		GROUP BY date
	`).Scan(&sales)
	return context.JSON(sales)
}
