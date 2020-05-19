package main

import (
	"fmt"
)

func addEquipment() {
	var (
		Name   string
		Code   string
		Price  string
		Remark string
		Type   string
	)
	fmt.Println("输入器材名称")
	fmt.Scanln(&Name)
	fmt.Println("输入器材编号")
	fmt.Scanln(&Code)
	fmt.Println("输入器材价格")
	fmt.Scanln(&Price)
	fmt.Println("输入器材备注")
	fmt.Scanln(&Remark)
	fmt.Println("输入器材种类")
	fmt.Scanln(&Type)
	fmt.Println("器材:" + Name + "\n编号:" + Code + "\n单价:" + Price + "\n" + Remark + "你确认你的输入无误吗，如果正确，请输入yes")
	var status string
	fmt.Scanln(&status)
	if status == "yes" {
		DB.Model(&Equipment{}).Create(&Equipment{
			Name:   Name,
			Price:  Price,
			Type:   Type,
			Remark: Remark,
			Code:   Code,
		})
	} else {
		fmt.Println("录入已取消")
	}
	return
}

func in() {
	var Code string
	var count int
	var person string
	var Usage string
	fmt.Println("请输入你的姓名")
	fmt.Scanln(&person)
	fmt.Println("请输入入仓物品编号")
	fmt.Scanln(&Code)
	fmt.Println("请输入入仓物品数量")
	fmt.Scanln(&count)
	fmt.Println("请输入备注")
	fmt.Scanln(&Usage)
	fmt.Println(person + "入仓" + string(count) + "件" + Code + "你确认你的输入无误吗，如果正确，请输入yes")
	var status string
	fmt.Scanln(&status)
	EquipmentID := findEquipmentIDByCode(Code)
	if status == "yes" && EquipmentID != 0 {
		DB.Model(&Record{}).Create(&Record{
			Staff:       person,
			EquipmentID: EquipmentID,
			Amount:      count,
			Usage:       Usage,
		})
	} else {
		fmt.Println("录入已取消")
	}
	return
}

func out() {
	var Code string
	var count int
	var person string
	var Usage string
	fmt.Println("请输入你的姓名")
	fmt.Scanln(&person)
	fmt.Println("请输入出仓物品编号")
	fmt.Scanln(&Code)
	fmt.Println("请输入出仓物品数量")
	fmt.Scanln(&count)
	fmt.Println("请输入用途")
	fmt.Scanln(&Usage)
	fmt.Println(person + "出仓" + string(count) + "件" + Code + "你确认你的输入无误吗，如果正确，请输入yes")
	var status string
	fmt.Scanln(&status)
	EquipmentID := findEquipmentIDByCode(Code)
	if status == "yes" && EquipmentID != 0 {
		DB.Model(&Record{}).Create(&Record{
			Staff:       person,
			EquipmentID: EquipmentID,
			Amount:      count,
			Usage:       Usage,
		})
	} else {
		fmt.Println("录入已取消")
	}
	return
}

func findEquipmentIDByCode(Code string) uint {
	data := new(Equipment)
	DB.Where(&Equipment{Code: Code}).Find(&data)
	return data.ID
}
