package structure

import (
	"goproject/constants"
)

type Status struct {
	Room      *Room
	Inventory *Inventory
	Equipment *Equipment
}

type Room struct {
	Doors      map[constants.Direction]*Door
	Directions map[constants.Direction]*Room
	GoalFlag   bool
	Items      map[constants.ItemType]int
	Monsters   map[constants.MonsterType]int
}

type Door struct {
	Closed   bool
	DoorType constants.DoorType
}

type Item struct {
	ItemType constants.ItemType
}

type Armor struct {
	Item
}

type Weapon struct {
	Item
}

type Equipment struct {
	Top       constants.ItemType
	Pants     constants.ItemType
	Shoes     constants.ItemType
	LeftHand  constants.ItemType
	RightHand constants.ItemType
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

type Inventory map[constants.ItemType]int

type DropPercentage struct {
	Percentage float32
	Num        int
}

type MonsterPositionAndType struct {
	RoomPosition Position
	MonsterType  constants.MonsterType
}

type Monster struct {
	Health   int
	Attack   int
	Defense  int
	DropItem []DropItem
}

type DropItem struct {
	ItemType       constants.ItemType
	DropPercentage DropPercentage
}
