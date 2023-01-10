package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "192.168.0.77"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "latihandb"
)

// go get gorm.io/driver/postgres
// go get gorm.io/gorm
var (
	err error
	db  *gorm.DB
)

func dbConnection() *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		host, port, user, password, dbname)

	db, err2 := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err2 != nil {
		fmt.Println("Error Connection", err2.Error())
		panic(err2)
	}
	return db
}

func main() {

	db := dbConnection()
	db.Debug()
}
