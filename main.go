package main

import (
	"goproject/constants"
	"goproject/game/initialize"
	"goproject/game/terminal"
	"goproject/structure"
)

func gameStart(status *structure.Status) {
	terminal.SetUserInputAsUserName(status)
	terminal.RunTerminal(status)
}

func main() {
	rowLen := 10
	colLen := 10
	roomPositions := []structure.Position{
		{Row: 0, Col: 4},
		{Row: 1, Col: 4},
		{Row: 2, Col: 1}, {Row: 2, Col: 4},
		{Row: 3, Col: 1}, {Row: 3, Col: 3}, {Row: 3, Col: 4}, {Row: 3, Col: 5}, {Row: 3, Col: 6}, {Row: 3, Col: 7}, {Row: 3, Col: 8}, {Row: 3, Col: 9},
		{Row: 4, Col: 1}, {Row: 4, Col: 4}, {Row: 4, Col: 6},
		{Row: 5, Col: 1}, {Row: 5, Col: 4}, {Row: 5, Col: 6},
		{Row: 6, Col: 1}, {Row: 6, Col: 2}, {Row: 6, Col: 3}, {Row: 6, Col: 4}, {Row: 6, Col: 6},
		{Row: 7, Col: 6}, {Row: 7, Col: 7}, {Row: 7, Col: 8}, {Row: 7, Col: 9},
		{Row: 8, Col: 6}, {Row: 8, Col: 9},
		{Row: 9, Col: 6}, {Row: 9, Col: 9},
	}
	doorPositionAndType := []structure.DoorPositionAndType{
		{RoomPosition: structure.Position{Row: 2, Col: 1}, Direction: constants.West, DoorType: constants.LockedDoor},
		{RoomPosition: structure.Position{Row: 3, Col: 4}, Direction: constants.West, DoorType: constants.WoodDoor},
		{RoomPosition: structure.Position{Row: 3, Col: 8}, Direction: constants.East, DoorType: constants.GlassDoor},
		{RoomPosition: structure.Position{Row: 6, Col: 2}, Direction: constants.West, DoorType: constants.GlassDoor},
		{RoomPosition: structure.Position{Row: 6, Col: 4}, Direction: constants.West, DoorType: constants.GlassDoor},
		{RoomPosition: structure.Position{Row: 8, Col: 6}, Direction: constants.South, DoorType: constants.WoodDoor},
	}
	startPosition := structure.Position{Row: 9, Col: 9}
	endPosition := structure.Position{Row: 2, Col: 1}
	endDirection := constants.West
	itemPositionAndType := []structure.ItemPositionAndType{
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
	monsterWithPosition := []structure.MonsterWithPosition{
		{
			RoomPosition: structure.Position{Row: 3, Col: 4},
			Monster: structure.Monster{
				MonsterType: constants.Squirrel,
				Attribute: structure.Attribute{
					Health:  50,
					Attack:  5,
					Defense: 0,
				},
				DropItem: []structure.DropItem{
					{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.7, Num: 1}},
				},
			},
		},
		{
			RoomPosition: structure.Position{Row: 4, Col: 6},
			Monster: structure.Monster{
				MonsterType: constants.Squirrel,
				Attribute: structure.Attribute{
					Health:  50,
					Attack:  5,
					Defense: 0,
				},
				DropItem: []structure.DropItem{
					{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.7, Num: 1}},
				},
			},
		},
		{
			RoomPosition: structure.Position{Row: 6, Col: 1},
			Monster: structure.Monster{
				MonsterType: constants.Deer,
				Attribute: structure.Attribute{
					Health:  100,
					Attack:  10,
					Defense: 5,
				},
				DropItem: []structure.DropItem{
					{ItemType: constants.Key, DropPercentage: structure.DropPercentage{Percentage: 1.0, Num: 1}},
				},
			},
		},
		{
			RoomPosition: structure.Position{Row: 6, Col: 4},
			Monster: structure.Monster{
				MonsterType: constants.Rabbit,
				Attribute: structure.Attribute{
					Health:  70,
					Attack:  7,
					Defense: 3,
				},
				DropItem: []structure.DropItem{
					{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.5, Num: 1}},
					{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.3, Num: 2}},
				},
			},
		},
	}
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
	boxPositionAndDropItems := []structure.BoxPositionAndDropItem{
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

	status := initialize.InitGameAndReturnStatus(rowLen, colLen, &roomPositions, &doorPositionAndType, startPosition, endPosition, endDirection, &itemPositionAndType, &monsterWithPosition, &boxPositionAndDropItems)
	gameStart(status)
}
