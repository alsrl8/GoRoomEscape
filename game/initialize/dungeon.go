package initialize

import (
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
)

func initDungeonMap(rowLen int, colLen int) [][]*structure.Room {
	var grid [][]*structure.Room
	for i := 0; i < rowLen; i++ {
		grid = append(grid, make([]*structure.Room, colLen))
	}
	return grid
}

func generateRoom() *structure.Room {
	return &structure.Room{
		Directions: make(map[constants.Direction]structure.Location),
		Doors:      make(map[constants.Direction]*structure.Door),
		Items:      make(map[constants.ItemType]int),
		Object:     make(map[constants.ObjectType]int),
	}
}

func getNextRoom(grid [][]*structure.Room, position structure.Position, direction constants.Direction) *structure.Room {
	room := grid[position.Row][position.Col]
	if room.Directions[direction] == nil {
		return nil
	}
	return room.Directions[direction].(*structure.Room)
}

func createEmptyRooms(grid [][]*structure.Room, roomPositions *[]structure.Position) {
	for _, pos := range *roomPositions {
		grid[pos.Row][pos.Col] = generateRoom()
	}
}

func connectAdjacentRooms(grid [][]*structure.Room) {
	rowLen, colLen := len(grid), len(grid[0])

	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == nil {
				continue
			}

			for _, d := range constants.DirectionList {
				nr, nc := row+constants.DRow[d], col+constants.DCol[d]
				if nr < 0 || rowLen <= nr || nc < 0 || colLen <= nc {
					continue
				} else if grid[nr][nc] == nil {
					continue
				}

				location := grid[row][col]
				near := grid[nr][nc]
				location.Connect(near, d)
			}
		}
	}
}

func buildDoorsBetweenRooms(grid [][]*structure.Room, doorPositionAndType *[]structure.DoorPositionAndType) {
	for _, door := range *doorPositionAndType {
		room := grid[door.RoomPosition.Row][door.RoomPosition.Col]
		room.Doors[door.Direction] = &structure.Door{Closed: true, DoorType: door.DoorType}

		oppositeRoom := getNextRoom(grid, door.RoomPosition, door.Direction)
		if oppositeRoom == nil {
			continue
		}
		counterDirection := constants.GetCounterDirection(door.Direction)
		oppositeRoom.Doors[counterDirection] = room.Doors[door.Direction]
	}
}

func addDungeonExit(currentLocation structure.Location, dungeon [][]*structure.Room, endPosition structure.Position) {
	room := dungeon[endPosition.Row][endPosition.Col]
	room.Object[constants.DungeonExit] += 1
	room.Directions[constants.Exit] = currentLocation
}

func putItemsOnRooms(grid [][]*structure.Room, itemPositionAndType *[]structure.ItemPositionAndType) {
	for _, item := range *itemPositionAndType {
		grid[item.RoomPosition.Row][item.RoomPosition.Col].Items[item.ItemType] += 1
	}
}

func putMonstersOnRooms(grid [][]*structure.Room, monsterWithPosition *[]structure.MonsterWithPosition) {
	for _, monsterInfo := range *monsterWithPosition {
		grid[monsterInfo.RoomPosition.Row][monsterInfo.RoomPosition.Col].Monster = &structure.Monster{
			MonsterType: monsterInfo.Monster.MonsterType,
			Attribute:   monsterInfo.Monster.Attribute,
			DropItems:   monsterInfo.Monster.DropItems,
		}
	}
}

func GenerateDungeon(status *structure.Status, stageNum constants.StageNum) *structure.Room {
	rowLen := data.GetDungeonRowLen(stageNum)
	colLen := data.GetDungeonColLen(stageNum)
	roomPositions := data.GetDungeonRoomPositions(stageNum)
	dungeon := initDungeonMap(rowLen, colLen)
	createEmptyRooms(dungeon, roomPositions)
	connectAdjacentRooms(dungeon)

	startPosition := data.GetDungeonStartPosition(stageNum)
	endPosition := data.GetDungeonExitPosition(stageNum)
	addDungeonExit(status.Location, dungeon, endPosition)

	doorPositionAndType := data.GetDungeonDoorPositionAndType(stageNum)
	buildDoorsBetweenRooms(dungeon, doorPositionAndType)

	monsterWithPosition := data.GetMonsterWithPositionData(stageNum)
	putMonstersOnRooms(dungeon, &monsterWithPosition)

	itemPositionAndType := data.GetItemPositionAndType(stageNum)
	putItemsOnRooms(dungeon, itemPositionAndType)

	return dungeon[startPosition.Row][startPosition.Col]
}
