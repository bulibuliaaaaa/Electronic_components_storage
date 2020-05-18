package main

import "github.com/jinzhu/gorm"

type Equipment struct {
	gorm.Model
	Name   string
	Price  string
	Count  int
	Type   string
	Remark string
}

type Record struct {
	gorm.Model
	Staff       string
	EquipmentID uint
	Amount      int
	Usage       string
}
