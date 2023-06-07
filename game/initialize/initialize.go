package initialize

import (
	"goproject/constants"
	"goproject/game/command"
	"goproject/game/data"
	"goproject/structure"
)

func initGrid(rowLen int, colLen int) *[][]*structure.Room {
	var grid [][]*structure.Room
	for i := 0; i < rowLen; i++ {
		grid = append(grid, make([]*structure.Room, colLen))
	}
	return &grid
}

func generateRoom(goalFlag bool) *structure.Room {
	return &structure.Room{
		Directions: make(map[constants.Direction]*structure.Location),
		Doors:      make(map[constants.Direction]*structure.Door),
		GoalFlag:   goalFlag,
		Items:      make(map[constants.ItemType]int),
	}
}

func getNextRoom(grid *[][]*structure.Room, position structure.Position, direction constants.Direction) *structure.Room {
	room := (*grid)[position.Row][position.Col]
	return (*room.Directions[direction]).(*structure.Room)
}

func createEmptyRooms(grid *[][]*structure.Room, roomPositions *[]structure.Position) {
	for _, pos := range *roomPositions {
		(*grid)[pos.Row][pos.Col] = generateRoom(false)
	}
}

func connectTwoRooms(fromRoom *structure.Room, toRoom *structure.Room, direction constants.Direction) {
	if fromRoom.Directions[direction] != nil {
		return
	}

	var _toRoom structure.Location = toRoom
	fromRoom.Directions[direction] = &_toRoom
	counterDirection := command.GetCounterDirection(direction)
	var _fromRoom structure.Location = fromRoom
	toRoom.Directions[counterDirection] = &_fromRoom
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
	var goalRoom structure.Location = generateRoom(true)
	room.Directions[direction] = &goalRoom
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
			DropItems:   monsterInfo.Monster.DropItems,
		}
	}
}

func InitGameAndReturnStatus() *structure.Status {
	rowLen := data.GetRowLen()
	colLen := data.GetColLen()
	roomPositions := data.GetRoomPositions()
	grid := initGrid(rowLen, colLen)
	createEmptyRooms(grid, roomPositions)
	connectAdjacentRooms(grid)

	startPosition := data.GetStartPosition()
	endPosition := data.GetEndPosition()
	endDirection := data.GetEndDirection()
	addEndPoint(grid, endPosition, endDirection)

	doorPositionAndType := data.GetDoorPositionAndType()
	buildDoorsBetweenRooms(grid, doorPositionAndType)

	monsterWithPosition := data.GetMonsterWithPositionData()
	putMonstersOnRooms(grid, &monsterWithPosition)

	itemPositionAndType := data.GetItemPositionAndType()
	putItemsOnRooms(grid, itemPositionAndType)

	status := initStatus((*grid)[startPosition.Row][startPosition.Col])
	return status
}

func initStatus(startRoom *structure.Room) *structure.Status {
	var location structure.Location = startRoom
	status := structure.Status{
		Location:  &location,
		Inventory: &structure.Inventory{},
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
