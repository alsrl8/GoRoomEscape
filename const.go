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

var DirStringEngMap = map[int]string{
	East:  "E",
	North: "N",
	West:  "W",
	South: "S",
}

// State of wall
const (
	Wall = iota
	Clear
	ClosedDoor
	OpenDoor
	GlassDoor
	LockedDoor
)

var WallStringMap = map[int]string{
	Wall:       "벽",
	Clear:      "방",
	ClosedDoor: "닫힌 문",
	OpenDoor:   "열린 문",
	GlassDoor:  "유리문",
	LockedDoor: "잠긴 문",
}

var StringWallMap = map[string]int{
	"벽":    Wall,
	"방":    Clear,
	"닫힌 문": ClosedDoor,
	"열린 문": OpenDoor,
	"유리문":  GlassDoor,
	"잠긴 문": LockedDoor,
}

// Item
const (
	Empty = iota
	Hammer
	Key
)

var ItemStringMap = map[int]string{
	Empty:  "없음",
	Hammer: "망치",
	Key:    "열쇠",
}

var ItemDescMap = map[int]string{
	Hammer: "유리문을 부수고 지나갈 수 있게 한다.",
	Key:    "잠긴 문을 열고 지나갈 수 있게 한다.",
}

var StringItemMap = map[string]int{
	"망치": Hammer,
	"열쇠": Key,

	"Hammer": Hammer,
	"Key":    Key,
}

const ItemTypeNum = 2

const WrongInput = "!!!잘못 입력하였습니다!!!"
const CanNotMoveWarning = "!!!그 방향으로 움직일 수 없습니다!!!"
const NoSuchItem = "!!!그런 아이템은 없습니다!!!"
const NoSuchDoor = "!!!그런 문은 없습니다!!!"
