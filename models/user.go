package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	PhoneNo   string `json:"phone_no"`
	Password  string `json:"password"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

// SetPassword sets the password for the user by hashing the provided password.
// It takes a string parameter `password` and updates the `Password` field of the `User` struct.
// If an error occurs during the password hashing process, it panics.
func (user *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
}

// IsCorrectPassword checks if the provided password matches the user's stored password.
// It uses bcrypt to compare the hashed password with the plain-text password.
// Returns true if the passwords match, false otherwise.
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Take retrieves a list of users from the database with the specified limit and offset.
// It returns the list of users as an interface{}.
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	products := make([]User, 0)
	db.Preload("Role").Offset(offset).Limit(limit).Find(&products)
	return products
}

// Count returns the total number of User records in the database.
func (user *User) Count(gormDb *gorm.DB) int64 {
	var count int64
	gormDb.Model(&User{}).Count(&count)
	return count
}
