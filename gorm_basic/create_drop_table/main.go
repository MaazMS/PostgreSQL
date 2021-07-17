package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p"
	dbname   = "users"
)

type UserInfo struct {
	Id         int
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

	// take the address of 'UserInfo{}', it also use to create table
	db.AutoMigrate(&UserInfo{})

	// create table
	db.Migrator().CreateTable(&UserInfo{})
	// Drop Table
	db.Migrator().DropTable(&UserInfo{})

	fmt.Println("Table created and Drop Table")

}
