package structure

import "goproject/constants"

type Position struct {
	Row int
	Col int
}

type Room struct {
	Doors      map[constants.Direction]*Door
	Directions map[constants.Direction]*Room
	GoalFlag   bool
	Items      map[constants.ItemType]int
	Monster    *Monster
	NpcMap     map[constants.NpcType]int
	Place      constants.PlaceType
}

type Door struct {
	Closed   bool
	DoorType constants.DoorType
}

type DoorPositionAndType struct {
	RoomPosition Position
	Direction    constants.Direction
	DoorType     constants.DoorType
}
