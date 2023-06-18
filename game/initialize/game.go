package initialize

import (
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
)

func InitGame() *structure.Status {
	rowLen := data.GetMapRowLen()
	colLen := data.GetMapColLen()
	gameMap := initGameMap(rowLen, colLen)

	objectPositionAndType := data.GetMapObjectPositionAndType()
	gameMap.PutObject(objectPositionAndType)

	npcPositionAndType := data.GetMapNpcPositionAndType()
	gameMap.PutNpc(npcPositionAndType)

	startPosition := data.GetMapStartPosition()
	var startLocation structure.Location = gameMap.Grid[startPosition.Row][startPosition.Col]

	return initStatus(startLocation)
}

func initGameMap(rowLen, colLen int) *structure.GameMap {
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
	return &structure.GameMap{Grid: grid}
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
