package structure

import (
	"goproject/constants"
)

type Wearable interface {
	Equip(itemType constants.ItemType)
	Disarm()
	IsWearable(itemType constants.ItemType) bool
}

type Location interface {
	Move(direction constants.Direction) *Location
	ShowInfo()
}
