package data

import (
	"goproject/constants"
	"goproject/structure"
)

func GetDungeonRowLen(stageNum constants.StageNum) int {
	switch stageNum {
	case constants.Stage0:
		return 2
	case constants.Stage2:
		return 10
	default:
		panic("Invalid Dungeon Level")
	}
}

func GetDungeonColLen(stageNum constants.StageNum) int {
	switch stageNum {
	case constants.Stage0:
		return 2
	case constants.Stage2:
		return 10
	default:
		panic("Invalid Dungeon Level")
	}
}

func GetDungeonRoomPositions(stageNum constants.StageNum) *[]structure.Position {
	switch stageNum {
	case constants.Stage0:
		return &[]structure.Position{
			{Row: 0, Col: 0}, {Row: 0, Col: 1},
			{Row: 1, Col: 0}, {Row: 1, Col: 1},
		}
	case constants.Stage2:
		return &[]structure.Position{
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
	default:
		panic("Invalid Dungeon Level")
	}
}

func GetDungeonDoorPositionAndType(stageNum constants.StageNum) *[]structure.DoorPositionAndType {
	switch stageNum {
	case constants.Stage0:
		return &[]structure.DoorPositionAndType{
			{RoomPosition: structure.Position{Row: 0, Col: 0}, Direction: constants.East, DoorType: constants.WoodDoor},
		}
	case constants.Stage2:
		return &[]structure.DoorPositionAndType{
			{RoomPosition: structure.Position{Row: 2, Col: 1}, Direction: constants.West, DoorType: constants.LockedDoor},
			{RoomPosition: structure.Position{Row: 3, Col: 4}, Direction: constants.West, DoorType: constants.WoodDoor},
			{RoomPosition: structure.Position{Row: 3, Col: 8}, Direction: constants.East, DoorType: constants.GlassDoor},
			{RoomPosition: structure.Position{Row: 6, Col: 2}, Direction: constants.West, DoorType: constants.GlassDoor},
			{RoomPosition: structure.Position{Row: 6, Col: 4}, Direction: constants.West, DoorType: constants.GlassDoor},
			{RoomPosition: structure.Position{Row: 8, Col: 6}, Direction: constants.South, DoorType: constants.WoodDoor},
		}
	default:
		panic("Invalid Dungeon Level")
	}
}

func GetDungeonStartPosition(stageNum constants.StageNum) structure.Position {
	switch stageNum {
	case constants.Stage0:
		return structure.Position{Row: 1, Col: 1}
	case constants.Stage2:
		return structure.Position{Row: 9, Col: 9}
	default:
		panic("Invalid Dungeon Level")
	}
}

func GetDungeonExitPosition(stageNum constants.StageNum) structure.Position {
	switch stageNum {
	case constants.Stage0:
		return structure.Position{Row: 0, Col: 0}
	case constants.Stage2:
		return structure.Position{Row: 2, Col: 1}
	default:
		panic("Invalid Dungeon Level")
	}
}
