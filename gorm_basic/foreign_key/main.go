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
	Book []Book `gorm:"foreignkey:UserInfoId"`
}
type Book struct {
	Id         uint
	UserInfoId uint
	Title      string
}

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the Database")
	}

	// take the address of 'UserInfo{}', it also use to create table
	// create table
	db.Migrator().CreateTable(&UserInfo{}, &Book{})

	fmt.Println("foreign key created")

}

/*
2021/07/17 09:13:50 /home/maaz/go/src/gorm.io/driver/postgres/migrator.go:161 SLOW SQL >= 200ms
[377.269ms] [rows:0] CREATE TABLE "user_infos" ("id" bigserial,"name" text,PRIMARY KEY ("id"))

2021/07/17 09:13:50 /home/maaz/go/src/gorm.io/driver/postgres/migrator.go:161 SLOW SQL >= 200ms
[599.283ms] [rows:0] CREATE TABLE "books" ("id" bigserial,"user_info_id" bigint,"title" text,PRIMARY KEY ("id"),CONSTRAINT "fk_user_infos_book" FOREIGN KEY ("user_info_id") REFERENCES "user_infos"("id"))
Table created

*/
