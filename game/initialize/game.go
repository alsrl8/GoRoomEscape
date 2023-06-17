package initialize

import (
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
)

func InitGame() *structure.Status {
	rowLen := 2
	colLen := 2
	area := initGameMap(rowLen, colLen)
	area[0][0].Object[constants.DungeonEntrance] += 1
	area[0][0].Npc[constants.Merchant] += 1
	area[0][1].Npc[constants.GodOfDeath] += 1
	area[1][0].Npc[constants.Blacksmith] += 1
	var startLocation structure.Location = area[0][0]
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
