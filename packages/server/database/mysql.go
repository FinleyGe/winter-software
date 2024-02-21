package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:@tcp(127.0.0.1:3306)/ss?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("DatabaseStartFailed: ", err)
	} else {
		fmt.Println("DatabaseStartSuccess")
	}

	err = autoMigrate(db)
	if err != nil {
		log.Fatal("DatabaseMigrateFailed: ", err)
	}

	DB = db
}
