package command

import (
	"errors"
	"fmt"
	"goproject/constants"
	"goproject/game/initialize"
	"goproject/structure"
)

func Move(status *structure.Status, direction constants.Direction) {
	if !canMove(status.Location, direction) {
		fmt.Println(constants.CanNotMoveWarning, constants.DirStringMap[direction])
		return
	}
	status.Location = status.Location.Move(direction)
}

func canMove(location structure.Location, direction constants.Direction) bool {
	return location.CanMove(direction)
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

func findDoorInRoom(room *structure.Room, doorType constants.DoorType) *structure.Door {
	for _, direction := range constants.DirectionList {
		if room.Doors[direction] == nil {
			continue
		} else if doorType != room.Doors[direction].DoorType {
			continue
		}
		return room.Doors[direction]
	}
	return nil
}

func openDoor(door *structure.Door) {
	switch door.DoorType {
	case constants.WoodDoor:
		door.Closed = false
	case constants.GlassDoor:
		fmt.Println(constants.CanNotOpenSuchDoor, constants.DoorTypeStringMap[constants.GlassDoor])
	case constants.LockedDoor:
		fmt.Println(constants.CanNotOpenSuchDoor, constants.DoorTypeStringMap[constants.LockedDoor])
	}
}

func breakGlassDoor(room *structure.Room) {
	glassDoorDirection := findDoorDirectionByType(room, constants.GlassDoor)
	room.Doors[glassDoorDirection] = nil
	oppositeRoom := room.Directions[glassDoorDirection].(*structure.Room)
	if oppositeRoom != nil {
		counterDirection := constants.GetCounterDirection(glassDoorDirection)
		oppositeRoom.Doors[counterDirection] = nil
	}
	fmt.Println(constants.SucceedBreakingGlassDoor, constants.DirStringMap[glassDoorDirection])
}

func unlockLockedDoor(room *structure.Room) {
	lockedDoorDirection := findDoorDirectionByType(room, constants.LockedDoor)
	room.Doors[lockedDoorDirection].DoorType = constants.WoodDoor
	oppositeRoom := room.Directions[lockedDoorDirection].(*structure.Room)
	if oppositeRoom != nil {
		counterDirection := constants.GetCounterDirection(lockedDoorDirection)
		if oppositeRoom.Doors[counterDirection] != nil {
			oppositeRoom.Doors[counterDirection].DoorType = constants.WoodDoor
		}
	}
	fmt.Println(constants.SucceedUnlockLockedDoor, constants.DirStringMap[lockedDoorDirection])
}

func OpenDoorByName(room *structure.Room, doorType constants.DoorType) {
	door := findDoorInRoom(room, doorType)
	if !door.Closed {
		fmt.Println(constants.AlreadyOpenDoor, constants.DoorTypeStringMap[doorType])
		return
	}

	openDoor(door)
}

func CloseDoorByName(room *structure.Room, doorType constants.DoorType) {
	door := findDoorInRoom(room, doorType)
	if door.Closed {
		fmt.Println(constants.AlreadyClosedDoor, constants.DoorTypeStringMap[doorType])
		return
	}

	door.Closed = true
}

func PrintLine() {
	fmt.Println(constants.LineDivider)
}

func ShowUserNameAndStatus(status *structure.Status) {
	PrintLine()
	fmt.Printf(constants.PlayerStatus, status.Name, status.Attribute.Health, status.Attribute.Attack, status.Attribute.Defense)
}

func ShowInventory(inventory structure.Inventory) {
	PrintLine()
	fmt.Printf(constants.ItemInfoTitle)
	for itemType, itemNum := range inventory {
		if itemNum == 0 {
			continue
		}
		fmt.Printf(constants.ItemTypeAndNum, constants.ItemTypeStringMap[itemType], itemNum)
	}
	fmt.Println()
}

func ValidateDoorExist(room *structure.Room, doorType constants.DoorType) error {
	if findDoorInRoom(room, doorType) == nil {
		return errors.New(constants.NoSuchDoor)
	}
	return nil
}

func EnterDungeon(status *structure.Status) *structure.Status {
	initialize.InitDungeon(status, 0)
	return status
}
