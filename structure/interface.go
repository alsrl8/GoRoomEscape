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
	GetLocationType() constants.LocationType
	GetObjectMap() map[constants.ObjectType]int
	Move(direction constants.Direction) Location
	ShowInfo()
	CanMove(direction constants.Direction) bool
	Connect(near Location, direction constants.Direction)
}
