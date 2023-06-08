package structure

import (
	"fmt"
	"goproject/constants"
)

type Area struct {
	constants.LocationType
	Directions map[constants.Direction]Location
	Object     map[constants.ObjectType]int
}

func (area *Area) GetLocationType() constants.LocationType {
	return area.LocationType
}

func (area *Area) Move(direction constants.Direction) Location {
	return (*area).Directions[direction]
}

func (area *Area) CanMove(direction constants.Direction) bool {
	if area.Directions[direction] == nil {
		return false
	}
	return true
}

func (area *Area) Connect(near Location, direction constants.Direction) {
	area.Directions[direction] = near
}

func (area *Area) ShowInfo() {
	area.showObjectInArea()
	area.showNearAreaInfo()
	area.showMovableArea()
}

func (area *Area) showObjectInArea() {
	fmt.Println(constants.LineDivider)
	for objectType, objectNum := range area.Object {
		if objectNum <= 0 {
			continue
		}
		fmt.Printf(constants.ObjectTypeAndNum, constants.ObjectTypeStringMap[objectType], objectNum)
	}
	fmt.Println()
}

func (area *Area) showNearAreaInfo() {
	fmt.Println(constants.LineDivider)
	for _, dir := range constants.DirectionList {
		fmt.Printf(constants.DirectionInfoWithRoomInfo, constants.DirStringMap[dir], constants.DirStringEngMap[dir], area.getNearAreaInfo(dir))
	}
}

func (area *Area) getNearAreaInfo(direction constants.Direction) string {
	if area.Directions[direction] == nil {
		return fmt.Sprintf(constants.LocationTypeStringMap[constants.VoidArea])
	} else {
		return fmt.Sprintf(constants.LocationTypeStringMap[area.GetLocationType()])
	}
}

func (area *Area) showMovableArea() {
	fmt.Println(constants.LineDivider)
	fmt.Printf(constants.MovableDirectionTitle)
	for _, d := range constants.DirectionList {
		if !area.CanMove(d) {
			continue
		}
		fmt.Printf(constants.DirectionInfo, constants.DirStringMap[d], constants.DirStringEngMap[d])
	}
	fmt.Println()
}
