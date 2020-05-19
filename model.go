package main

import (
	"github.com/jinzhu/gorm"
)

type Equipment struct {
	gorm.Model
	Name   string `addEquipment:"yes" zh:"名称"`
	Price  string `addEquipment:"yes" zh:"名称"`
	Count  int    `zh:"名称"`
	Type   string `addEquipment:"yes" zh:"名称"`
	Remark string `addEquipment:"yes" zh:"名称"`
	Code   string `addEquipment:"yes" zh:"名称"`
}

type Record struct {
	gorm.Model
	Staff       string
	EquipmentID uint
	Amount      int
	Usage       string
}
