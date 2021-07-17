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
	Id   uint
	Name string
	Book []Book `gorm:"many2many:user_book"`
}
type Book struct {
	Id    uint
	Title string
}

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the Database")
	}

	// take the address of 'UserInfo{}', it also use to create table
	// create table
	db.AutoMigrate(&UserInfo{}, &Book{})

	fmt.Println("many2many relationship created")

}

/*
one user have multiple book and multiple user have one book,
user_book is many to many table name
2021/07/17 09:21:58 /home/maaz/go/src/gorm.io/driver/postgres/migrator.go:161 SLOW SQL >= 200ms
[208.509ms] [rows:0] CREATE TABLE "user_infos" ("id" bigserial,"name" text,PRIMARY KEY ("id"))
Table created

*/
