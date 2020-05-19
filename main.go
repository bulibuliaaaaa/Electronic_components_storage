package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

var DB *gorm.DB

func main() {
	initDB()
	CheckStock()
	var a string
	for true {
		showMenu()
		fmt.Scanln(&a)
		switch a {
		case "1":
			{
				in()
				break
			}
		case "2":
			{
				out()
				break
			}

		case "3":
			{
				addEquipment()
				break
			}
		case "4":
			{
				GetSummary()
				break
			}
		case "5":
			{
				FindByType()
				break
			}
		case "0":
			{
				fmt.Println("退出")
				return
			}
		}
		a = ""
	}
}

func showMenu() {
	fmt.Println("1:入库")
	fmt.Println("2:出库")
	fmt.Println("3:添加/编辑元件")
	fmt.Println("4:概要")
	fmt.Println("5:按种类查找")
	fmt.Println("0:退出")
}

func initDB() {
	var db *gorm.DB
	db, err := gorm.Open("sqlite3", "./database.db")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	fmt.Println("connect to DB successfully!")
	DB = db
	DB.AutoMigrate(&Equipment{}, &Record{})
}

func CheckStock() {
	data := make([]Equipment, 0)
	DB.Model(&Equipment{}).Find(&data)
	for _, v := range data {
		if v.Count < 10 {
			fmt.Println(v.Name + v.Code + "库存不足10件")
		}
	}
}

func GetSummary() {
	Equipments := make([]Equipment, 0)
	DB.Model(&Equipment{}).Find(Equipments)
	data := make([]Record, 0)
	DB.Model(&Record{}).Find(&data)
	for _, v := range Equipments {
		DB.Model(&Record{}).Where(&Record{
			EquipmentID: v.ID,
		}).Find(&data)
		Consume := 0
		for _, v2 := range data {
			Consume = Consume + v2.Amount
		}
		price, _ := strconv.Atoi(v.Price)
		fmt.Println(v.Name + v.Code + "消耗了" + string(Consume) + "件,共" + string(Consume*price) + "元,现有" + string(v.Count) + "件，共" + string(v.Count*price) + "元")
	}
}

func FindByType() {
	Equipments := make([]Equipment, 0)
	DB.Model(&Equipment{}).Find(Equipments)
	Types := make(map[string]bool)
	for _, v := range Equipments {
		Types[v.Type] = true
	}
	for k, _ := range Types {
		fmt.Println(k)
	}
	var type2 string
	fmt.Scanln(&type2)
	DB.Model(&Equipment{}).Where(&Equipment{
		Type: type2,
	}).Find(Equipments)
	for _, v := range Equipments {
		fmt.Println(v.Name)
	}
}
