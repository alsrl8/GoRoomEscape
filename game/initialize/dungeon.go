package initialize

import (
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
)

func initDungeonMap(rowLen int, colLen int) *structure.Dungeon {
	var grid [][]*structure.Room
	for i := 0; i < rowLen; i++ {
		grid = append(grid, make([]*structure.Room, colLen))
	}
	return &structure.Dungeon{
		Grid: grid,
	}
}

func GenerateDungeon(status *structure.Status, stageNum constants.StageNum) *structure.Room {
	rowLen := data.GetDungeonRowLen(stageNum)
	colLen := data.GetDungeonColLen(stageNum)
	roomPosition := data.GetDungeonRoomPosition(stageNum)
	dungeon := initDungeonMap(rowLen, colLen)
	dungeon.CreateEmptyRoom(roomPosition)
	dungeon.ConnectAdjacentRoom()

	exitPosition := data.GetDungeonExitPosition(stageNum)
	dungeon.AddExit(status.Location, exitPosition)

	doorPositionAndType := data.GetDungeonDoorPositionAndType(stageNum)
	dungeon.BuildDoor(doorPositionAndType)

	monsterWithPosition := data.GetMonsterWithPositionData(stageNum)
	dungeon.PutMonster(&monsterWithPosition)

	itemPositionAndType := data.GetItemPositionAndType(stageNum)
	dungeon.PutItem(itemPositionAndType)

	startPosition := data.GetDungeonStartPosition(stageNum)
	return dungeon.Grid[startPosition.Row][startPosition.Col]
}
