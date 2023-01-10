package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"latihancrudgorm/models"
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

func DbConnection() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Connection", err.Error())
	}

}

func Migrate() {
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Product{})
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDb() *gorm.DB {
	return db
}
