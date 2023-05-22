package initialize

import (
	"goproject/constants"
	"goproject/game/command"
	"goproject/structure"
)

func initGrid(rowLen int, colLen int) *[][]*structure.Room {
	var grid [][]*structure.Room
	for i := 0; i < rowLen; i++ {
		grid = append(grid, make([]*structure.Room, colLen))
	}
	return &grid
}

func newRoom(goalFlag bool) *structure.Room {
	return &structure.Room{
		Directions: make(map[constants.Direction]*structure.Room),
		Doors:      make(map[constants.Direction]*structure.Door),
		GoalFlag:   goalFlag,
		Items:      make(map[constants.ItemType]int),
	}
}

func getNextRoom(grid *[][]*structure.Room, position structure.Position, direction constants.Direction) *structure.Room {
	room := (*grid)[position.Row][position.Col]
	return room.Directions[direction]
}

func createEmptyRooms(grid *[][]*structure.Room, roomPositions *[]structure.Position) {
	for _, pos := range *roomPositions {
		(*grid)[pos.Row][pos.Col] = newRoom(false)
	}
}

func connectTwoRooms(fromRoom *structure.Room, toRoom *structure.Room, direction constants.Direction) {
	if fromRoom.Directions[direction] != nil {
		return
	}

	fromRoom.Directions[direction] = toRoom
	counterDirection := command.GetCounterDirection(direction)
	toRoom.Directions[counterDirection] = fromRoom
}

func connectAdjacentRooms(grid *[][]*structure.Room) {
	rowLen, colLen := len(*grid), len((*grid)[0])

	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if (*grid)[row][col] == nil {
				continue
			}

			for _, d := range constants.DirectionList {
				nr, nc := row+constants.DRow[d], col+constants.DCol[d]
				if nr < 0 || rowLen <= nr || nc < 0 || colLen <= nc {
					continue
				} else if (*grid)[nr][nc] == nil {
					continue
				}

				connectTwoRooms((*grid)[row][col], (*grid)[nr][nc], d)
			}
		}
	}
}

func buildDoorsBetweenRooms(grid *[][]*structure.Room, doorPositionAndType *[]structure.DoorPositionAndType) {
	for _, door := range *doorPositionAndType {
		room := (*grid)[door.RoomPosition.Row][door.RoomPosition.Col]
		room.Doors[door.Direction] = &structure.Door{Closed: true, DoorType: door.DoorType}

		oppositeRoom := getNextRoom(grid, door.RoomPosition, door.Direction)
		if oppositeRoom == nil {
			continue
		}
		counterDirection := command.GetCounterDirection(door.Direction)
		oppositeRoom.Doors[counterDirection] = room.Doors[door.Direction]
	}
}

func addEndPoint(grid *[][]*structure.Room, endPosition structure.Position, direction constants.Direction) {
	room := (*grid)[endPosition.Row][endPosition.Col]
	room.Directions[direction] = newRoom(true)
}

func putItemsOnRooms(grid *[][]*structure.Room, itemPositionAndType *[]structure.ItemPositionAndType) {
	for _, item := range *itemPositionAndType {
		(*grid)[item.RoomPosition.Row][item.RoomPosition.Col].Items[item.ItemType] += 1
	}
}

func putMonstersOnRooms(grid *[][]*structure.Room, monsterWithPosition *[]structure.MonsterWithPosition) {
	for _, monsterInfo := range *monsterWithPosition {
		(*grid)[monsterInfo.RoomPosition.Row][monsterInfo.RoomPosition.Col].Monster = &structure.Monster{
			MonsterType: monsterInfo.Monster.MonsterType,
			Attribute:   monsterInfo.Monster.Attribute,
			DropItem:    monsterInfo.Monster.DropItem,
		}
	}
}

func putBoxesOnRooms(grid *[][]*structure.Room, boxPositionAndDropItem *[]structure.BoxPositionAndDropItem) {
	for _, boxInfo := range *boxPositionAndDropItem {
		(*grid)[boxInfo.RoomPosition.Row][boxInfo.RoomPosition.Col].DropItem = boxInfo.DropItem
	}
}

func InitGameAndReturnStatus(
	rowLen int,
	colLen int,
	roomPositions *[]structure.Position,
	doorPositionAndType *[]structure.DoorPositionAndType,
	startPosition structure.Position,
	endPosition structure.Position,
	endDirection constants.Direction,
	itemPositionAndType *[]structure.ItemPositionAndType,
	monsterWithPosition *[]structure.MonsterWithPosition,
	boxPositionAndDropItem *[]structure.BoxPositionAndDropItem,
) *structure.Status {
	var grid = initGrid(rowLen, colLen)
	createEmptyRooms(grid, roomPositions)
	connectAdjacentRooms(grid)
	buildDoorsBetweenRooms(grid, doorPositionAndType)
	addEndPoint(grid, endPosition, endDirection)
	putItemsOnRooms(grid, itemPositionAndType)
	putMonstersOnRooms(grid, monsterWithPosition)
	putBoxesOnRooms(grid, boxPositionAndDropItem)
	status := initStatus((*grid)[startPosition.Row][startPosition.Col])
	return status
}

func initStatus(startRoom *structure.Room) *structure.Status {
	status := structure.Status{
		Room:      startRoom,
		Inventory: &structure.Inventory{},
		Equipment: &structure.Equipment{},
		Attribute: structure.Attribute{
			Health:  50,
			Attack:  3,
			Defense: 3,
		},
	}
	return &status
}
