package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// user have Id, FirstName, LastName, Email, Password, RoleId, Role
type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

// SetPassword Generate new password
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

// ComparePassword
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (user *User) Take(db *gorm.DB, limit int, offest int) interface{} {
	var products []User
	db.Preload("Role").Offset(offest).Limit(limit).Find(&products)
	return products
}
