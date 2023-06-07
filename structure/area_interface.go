package structure

import "goproject/constants"

type Location interface {
	Move(direction constants.Direction) *Location
}

func (area *Area) Move(direction constants.Direction) *Location {
	return (*area).Directions[direction]
}

func (room *Room) Move(direction constants.Direction) *Location {
	return (*room).Directions[direction]
}
