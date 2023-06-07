package structure

import (
	"fmt"
	"goproject/constants"
)

type Room struct {
	Doors      map[constants.Direction]*Door
	Directions map[constants.Direction]*Location
	GoalFlag   bool
	Items      map[constants.ItemType]int
	Monster    *Monster
	NpcMap     map[constants.NpcType]int
	Place      constants.PlaceType
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

func (room *Room) Move(direction constants.Direction) *Location {
	return (*room).Directions[direction]
}

func (room *Room) ShowInfo() {
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
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.EmptyRoom])
	}
}
