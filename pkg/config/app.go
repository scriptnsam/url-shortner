package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	d, err := gorm.Open("mysql", "username:password@tcp(IP-Address:3306)/u_shortner?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

	// Check if the connection is successful
	if err := db.DB().Ping(); err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	fmt.Println("Connected to the database")
	return db
}

func GetDB() *gorm.DB {
	return db
}
