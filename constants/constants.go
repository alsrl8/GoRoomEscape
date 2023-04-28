package constants

type SpaceType int

const (
	Wall SpaceType = iota
	EmptyRoom
)

var SpaceTypeStringMap = map[SpaceType]string{
	Wall:      "벽",
	EmptyRoom: "방",
}

// Direction
type Direction int

var DirectionList = []Direction{East, North, West, South}

const (
	East Direction = iota
	North
	West
	South
)

var DirStringMap = map[Direction]string{
	East:  "동쪽",
	North: "북쪽",
	West:  "서쪽",
	South: "남쪽",
}

var DirStringEngMap = map[Direction]string{
	East:  "E",
	North: "N",
	West:  "W",
	South: "S",
}

// Door Type
type DoorType int

const (
	WoodDoor DoorType = iota
	GlassDoor
	LockedDoor
)

var DoorTypeStringMap = map[DoorType]string{
	WoodDoor:   "나무문",
	GlassDoor:  "유리문",
	LockedDoor: "잠긴문",
}

var DoorCloseStateStringMap = map[bool]string{
	true:  "닫힘",
	false: "열림",
}

// Item Type
type ItemType int

const (
	Hammer ItemType = iota
	Key
)

var ItemTypeStringMap = map[ItemType]string{
	Hammer: "망치",
	Key:    "열쇠",
}

// // Common Messages
// const SucceedOpeningDoor = "문을 열었습니다."
// const SucceedClosingDoor = "문을 닫았습니다."

// Warnings
const WrongInput = "!!!잘못 입력하였습니다!!!"
const CanNotMoveWarning = "!!!그 방향으로 움직일 수 없습니다!!!"
const NoSuchItem = "!!!그런 아이템은 없습니다!!!"
const NoSuchDoor = "!!!그런 문은 없습니다!!!"
const CanNotOpenSuchDoor = "!!!열 수 없는 종류의 문입니다!!!"
const CanNotCloseSuchDoor = "!!!닫을 수 없는 종류의 문입니다!!!"
const AlreadyOpenDoor = "!!!이미 문이 열려 있습니다!!!"
const AlreadyClosedDoor = "!!!이미 문이 닫혀 있습니다!!!"
const NotEnoughItemsToOpenDoor = "!!!아이템이 모자라 문을 열 수 없습니다!!!"
