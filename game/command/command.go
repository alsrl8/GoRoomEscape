package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

func CanMove(room *structure.Room, direction constants.Direction) bool {
	if room.Directions[direction] == nil {
		return false
	} else if room.Doors[direction] != nil && room.Doors[direction].Closed {
		return false
	}
	return true
}

func Move(room *structure.Room, direction constants.Direction) *structure.Room {
	return room.Directions[direction]
}

func findDoorDirectionByType(room *structure.Room, doorType constants.DoorType) constants.Direction {
	for _, direction := range constants.DirectionList {
		if room.Doors[direction] == nil {
			continue
		} else if room.Doors[direction].DoorType != doorType {
			continue
		}
		return direction
	}
	return constants.NoDirection
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

func IsItemsEnoughToOpenDoor(door *structure.Door, bag *structure.Inventory) bool {
	switch door.DoorType {
	case constants.GlassDoor:
		if (*bag)[constants.Hammer] == 0 {
			return false
		}
	case constants.LockedDoor:
		if (*bag)[constants.Key] == 0 {
			return false
		}
	}
	return true
}

func OpenDoor(door *structure.Door, bag *structure.Inventory) {
	switch door.DoorType {
	case constants.WoodDoor:
		door.Closed = false
	case constants.GlassDoor:
		fmt.Println(constants.CanNotOpenSuchDoor, constants.DoorTypeStringMap[constants.GlassDoor])
	case constants.LockedDoor:
		fmt.Println(constants.CanNotOpenSuchDoor, constants.DoorTypeStringMap[constants.LockedDoor])
	}
}

func IsItemInBag(bag *structure.Inventory, itemName string) bool {
	item, ok := constants.StringItemTypeMap[itemName]
	if !ok {
		return false
	} else if (*bag)[item] == 0 {
		return false
	}
	return true
}

func GetCounterDirection(direction constants.Direction) constants.Direction {
	return constants.DirectionList[(direction+2)%4]
}

func BreakGlassDoorAndReduceHammer(room *structure.Room, bag *structure.Inventory) {
	glassDoorDirection := findDoorDirectionByType(room, constants.GlassDoor)
	room.Doors[glassDoorDirection] = nil
	oppositeRoom := room.Directions[glassDoorDirection]
	if oppositeRoom != nil {
		counterDirection := GetCounterDirection(glassDoorDirection)
		oppositeRoom.Doors[counterDirection] = nil
	}
	(*bag)[constants.Hammer] -= 1
	fmt.Println(constants.SucceedBreakingGlassDoor, constants.DirStringMap[glassDoorDirection])
}

func UnlockLockedDoorAndReduceKey(room *structure.Room, bag *structure.Inventory) {
	lockedDoorDirection := findDoorDirectionByType(room, constants.LockedDoor)
	room.Doors[lockedDoorDirection].DoorType = constants.WoodDoor
	oppositeRoom := room.Directions[lockedDoorDirection]
	if oppositeRoom != nil {
		counterDirection := GetCounterDirection(lockedDoorDirection)
		if oppositeRoom.Doors[counterDirection] != nil {
			oppositeRoom.Doors[counterDirection].DoorType = constants.WoodDoor
		}
	}
	(*bag)[constants.Key] -= 1
	fmt.Println(constants.SucceedUnlockLockedDoor, constants.DirStringMap[lockedDoorDirection])
}
