package command

import (
	"goproject/constants"
	"goproject/structure"
)

func CanMove(room *structure.Room, direction constants.Direction) bool {
	if room.Directions[direction] == nil {
		return false
	} else if room.Doors[direction] != nil && room.Doors[direction].IsClosed {
		return false
	}
	return true
}

func Move(room *structure.Room, direction constants.Direction) *structure.Room {
	return room.Directions[direction]
}

func FindDoorByName(room *structure.Room, doorName string) *structure.Door {
	for _, direction := range constants.DirectionList {
		if room.Doors[direction] == nil {
			continue
		}
		doorType := room.Doors[direction].DoorType
		if constants.DoorTypeStringMap[doorType] != doorName {
			continue
		}

		return room.Doors[direction]
	}
	return nil
}

func IsItemsEnoughToOpenDoor(door *structure.Door) bool {
	return true
}

func OpenDoor(door *structure.Door) bool {
	door.IsClosed = false
	return true
}
