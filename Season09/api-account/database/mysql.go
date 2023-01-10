package database

import (
	"api-account/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "192.168.0.77"
	port     = 3306
	user     = "root"
	password = "password"
	dbname   = "account"
)

// go get gorm.io/driver/postgres
// go get gorm.io/gorm
var (
	err error
	db  *gorm.DB
)

func init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connection", err.Error())
	}
	db.Debug().Migrator().DropTable(&models.Acccount{})
	db.Debug().Migrator().CreateTable(&models.Acccount{})
}

func Migrate() {
	//db.Migrator().DropTable(&models.User{})
	//db.Migrator().DropTable(&models.Product{})
	//db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDb() *gorm.DB {
	return db
}
