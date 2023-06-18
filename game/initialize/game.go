package initialize

import (
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
)

func InitGame() *structure.Status {
	rowLen := data.GetMapRowLen()
	colLen := data.GetMapColLen()
	area := initGameMap(rowLen, colLen)

	startPosition := data.GetMapStartPosition()
	var startLocation structure.Location = area[startPosition.Row][startPosition.Col]

	objectPositionAndType := data.GetMapObjectPositionAndType()
	putObjectOnArea(area, objectPositionAndType)

	npcPositionAndType := data.GetMapNpcPositionAndType()
	putNpcOnArea(area, npcPositionAndType)

	return initStatus(startLocation)
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
				Npc:          make(map[constants.NpcType]int),
			}
		}
		grid = append(grid, row)
	}
	connectAdjacentArea(grid)
	return grid
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

func putObjectOnArea(grid [][]*structure.Area, objectPositionAndType *[]structure.ObjectPositionAndType) {
	for _, object := range *objectPositionAndType {
		grid[object.Position.Row][object.Position.Col].Object[object.ObjectType] += 1
	}
}

func putNpcOnArea(grid [][]*structure.Area, npcPositionAndType *[]structure.NpcPositionAndType) {
	for _, npc := range *npcPositionAndType {
		grid[npc.Position.Row][npc.Position.Col].Npc[npc.NpcType] += 1
	}
}

func initStatus(startLocation structure.Location) *structure.Status {
	status := structure.Status{
		Location:  startLocation,
		Inventory: structure.Inventory{},
		Equipment: &structure.Equipment{},
		Attribute: data.GetAttribute(),
		BodyPartForArmor: &structure.BodyPartForArmor{
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
		BodyPartForWeapon: &structure.BodyPartForWeapon{
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
