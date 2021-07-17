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
	Name    string `gorm:"type:VARCHAR(50)"`
	Email   string `gorm:"unique"`
	Phone   int    `gorm:"size:10"`
	Address string `gorm:"not null"`
	Salary  int    `gorm:"null"`
}

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the Database")
	}

	// take the address of 'UserInfo{}', it also use to create table
	db.Migrator().CreateTable(&UserInfo{})

	fmt.Println("Table created")

}

// [261.076ms] [rows:0] CREATE TABLE "user_infos" ("name" VARCHAR(50),"email" text UNIQUE,"phone" smallint,"address" text NOT NULL,"salary" bigint)
