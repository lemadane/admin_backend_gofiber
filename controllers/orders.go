package controllers

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/middlewares"
	"github.com/lemadane/admin_backend_gofiber/models"
	"github.com/lemadane/admin_backend_gofiber/utils"
)

func AllOrders(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "orders"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(context.Query("page", "1"))
	orderDto := utils.Paginate(db.Session(), &models.Order{}, page)
	return context.JSON(orderDto)
}

func Export(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "orders"); err != nil {
		return err
	}
	filePath := "./csv/orders.csv"
	if err := createFile(filePath); err != nil {
		return err
	}
	return context.Download(filePath)
}

func createFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var orders []models.Order
	db.Session().Preload("OrderItems").Find(&orders)
	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})
	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.Firstname + " " + order.Lastname,
			order.Email,
			"",
			"",
			"",
		}

		if err := writer.Write(data); err != nil {
			return err
		}

		for _, orderItem := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			}
			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}

	return nil
}
