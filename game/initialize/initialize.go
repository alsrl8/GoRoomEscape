package initialize

import (
	"goproject/constants"
	"goproject/structure"
)

func initGrid(rowLen int, colLen int) *[][]*structure.Room {
	var grid [][]*structure.Room
	for i := 0; i < 6; i++ {
		grid = append(grid, make([]*structure.Room, colLen))
	}
	return &grid
}

func newRoom(isGoal bool) *structure.Room {
	return &structure.Room{
		Directions: make(map[constants.Direction]*structure.Room),
		Doors:      make(map[constants.Direction]*structure.Door),
		IsGoal:     isGoal,
	}
}

func getNextRoom(grid *[][]*structure.Room, position structure.Position, direction constants.Direction) *structure.Room {
	room := (*grid)[position.Row][position.Col]
	return room.Directions[direction]
}

func getCounterDirection(direction constants.Direction) constants.Direction {
	return constants.DirectionList[(direction+2)%4]
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
	counterDirection := getCounterDirection(direction)
	toRoom.Directions[counterDirection] = fromRoom
}

func connectAdjacentRooms(grid *[][]*structure.Room) {
	dRow, dCol := [4]int{0, -1, 0, 1}, [4]int{1, 0, -1, 0}
	rowLen, colLen := len(*grid), len((*grid)[0])

	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if (*grid)[row][col] == nil {
				continue
			}

			for _, d := range constants.DirectionList {
				nr, nc := row+dRow[d], col+dCol[d]
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
		room.Doors[door.Direction] = &structure.Door{IsClosed: true, DoorType: door.DoorType}

		oppositeRoom := getNextRoom(grid, door.RoomPosition, door.Direction)
		if oppositeRoom == nil {
			return
		}
		counterDirection := getCounterDirection(door.Direction)
		oppositeRoom.Doors[counterDirection] = room.Doors[door.Direction]
	}
}

func addEndPoint(grid *[][]*structure.Room, endPosition structure.Position, direction constants.Direction) {
	room := (*grid)[endPosition.Row][endPosition.Col]
	room.Directions[direction] = newRoom(true)
}

func InitGameAndReturnStartPoint(rowLen int, colLen int, roomPositions *[]structure.Position, doorPositionAndType *[]structure.DoorPositionAndType, startPosition structure.Position, endPosition structure.Position, endDirection constants.Direction) *structure.Room {

	var grid = initGrid(rowLen, colLen)
	createEmptyRooms(grid, roomPositions)
	connectAdjacentRooms(grid)
	buildDoorsBetweenRooms(grid, doorPositionAndType)
	addEndPoint(grid, endPosition, endDirection)

	return (*grid)[startPosition.Row][startPosition.Col]
}
