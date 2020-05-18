package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func main() {
	initDB()
	var a string
	fmt.Scanln(&a)
	switch a {
	case "1":
		{
			print("1:入库")
		}
	case "2":
		{
			print("2:出库")
		}
	case "0":
		{
			print("0:退出")
			return
		}
	}
}

func initDB() {
	var db *gorm.DB
	db, err := gorm.Open("sqlite3", "./database.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	fmt.Println("connect to DB successfully!")
	DB = db
	DB.AutoMigrate(&Equipment{}, &Record{})
}
