package structure

import (
	"fmt"
	"goproject/constants"
)

type Area struct {
	Directions map[constants.Direction]*Location
}

func (area *Area) Move(direction constants.Direction) *Location {
	return (*area).Directions[direction]
}

func (area *Area) ShowInfo() {
	fmt.Println(constants.LineDivider)
	for _, dir := range constants.DirectionList {
		fmt.Printf(constants.DirectionInfoWithRoomInfo, constants.DirStringMap[dir], constants.DirStringEngMap[dir], area.getNearAreaInfo(dir))
	}
}

func (area *Area) getNearAreaInfo(direction constants.Direction) string {
	if area.Directions[direction] == nil {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.EmptyRoom])
	}
}
