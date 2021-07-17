package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p"
	dbname   = "users"
)

type UserInfo struct {
	gorm.Model
	First_name string
	Last_name  string
	Age        int
}

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the Database")
	}

	// create table, take the address of 'UserInfo{}'
	db.Migrator().CreateTable(&UserInfo{})

	user := UserInfo{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
		First_name: "xyz",
		Last_name:  "abcd",
		Age:        100,
	}
	db.Create(&user)
	fmt.Println("gorm model created")

}
