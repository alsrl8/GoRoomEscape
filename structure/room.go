package structure

import (
	"fmt"
	"goproject/constants"
)

type Room struct {
	constants.LocationType
	Doors      map[constants.Direction]*Door
	Directions map[constants.Direction]Location
	Items      map[constants.ItemType]int
	Monster    *Monster
	NpcMap     map[constants.NpcType]int
	Object     map[constants.ObjectType]int
}

type Door struct {
	Closed   bool
	DoorType constants.DoorType
}

type DoorPositionAndType struct {
	RoomPosition Position
	Direction    constants.Direction
	DoorType     constants.DoorType
}

func (room *Room) GetLocationType() constants.LocationType {
	return room.LocationType
}

func (room *Room) Move(direction constants.Direction) Location {
	return (*room).Directions[direction]
}

func (room *Room) CanMove(direction constants.Direction) bool {
	if room.Directions[direction] == nil {
		return false
	} else if room.Doors[direction] != nil && room.Doors[direction].Closed {
		return false
	}
	return true
}

func (room *Room) Connect(near Location, direction constants.Direction) {
	room.Directions[direction] = near
}

func (room *Room) ShowInfo() {
	room.showObjectInRoom()
	room.showMonsterInRoom()
	room.showItemInRoom()
	room.showNearRoomInfo()
	room.showMovableRoom()
}

func (room *Room) GetObjectMap() map[constants.ObjectType]int {
	return room.Object
}

func (room *Room) showMonsterInRoom() {
	if room.Monster == nil {
		return
	}

	fmt.Println(constants.LineDivider)
	monster := room.Monster
	fmt.Printf(constants.MonsterStatus, constants.MonsterTypeStringMap[monster.MonsterType], monster.Attribute.Health, monster.Attribute.Attack, monster.Attribute.Defense)
}

func (room *Room) showObjectInRoom() {
	fmt.Println(constants.LineDivider)
	for objectType, objectNum := range room.Object {
		if objectNum <= 0 {
			continue
		}
		fmt.Printf(constants.ObjectTypeAndNum, constants.ObjectTypeStringMap[objectType], objectNum)
	}
	fmt.Println()
}

func (room *Room) showItemInRoom() {
	fmt.Println(constants.LineDivider)
	for itemType, itemNum := range room.Items {
		if itemNum <= 0 {
			continue
		}
		fmt.Printf(constants.ItemTypeAndNum, constants.ItemTypeStringMap[itemType], itemNum)
	}
	fmt.Println()
}

func (room *Room) showNearRoomInfo() {
	fmt.Println(constants.LineDivider)
	for _, dir := range constants.DirectionList {
		fmt.Printf(constants.DirectionInfoWithRoomInfo, constants.DirStringMap[dir], constants.DirStringEngMap[dir], room.getNearRoomInfo(dir))
	}
}

func (room *Room) getNearRoomInfo(direction constants.Direction) string {
	if room.Doors[direction] != nil {
		doorType := room.Doors[direction].DoorType
		isClosed := room.Doors[direction].Closed
		return fmt.Sprintf(constants.DirectionInfo, constants.DoorTypeStringMap[doorType], constants.DoorCloseStateStringMap[isClosed])
	}

	if room.Directions[direction] == nil {
		return fmt.Sprintf(constants.LocationTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.LocationTypeStringMap[room.GetLocationType()])
	}
}

func (room *Room) showMovableRoom() {
	fmt.Println(constants.LineDivider)
	fmt.Printf(constants.MovableDirectionTitle)
	for _, d := range constants.DirectionList {
		if !room.CanMove(d) {
			continue
		}
		fmt.Printf(constants.DirectionInfo, constants.DirStringMap[d], constants.DirStringEngMap[d])
	}
	fmt.Println()
}
