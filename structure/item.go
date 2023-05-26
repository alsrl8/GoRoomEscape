package structure

import "goproject/constants"

type Item struct {
	ItemType constants.ItemType
}

type ItemPositionAndType struct {
	RoomPosition Position
	ItemType     constants.ItemType
}

type DropPercentage struct {
	Percentage float64
	Num        int
}

type DropItem struct {
	ItemType       constants.ItemType
	DropPercentage DropPercentage
}
