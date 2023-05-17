package main

import (
	"goproject/constants"
	"goproject/game/initialize"
	"goproject/game/terminal"
	"goproject/structure"
)

func gameStart(status *structure.Status) {
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
	// boxItemPercentageMap := map[constants.ItemType]map[int]float32{
	// 	constants.WoodSword:    {1: 0.2},
	// 	constants.IronSword:    {1: 0.15},
	// 	constants.LeatherCloth: {1: 0.07},
	// 	constants.LeatherPants: {1: 0.08},
	// 	constants.LeatherHat:   {1: 0.1},
	// 	constants.HealPotion:   {1: 0.15, 2: 0.1, 3: 0.05},
	// 	constants.Nothing:      {1: 0.1},
	// }

	status := initialize.InitGameAndReturnStatus(rowLen, colLen, &roomPositions, &doorPositionAndType, startPosition, endPosition, endDirection, &itemPositionAndType)
	gameStart(status)
}
