package models

// permissions is many2many relationship and table name of many to many relation is role_permissions
// Role have Id, Name, and Permissions of Role.
type Role struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
