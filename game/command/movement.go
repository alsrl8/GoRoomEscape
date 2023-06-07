package command

import (
	"errors"
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

func Move(status *structure.Status, direction constants.Direction) {
	room := GetCurrentRoom(status)
	if !canMove(room, direction) {
		fmt.Println(constants.CanNotMoveWarning, constants.DirStringMap[direction])
		return
	}
	status.Location = (*status.Location).Move(direction)
}

func canMove(room *structure.Room, direction constants.Direction) bool {
	if room.Directions[direction] == nil {
		return false
	} else if room.Doors[direction] != nil && room.Doors[direction].Closed {
		return false
	}
	return true
}

func getNextRoomInDirection(room *structure.Room, direction constants.Direction) *structure.Room {
	return (*room.Directions[direction]).(*structure.Room)
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

func GetCounterDirection(direction constants.Direction) constants.Direction {
	return constants.DirectionList[(direction+2)%4]
}

func breakGlassDoor(room *structure.Room, inventory *structure.Inventory) {
	glassDoorDirection := findDoorDirectionByType(room, constants.GlassDoor)
	room.Doors[glassDoorDirection] = nil
	oppositeRoom := (*room.Directions[glassDoorDirection]).(*structure.Room)
	if oppositeRoom != nil {
		counterDirection := GetCounterDirection(glassDoorDirection)
		oppositeRoom.Doors[counterDirection] = nil
	}
	fmt.Println(constants.SucceedBreakingGlassDoor, constants.DirStringMap[glassDoorDirection])
}

func unlockLockedDoor(room *structure.Room, inventory *structure.Inventory) {
	lockedDoorDirection := findDoorDirectionByType(room, constants.LockedDoor)
	room.Doors[lockedDoorDirection].DoorType = constants.WoodDoor
	oppositeRoom := (*room.Directions[lockedDoorDirection]).(*structure.Room)
	if oppositeRoom != nil {
		counterDirection := GetCounterDirection(lockedDoorDirection)
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

func getNextRoomInfo(room *structure.Room, direction constants.Direction) string {
	if room.Doors[direction] != nil {
		doorType := room.Doors[direction].DoorType
		isClosed := room.Doors[direction].Closed
		return fmt.Sprintf(constants.DirectionInfo, constants.DoorTypeStringMap[doorType], constants.DoorCloseStateStringMap[isClosed])
	}

	if room.Directions[direction] == nil {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.EmptyRoom])
	}
}

func PrintLine() {
	fmt.Println(constants.LineDivider)
}

func ShowRoomInfo(room *structure.Room) {
	showObjectsInDirections(room)
	showItemsInRoom(room)
	if isMonsterExistInRoom(room) {
		PrintMonsterInRoom(room)
	}
}

func ShowUserNameAndStatus(status *structure.Status) {
	PrintLine()
	fmt.Printf(constants.PlayerStatus, status.Name, status.Attribute.Health, status.Attribute.Attack, status.Attribute.Defense)
}

func showObjectsInDirections(room *structure.Room) {
	PrintLine()
	for _, d := range constants.DirectionList {
		fmt.Printf(constants.DirectionInfoWithRoomInfo, constants.DirStringMap[d], constants.DirStringEngMap[d], getNextRoomInfo(room, d))
	}
}

func showItemsInRoom(room *structure.Room) {
	PrintLine()
	fmt.Printf("방에 있는 아이템 >>> ")
	for itemType, itemNum := range room.Items {
		if itemNum <= 0 {
			continue
		}
		fmt.Printf(constants.ItemTypeAndNum+" ", constants.ItemTypeStringMap[itemType], itemNum)
	}
	fmt.Println()
}

func ShowInventory(inventory *structure.Inventory) {
	PrintLine()
	fmt.Printf(constants.ItemInfoTitle)
	for itemType, itemNum := range *inventory {
		if itemNum == 0 {
			continue
		}
		fmt.Printf(constants.ItemTypeAndNum, constants.ItemTypeStringMap[itemType], itemNum)
	}
	fmt.Println()
}

func ShowMovableDirections(room *structure.Room) {
	PrintLine()
	fmt.Printf(constants.MovableDirectionTitle)
	for _, d := range constants.DirectionList {
		if !canMove(room, d) {
			continue
		}
		fmt.Printf(constants.DirectionInfo, constants.DirStringMap[d], constants.DirStringEngMap[d])
	}
	fmt.Println()
}

func ValidateDoorExist(room *structure.Room, doorType constants.DoorType) error {
	if findDoorInRoom(room, doorType) == nil {
		return errors.New(constants.NoSuchDoor)
	}
	return nil
}
