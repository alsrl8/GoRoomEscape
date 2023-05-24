package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

func Move(room *structure.Room, direction constants.Direction) *structure.Room {
	if !canMove(room, direction) {
		fmt.Println(constants.CanNotMoveWarning, constants.DirStringMap[direction], constants.DirStringEngMap[direction])
		return room
	} else if isMonsterExistInRoom(room) {
		fmt.Println(constants.MonsterExistsInTheRoom, constants.MonsterTypeStringMap[room.Monster.MonsterType])
		return room
	}
	return getNextRoomInDirection(room, direction)
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

func findDoorByName(room *structure.Room, doorName string) *structure.Door {
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

func breakGlassDoorAndReduceHammer(room *structure.Room, inventory *structure.Inventory) {
	glassDoorDirection := findDoorDirectionByType(room, constants.GlassDoor)
	room.Doors[glassDoorDirection] = nil
	oppositeRoom := room.Directions[glassDoorDirection]
	if oppositeRoom != nil {
		counterDirection := GetCounterDirection(glassDoorDirection)
		oppositeRoom.Doors[counterDirection] = nil
	}
	removeItemInInventory(inventory, constants.Hammer)
	fmt.Println(constants.SucceedBreakingGlassDoor, constants.DirStringMap[glassDoorDirection])
}

func unlockLockedDoorAndReduceKey(room *structure.Room, inventory *structure.Inventory) {
	lockedDoorDirection := findDoorDirectionByType(room, constants.LockedDoor)
	room.Doors[lockedDoorDirection].DoorType = constants.WoodDoor
	oppositeRoom := room.Directions[lockedDoorDirection]
	if oppositeRoom != nil {
		counterDirection := GetCounterDirection(lockedDoorDirection)
		if oppositeRoom.Doors[counterDirection] != nil {
			oppositeRoom.Doors[counterDirection].DoorType = constants.WoodDoor
		}
	}
	removeItemInInventory(inventory, constants.Key)
	fmt.Println(constants.SucceedUnlockLockedDoor, constants.DirStringMap[lockedDoorDirection])
}

func OpenDoorByName(room *structure.Room, inventory *structure.Inventory, doorName string) {
	door := findDoorByName(room, doorName)
	if door == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
		return
	} else if !door.Closed {
		fmt.Println(constants.AlreadyOpenDoor, doorName)
		return
	}

	openDoor(door)
}

func CloseDoorByName(room *structure.Room, doorName string) {
	door := findDoorByName(room, doorName)
	if door == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
		return
	} else if door.Closed {
		fmt.Println(constants.AlreadyClosedDoor, doorName)
		return
	}

	door.Closed = true
}

func getNextRoomInfo(room *structure.Room, direction constants.Direction) string {
	if room.Doors[direction] != nil {
		doorType := room.Doors[direction].DoorType
		isClosed := room.Doors[direction].Closed
		return fmt.Sprintf("%s(%s)", constants.DoorTypeStringMap[doorType], constants.DoorCloseStateStringMap[isClosed])
	}

	if room.Directions[direction] == nil {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.EmptyRoom])
	}
}

func printLine() {
	fmt.Println("==================================================")
}

func PrintUserNameAndStatus(status *structure.Status) {
	printLine()
	fmt.Printf("이름 : %s, Health: %d, Attack: %d, Defense: %d\n", status.Name, status.Attribute.Health, status.Attribute.Attack, status.Attribute.Defense)
}

func printObjectsInDirections(room *structure.Room) {
	printLine()
	for _, d := range constants.DirectionList {
		fmt.Printf("%s(%s) - %s\n", constants.DirStringMap[d], constants.DirStringEngMap[d], getNextRoomInfo(room, d))
	}
}

func printInventory(inventory *structure.Inventory) {
	printLine()
	fmt.Printf("아이템 정보 >>> ")
	for itemType, itemNum := range *inventory {
		if itemNum == 0 {
			continue
		}
		fmt.Printf("%s(%d) ", constants.ItemTypeStringMap[itemType], itemNum)
	}
	fmt.Println()
}

func printMovableDirections(room *structure.Room) {
	printLine()
	fmt.Printf("이동 가능한 방향 >>> ")
	for _, d := range constants.DirectionList {
		if !canMove(room, d) {
			continue
		}
		fmt.Printf("%s(%s) ", constants.DirStringMap[d], constants.DirStringEngMap[d])
	}
	fmt.Println()
}

func ShowRoomAndInventoryInfo(status *structure.Status) {
	room := status.Room

	PrintUserNameAndStatus(status)
	printObjectsInDirections(status.Room)
	printInventory(status.Inventory)
	printMovableDirections(status.Room)

	if isMonsterExistInRoom(room) {
		printLine()
		PrintMonsterInRoom(room)
	}

	printLine()
}

func PickUpItems(room *structure.Room, inventory *structure.Inventory) {
	for itemType, itemNum := range room.Items {
		if itemNum == 0 {
			continue
		}
		fmt.Printf("%s을 (%d)개 주웠습니다.\n", constants.ItemTypeStringMap[itemType], itemNum)
		room.Items[itemType] -= itemNum
		addItemToInventory(inventory, itemType)
	}
}
