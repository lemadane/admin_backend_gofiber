package models

// Role represents a user role in the system.
type Role struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
