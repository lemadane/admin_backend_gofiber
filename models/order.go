package models

import "gorm.io/gorm"

// Order represents an order in the system.
type Order struct {
	Id         uint        `json:"id"`
	Firstname  string      `json:"-"`
	Lastname   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Email      string      `json:"email"`
	Total      float32     `json:"total" gorm:"-"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

// OrderItem represents an item in an order.
type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

// Count returns the total number of records in the database for the given order.
func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(order).Count(&total)

	return total
}

// Take retrieves a list of orders from the database with the specified limit and offset.
// It calculates the total price for each order by summing the prices of all order items.
// It also concatenates the first name and last name to form the order's name.
// The retrieved orders are returned as a slice of Order structs.
func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)
	for i, _ := range orders {
		var total float32
		for _, item := range orders[i].OrderItems {
			total += item.Price * float32(item.Quantity)
		}
		orders[i].Name = orders[i].Firstname + " " + orders[i].Lastname
		orders[i].Total = total
	}
	return orders
}
