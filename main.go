package main

import (
	"goproject/constants"
	"goproject/game/initialize"
	"goproject/game/terminal"
	"goproject/structure"
)

func gameStart(startRoom *structure.Room) {
	terminal.RunTerminal(startRoom)
}

func main() {

	rowLen := 6
	colLen := 7
	roomPositions := []structure.Position{
		{Row: 0, Col: 4}, {Row: 0, Col: 5}, {Row: 0, Col: 6},
		{Row: 1, Col: 3}, {Row: 1, Col: 4},
		{Row: 2, Col: 0}, {Row: 2, Col: 1}, {Row: 2, Col: 2}, {Row: 2, Col: 3},
		{Row: 3, Col: 1}, {Row: 3, Col: 3},
		{Row: 4, Col: 1}, {Row: 4, Col: 3}, {Row: 4, Col: 4},
		{Row: 5, Col: 1}, {Row: 5, Col: 4},
	}
	doorPositionAndType := []structure.DoorPositionAndType{
		{RoomPosition: structure.Position{Row: 0, Col: 4}, Direction: constants.South, DoorType: constants.WoodDoor},
		{RoomPosition: structure.Position{Row: 0, Col: 6}, Direction: constants.East, DoorType: constants.LockedDoor},
		{RoomPosition: structure.Position{Row: 2, Col: 1}, Direction: constants.East, DoorType: constants.GlassDoor},
		{RoomPosition: structure.Position{Row: 3, Col: 4}, Direction: constants.East, DoorType: constants.WoodDoor},
	}
	startPositon := structure.Position{Row: 5, Col: 1}
	endPosition := structure.Position{Row: 0, Col: 6}
	endDirection := constants.East

	startRoom := initialize.InitGameAndReturnStartPoint(rowLen, colLen, &roomPositions, &doorPositionAndType, startPositon, endPosition, endDirection)
	gameStart(startRoom)
}
