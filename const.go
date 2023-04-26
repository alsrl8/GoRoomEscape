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

// Type of boundary
const (
	Wall = iota
	Clear
	WoodDoor
	GlassDoor
	LockedDoor
)

var WallStringMap = map[int]string{
	Wall:       "벽",
	Clear:      "방",
	WoodDoor:   "나무문",
	GlassDoor:  "유리문",
	LockedDoor: "잠긴문",
}

var StringWallMap = map[string]int{
	"벽":   Wall,
	"방":   Clear,
	"나무문": WoodDoor,
	"유리문": GlassDoor,
	"잠긴문": LockedDoor,
}

// State of Boundary
const (
	Open = iota
	Closed
)

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

// Common Messages
const SucceedOpeningDoor = "문을 열었습니다."

// Warnings
const WrongInput = "!!!잘못 입력하였습니다!!!"
const CanNotMoveWarning = "!!!그 방향으로 움직일 수 없습니다!!!"
const NoSuchItem = "!!!그런 아이템은 없습니다!!!"
const NoSuchDoor = "!!!그런 문은 없습니다!!!"
const CanNotOpenSuchDoor = "!!!열 수 없는 종류의 문입니다!!!"
const AlreadyOpenDoor = "!!!이미 문이 열려 있습니다!!!"
const NotEnoughItemsToOpenDoor = "!!!아이템이 모자라 문을 열 수 없습니다!!!"
