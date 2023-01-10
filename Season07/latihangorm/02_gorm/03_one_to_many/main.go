package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
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
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})
	db.Migrator().DropTable(&Calendar{})
	db.Migrator().CreateTable(&Calendar{})
	db.Migrator().DropTable(&Appointment{})
	db.Migrator().CreateTable(&Appointment{})

	db.Save(&User{
		Username:  "Iketutg",
		FirstName: "IKetut",
		LastName:  "Gunawan",
		Calendar: Calendar{
			Name: "Standup Meeting",
			Appointment: []Appointment{
				{Subject: "Interview"},
				{Subject: "Seminar"},
			},
		},
	})

}

// one to one
type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

// one to many
type Calendar struct {
	gorm.Model
	Name        string
	UserID      uint
	Appointment []Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	length      uint
	CalendarID  uint
}
