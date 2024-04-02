package models

import "gorm.io/gorm"

// Entity represents an entity in the system.
type Entity interface {
	// Count returns the total count of entities in the database.
	Count(db *gorm.DB) int64

	// Take retrieves a subset of entities from the database.
	// It takes the number of entities to retrieve (limit) and the offset to start from.
	// It returns the retrieved entities.
	Take(db *gorm.DB, limit int, offset int) interface{}
}