package command

import "goproject/structure"

func GetCurrentRoom(status *structure.Status) *structure.Room {
	room, ok := status.Location.(*structure.Room)
	if !ok {
		panic("Current Location Is Not A Room")
	}
	return room
}
