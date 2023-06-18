package structure

import "goproject/constants"

type GameMap struct {
	Grid [][]*Area
}

func (gameMap *GameMap) ConnectAdjacentArea() {
	grid := gameMap.Grid
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

func (gameMap *GameMap) PutObject(objectPositionAndType *[]ObjectPositionAndType) {
	grid := gameMap.Grid
	for _, object := range *objectPositionAndType {
		grid[object.Position.Row][object.Position.Col].Object[object.ObjectType] += 1
	}
}

func (gameMap *GameMap) PutNpc(npcPositionAndType *[]NpcPositionAndType) {
	grid := gameMap.Grid
	for _, npc := range *npcPositionAndType {
		grid[npc.Position.Row][npc.Position.Col].Npc[npc.NpcType] += 1
	}
}
