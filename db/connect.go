package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ormDb is a global variable that holds the connection to the database.
var ormDb *gorm.DB

// Connect establishes a connection to the database.
func Connect() {
	db, err := gorm.Open(
		mysql.Open(
			"root:mel@tcp(localhost:3306)/mysql?charset=utf8&parseTime=True&loc=Local",
		),
		&gorm.Config{},
	)
	if err != nil {
		panic(err.Error())
	}
	ormDb = db
}

// Session returns the global database connection.
func Session() *gorm.DB {
	return ormDb
}
