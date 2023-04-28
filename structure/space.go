package structure

import (
	"goproject/constants"
)

type Room struct {
	Doors      map[constants.Direction]*Door
	Directions map[constants.Direction]*Room
	IsGoal     bool
}

type Door struct {
	IsClosed bool
	DoorType constants.DoorType
}

type Position struct {
	Row int
	Col int
}

type DoorPositionAndType struct {
	RoomPosition Position
	Direction    constants.Direction
	DoorType     constants.DoorType
}
