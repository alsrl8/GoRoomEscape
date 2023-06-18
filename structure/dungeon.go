package structure

import "goproject/constants"

type Dungeon struct {
	Grid [][]*Room
}

func (dungeon *Dungeon) CreateEmptyRoom(roomPosition *[]Position) {
	grid := dungeon.Grid
	for _, pos := range *roomPosition {
		grid[pos.Row][pos.Col] = &Room{
			Directions: make(map[constants.Direction]Location),
			Doors:      make(map[constants.Direction]*Door),
			Items:      make(map[constants.ItemType]int),
			Object:     make(map[constants.ObjectType]int),
		}
	}
}

func (dungeon *Dungeon) ConnectAdjacentRoom() {
	grid := dungeon.Grid
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

func (dungeon *Dungeon) AddExit(currentLocation Location, exitPosition Position) {
	room := dungeon.Grid[exitPosition.Row][exitPosition.Col]
	room.Object[constants.DungeonExit] += 1
	room.Directions[constants.Exit] = currentLocation
}

func (dungeon *Dungeon) GetNextRoom(position Position, direction constants.Direction) *Room {
	grid := dungeon.Grid
	room := grid[position.Row][position.Col]
	if room.Directions[direction] == nil {
		return nil
	}
	return room.Directions[direction].(*Room)
}

func (dungeon *Dungeon) BuildDoor(doorPositionAndType *[]DoorPositionAndType) {
	grid := dungeon.Grid
	for _, door := range *doorPositionAndType {
		room := grid[door.Position.Row][door.Position.Col]
		room.Doors[door.Direction] = &Door{Closed: true, DoorType: door.DoorType}

		oppositeRoom := dungeon.GetNextRoom(door.Position, door.Direction)
		if oppositeRoom == nil {
			continue
		}
		counterDirection := constants.GetCounterDirection(door.Direction)
		oppositeRoom.Doors[counterDirection] = room.Doors[door.Direction]
	}
}

func (dungeon *Dungeon) PutMonster(monsterWithPosition *[]MonsterWithPosition) {
	grid := dungeon.Grid
	for _, monsterInfo := range *monsterWithPosition {
		grid[monsterInfo.Position.Row][monsterInfo.Position.Col].Monster = &Monster{
			MonsterType: monsterInfo.Monster.MonsterType,
			Attribute:   monsterInfo.Monster.Attribute,
			DropItems:   monsterInfo.Monster.DropItems,
		}
	}
}

func (dungeon *Dungeon) PutItem(itemPositionAndType *[]ItemPositionAndType) {
	grid := dungeon.Grid
	for _, item := range *itemPositionAndType {
		grid[item.Position.Row][item.Position.Col].Items[item.ItemType] += 1
	}
}
