package database

import (
	"../models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p"
	dbname   = "users"
)

// Connect of database using postgresql
func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//psqlconn := "host=localhost user=postgres password=p dbname=users port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the Database")
	}
	fmt.Println("Connection to the Database created")

	DB = database

	// AutoMigrate run auto migration for given models
	//database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})

}
