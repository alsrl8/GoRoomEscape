package structure

import (
	"goproject/constants"
)

type Room struct {
	Doors      map[constants.Direction]*Door
	Directions map[constants.Direction]*Room
	GoalFlag   bool
	Items      map[constants.ItemType]int
}

type Door struct {
	Closed   bool
	DoorType constants.DoorType
}

type Item struct {
	ItemType constants.ItemType
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

type ItemPositionAndType struct {
	RoomPosition Position
	ItemType     constants.ItemType
}

type Bag map[constants.ItemType]int

type DropPercentage struct {
	percentage int
	num        int
}
