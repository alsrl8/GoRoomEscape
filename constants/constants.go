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

// DoorType Door Type
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

// ItemType Item Type
type ItemType int

const (
	Nothing ItemType = iota
	Hammer
	Key
	Box
	WoodSword
	IronSword
	LeatherHat
	LeatherCloth
	LeatherPants
	LeatherShoes
	HealPotion
)

var ItemTypeStringMap = map[ItemType]string{
	Nothing:      "없음",
	Hammer:       "망치",
	Key:          "열쇠",
	Box:          "상자",
	WoodSword:    "목검",
	IronSword:    "철검",
	LeatherHat:   "가죽모자",
	LeatherCloth: "가죽옷",
	LeatherPants: "가죽바지",
	LeatherShoes: "가죽신발",
	HealPotion:   "회복약",
}

var StringItemTypeMap = map[string]ItemType{
	"없음":   Nothing,
	"망치":   Hammer,
	"열쇠":   Key,
	"상자":   Box,
	"목검":   WoodSword,
	"철검":   IronSword,
	"가죽모자": LeatherHat,
	"가죽옷":  LeatherCloth,
	"가죽바지": LeatherPants,
	"가죽신발": LeatherShoes,
	"회복약":  HealPotion,
}

var ItemTypeWearableMap = map[ItemType]bool{
	WoodSword:    true,
	IronSword:    true,
	LeatherCloth: true,
	LeatherPants: true,
	LeatherShoes: true,
}

type BodyPart int

var BodyPartList = []BodyPart{Top, Pants, Shoes, LeftHand, RightHand}

const (
	Top BodyPart = iota
	Pants
	Shoes
	LeftHand
	RightHand
)

var ItemTypeBodyPartMap = map[ItemType][]BodyPart{
	WoodSword:    {LeftHand, RightHand},
	IronSword:    {LeftHand, RightHand},
	LeatherCloth: {Top},
	LeatherPants: {Pants},
	LeatherShoes: {Shoes},
}

var BodyPartStringMap = map[BodyPart]string{
	Top:       "상의",
	Pants:     "하의",
	Shoes:     "신발",
	LeftHand:  "왼손",
	RightHand: "오른손",
}

// Common Messages
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
const CanNotWear = "!!!장비할 수 없는 아이템입니다!!!"
const NoBodyPartToWear = "!!!몸에 장비할 빈 공간이 없습니다!!!"
const NoItemInInventory = "!!!해당 아이템이 가방에 없습니다!!!"
