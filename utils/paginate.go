package utils

import (
	"math"

	"github.com/lemadane/admin_backend_gofiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Paginate is a utility function that retrieves paginated data from the database.
// It takes a *gorm.DB instance, an entity model, and a page number as input.
// It returns a fiber.Map containing the paginated data and metadata.
// The data field in the fiber.Map contains the paginated data retrieved from the database.
// The meta field in the fiber.Map contains metadata about the pagination, including the total number of records,
// the current page number, and the last page number.
// The limit for each page is set to 15 records.
// The offset is calculated based on the page number and limit.
func Paginate(db *gorm.DB, entity models.Entity, pageNum int) fiber.Map {
	limit := 15
	offset := (pageNum - 1) * limit

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      pageNum,
			"last_page": int(math.Ceil(float64(total) / float64(limit))),
		},
	}
}
