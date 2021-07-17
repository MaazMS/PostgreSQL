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

	users := []UserInfo{
		{
			First_name: "a",
			Last_name:  "a",
			Age:        1,
		},
		{
			First_name: "b",
			Last_name:  "b",
			Age:        2,
		},
		{
			First_name: "c",
			Last_name:  "c",
			Age:        3,
		},
	}

	for _, user := range users {
		db.Create(&user)
		fmt.Println("record inserting")
	}

	// Get the first record ordered by primary key,  SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&users)

	// Get one record, no specified order , SELECT * FROM users LIMIT 1;
	db.Take(&users)

	// Get last record, ordered by primary key desc, SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&users)

	// Retrieving objects with primary key , // SELECT * FROM users WHERE id = 1;
	db.First(&users, 5)

	// Get all records, SELECT * FROM users;
	db.Find(&users)

	//String Conditions
	// Get first matched record,  SELECT * FROM users WHERE first_name = 'a' ORDER BY id LIMIT 1;
	db.Where("first_name = ?", "a").First(&users)

	// Struct & Map Conditions, SELECT * FROM users WHERE first_name = "b" AND age = 2 ORDER BY id LIMIT 1;
	db.Where(&UserInfo{First_name: "b", Age: 2}).First(&users)

	//Selecting Specific Fields, SELECT name, age FROM users;
	db.Select("first_name", "age").Find(&users)

}
