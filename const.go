package main

// Direction
const (
	East = iota
	North
	West
	South
)

var DirStringMap = map[int]string{
	East:  "동쪽",
	North: "북쪽",
	West:  "서쪽",
	South: "남쪽",
}

// State of wall
const (
	Wall = iota
	Clear
	WoodDoorClosed
	WoodDoorOpen
	GlassDoorClosed
	GlassDoorOpen
	LockedDoorClosed
	LockedDoorOpen
)

// State of room
const (
	Empty = iota
	Hammer
	Key
)

const CanNotMoveWarning = "그 방향으로 움직일 수 없습니다."
