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

	// take the address of 'UserInfo{}'
	db.AutoMigrate(&UserInfo{})

	user := UserInfo{
		Id: 2,
	}
	// take the address of variable UserInfo, also have id to which value to delete
	db.Delete(&user)
	fmt.Println("delete record")

}
