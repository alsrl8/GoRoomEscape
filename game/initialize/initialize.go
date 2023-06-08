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

func initGameMap(rowLen, colLen int) [][]*structure.Area {
	var grid [][]*structure.Area
	for r := 0; r < rowLen; r++ {
		row := make([]*structure.Area, colLen)
		for c := 0; c < colLen; c++ {
			row[c] = &structure.Area{
				LocationType: constants.EmptyArea,
				Directions:   make(map[constants.Direction]structure.Location),
				Object:       make(map[constants.ObjectType]int),
			}
		}
		grid = append(grid, row)
	}
	connectAdjacentArea(grid)
	return grid
}

func generateRoom(goalFlag bool) *structure.Room {
	return &structure.Room{
		Directions: make(map[constants.Direction]structure.Location),
		Doors:      make(map[constants.Direction]*structure.Door),
		GoalFlag:   goalFlag,
		Items:      make(map[constants.ItemType]int),
	}
}

func getNextRoom(grid [][]*structure.Room, position structure.Position, direction constants.Direction) *structure.Room {
	room := grid[position.Row][position.Col]
	return room.Directions[direction].(*structure.Room)
}

func createEmptyRooms(grid [][]*structure.Room, roomPositions *[]structure.Position) {
	for _, pos := range *roomPositions {
		grid[pos.Row][pos.Col] = generateRoom(false)
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

func connectAdjacentArea(grid [][]*structure.Area) {
	rowLen, colLen := len(grid), len(grid[0])

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			location := grid[r][c]
			for _, d := range constants.DirectionList {
				nr, nc := r+constants.DRow[d], c+constants.DCol[d]
				if nr < 0 || rowLen <= nr || nc < 0 || colLen <= nc {
					continue
				}
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

func addDungeonExit(dungeon [][]*structure.Room, endPosition structure.Position, direction constants.Direction) {
	room := dungeon[endPosition.Row][endPosition.Col]
	var goalRoom structure.Location = generateRoom(true)
	room.Directions[direction] = goalRoom
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

func InitDungeon(status *structure.Status, stageNum int) {
	rowLen := data.GetDungeonRowLen(stageNum)
	colLen := data.GetDungeonColLen(stageNum)
	roomPositions := data.GetDungeonRoomPositions(stageNum)
	dungeon := initDungeonMap(rowLen, colLen)
	createEmptyRooms(dungeon, roomPositions)
	connectAdjacentRooms(dungeon)

	startPosition := data.GetDungeonStartPosition(stageNum)
	endPosition := data.GetDungeonExitPosition(stageNum)
	endDirection := data.GetDungeonExitDirection(stageNum)
	addDungeonExit(dungeon, endPosition, endDirection)

	doorPositionAndType := data.GetDungeonDoorPositionAndType(stageNum)
	buildDoorsBetweenRooms(dungeon, doorPositionAndType)

	monsterWithPosition := data.GetMonsterWithPositionData(stageNum)
	putMonstersOnRooms(dungeon, &monsterWithPosition)

	itemPositionAndType := data.GetItemPositionAndType(stageNum)
	putItemsOnRooms(dungeon, itemPositionAndType)

	var startLocation structure.Location = dungeon[startPosition.Row][startPosition.Col]
	status.Location = startLocation
}

func InitGame() *structure.Status {
	rowLen := 2
	colLen := 2
	area := initGameMap(rowLen, colLen)
	area[0][0].Object[constants.DungeonEntrance] += 1
	var startLocation structure.Location = area[0][0]
	return initStatus(startLocation)
}

func initStatus(startLocation structure.Location) *structure.Status {
	status := structure.Status{
		Location:  startLocation,
		Inventory: structure.Inventory{},
		Equipment: &structure.Equipment{},
		Attribute: data.GetAttribute(),
		BodyPartForArmor: structure.BodyPartForArmor{
			Top: &structure.Armor{
				Item:          structure.Item{ItemType: constants.Nothing},
				WearableItems: []constants.ItemType{constants.LeatherCloth},
			},
			Pants: &structure.Armor{
				Item:          structure.Item{ItemType: constants.Nothing},
				WearableItems: []constants.ItemType{constants.LeatherPants},
			},
			Shoes: &structure.Armor{
				Item:          structure.Item{ItemType: constants.Nothing},
				WearableItems: []constants.ItemType{constants.LeatherShoes},
			},
		},
		BodyPartForWeapon: structure.BodyPartForWeapon{
			LeftHand: &structure.Weapon{
				Item:          structure.Item{ItemType: constants.Nothing},
				WearableItems: []constants.ItemType{constants.WoodSword, constants.IronSword},
			},
			RightHand: &structure.Weapon{
				Item:          structure.Item{ItemType: constants.Nothing},
				WearableItems: []constants.ItemType{constants.WoodSword, constants.IronSword},
			},
		},
	}
	return &status
}
