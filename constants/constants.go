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
var DRow = [4]int{0, -1, 0, 1}
var DCol = [4]int{1, 0, -1, 0}

const (
	NoDirection Direction = iota - 1
	East
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

var StringDoorTypeMap = map[string]DoorType{
	"나무문": WoodDoor,
	"유리문": GlassDoor,
	"잠긴문": LockedDoor,
}

var DoorCloseStateStringMap = map[bool]string{
	true:  "닫힘",
	false: "열림",
}

// Item Type
type ItemType int

const (
	Nothing ItemType = iota
	Hammer
	Key
	Box
	WoodSword
	IronSword
	LeatherCloth
	LeatherPants
	LeatherHat
	HealPotion
)

var ItemTypeStringMap = map[ItemType]string{
	Nothing:      "없음",
	Hammer:       "망치",
	Key:          "열쇠",
	Box:          "상자",
	WoodSword:    "목검",
	IronSword:    "철검",
	LeatherCloth: "가죽옷",
	LeatherPants: "가죽바지",
	LeatherHat:   "가죽모자",
	HealPotion:   "회복약",
}

var StringItemTypeMap = map[string]ItemType{
	"없음":   Nothing,
	"망치":   Hammer,
	"열쇠":   Key,
	"상자":   Box,
	"목검":   WoodSword,
	"철검":   IronSword,
	"가죽옷":  LeatherCloth,
	"가죽바지": LeatherPants,
	"가죽모자": LeatherHat,
	"회복약":  HealPotion,
}

// // Common Messages
const SucceedOpeningDoor = "문을 열었습니다."
const SucceedClosingDoor = "문을 닫았습니다."
const SucceedBreakingGlassDoor = "유리문을 망치로 부쉈습니다."
const SucceedUnlockLockedDoor = "잠긴문을 열쇠로 열었습니다."

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
