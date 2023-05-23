package data

import (
	"goproject/constants"
	"goproject/structure"
)

var ItemTypeUsableMap = map[constants.ItemType]bool{
	constants.Hammer:     true,
	constants.Key:        true,
	constants.Box:        true,
	constants.HealPotion: true,
}

func GetItemPositionAndType() *[]structure.ItemPositionAndType {
	return &[]structure.ItemPositionAndType{
		{RoomPosition: structure.Position{Row: 0, Col: 4}, ItemType: constants.Hammer},
		{RoomPosition: structure.Position{Row: 0, Col: 4}, ItemType: constants.Box},
		{RoomPosition: structure.Position{Row: 2, Col: 1}, ItemType: constants.Key},
		{RoomPosition: structure.Position{Row: 3, Col: 6}, ItemType: constants.Box},
		{RoomPosition: structure.Position{Row: 3, Col: 9}, ItemType: constants.Box},
		{RoomPosition: structure.Position{Row: 6, Col: 4}, ItemType: constants.Hammer},
		{RoomPosition: structure.Position{Row: 7, Col: 6}, ItemType: constants.Hammer},
		{RoomPosition: structure.Position{Row: 7, Col: 9}, ItemType: constants.WoodSword},
		{RoomPosition: structure.Position{Row: 9, Col: 6}, ItemType: constants.Box},
	}
}

func GetBoxPositionAndDropItem() *[]structure.BoxPositionAndDropItem {
	boxDropItems := []structure.DropItem{
		{ItemType: constants.WoodSword, DropPercentage: structure.DropPercentage{Percentage: 0.2, Num: 1}},
		{ItemType: constants.IronSword, DropPercentage: structure.DropPercentage{Percentage: 0.15, Num: 1}},
		{ItemType: constants.LeatherCloth, DropPercentage: structure.DropPercentage{Percentage: 0.07, Num: 1}},
		{ItemType: constants.LeatherPants, DropPercentage: structure.DropPercentage{Percentage: 0.08, Num: 1}},
		{ItemType: constants.LeatherShoes, DropPercentage: structure.DropPercentage{Percentage: 0.1, Num: 1}},
		{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.15, Num: 1}},
		{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.1, Num: 2}},
		{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.05, Num: 3}},
	}
	return &[]structure.BoxPositionAndDropItem{
		{
			RoomPosition: structure.Position{Row: 0, Col: 4},
			DropItem:     &boxDropItems,
		},
		{
			RoomPosition: structure.Position{Row: 3, Col: 3},
			DropItem:     &boxDropItems,
		},
		{
			RoomPosition: structure.Position{Row: 3, Col: 9},
			DropItem:     &boxDropItems,
		},
		{
			RoomPosition: structure.Position{Row: 9, Col: 6},
			DropItem:     &boxDropItems,
		},
	}
}
