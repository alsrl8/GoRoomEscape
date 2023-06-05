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
	WoodShield
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

type MonsterType int

const (
	Squirrel MonsterType = iota
	Rabbit
	Deer
)

var MonsterTypeStringMap = map[MonsterType]string{
	Squirrel: "다람쥐",
	Rabbit:   "토끼",
	Deer:     "사슴",
}

var StringMonsterTypeMap = map[string]MonsterType{
	"다람쥐": Squirrel,
	"토끼":  Rabbit,
	"사슴":  Deer,
}

const LineDivider = "=================================================="

const SucceedOpeningDoor = "문을 열었습니다."
const SucceedClosingDoor = "문을 닫았습니다."
const SucceedBreakingGlassDoor = "유리문을 망치로 부쉈습니다."
const SucceedUnlockLockedDoor = "잠긴문을 열쇠로 열었습니다."

const GetItem = "아이템을 얻었습니다 >> %s: %d개\n"
const NoSpecificItemOnBody = "%s 아이템을 몸에 장비하고 있지 않습니다.\n"
const UseItem = "아이템(%s)을 사용했습니다.\n"
const EquipEquipment = "%s에 %s 장비\n"
const DisarmEquipment = "%s에서 장비(%s)를 해제했습니다.\n"

const PlayerStatus = "이름 : %s, Health: %d, Attack: %d, Defense: %d\n"
const MonsterStatus = "%s >> Health: %d, Attack: %d, Defense: %d\n"
const DirectionInfoWithRoomInfo = "%s(%s) - %s\n"
const DirectionInfo = "%s(%s) "
const MovableDirectionTitle = "이동 가능한 방향 >>> "
const ItemInfoTitle = "아이템 정보 >>> "
const ItemTypeAndNum = "%s(%d)"

const DuringBattle = "<<< 몬스터(%s)와 전투 중입니다 >>>\n"
const AttackMonster = "몬스터(%s)를 공격했습니다.\n"
const DamageMonster = "%s에게 %d 만큼 피해를 입혔습니다.\n"
const GetAttackedByMonster = "몬스터(%s)가 공격했습니다.\n"
const DamageByMonster = "몬스터(%s)로부터 %d 만큼 피해를 입었습니다.\n"
const CanNotAttack = "공격할 수 없습니다."
const CanNotGuard = "방어할 수 없습니다."
const OnGuard = "방어 중입니다."
const RunningSucceeded = "도망쳤습니다."
const RunningFailed = "도망치려 했으나 실패했습니다."
const RaiseShield = "방패를 들었습니다."

const KillMonster = "당신은 %s를 죽였습니다. ㅠㅠ\n"
const GetKilled = "당신은 죽었습니다. ㅠㅠ\n"

const InvalidUserName = "!!!사용자 이름이 잘못 입력되었습니다"
const WrongInput = "!!!잘못 입력하였습니다"
const CanNotMoveWarning = "!!!그 방향으로 움직일 수 없습니다"
const NoSuchItem = "!!!그런 아이템은 없습니다"
const NoSuchDoor = "!!!그런 문은 없습니다"
const NotEnoughItem = "!!!아이템 수가 모자랍니다"
const CanNotOpenSuchDoor = "!!!열 수 없는 종류의 문입니다"
const CanNotCloseSuchDoor = "!!!닫을 수 없는 종류의 문입니다"
const AlreadyOpenDoor = "!!!이미 문이 열려 있습니다"
const AlreadyClosedDoor = "!!!이미 문이 닫혀 있습니다"
const NotEnoughItemsToOpenDoor = "!!!아이템이 모자라 문을 열 수 없습니다"
const CanNotWear = "!!!장비할 수 없는 아이템입니다"
const NoBodyPartToWear = "!!!몸에 장비할 빈 공간이 없습니다"
const NoItemInInventory = "!!!해당 아이템이 가방에 없습니다"
const NoEquipmentOnBodyPart = "!!!해당 부위에 장비하고 있는 아이템이 없습니다"
const NoSuchEquipmentOnBody = "!!!장비하고 있는 아이템이 아닙니다"
const MonsterExistInTheRoom = "!!!방 안에 몬스터가 있습니다"
const NoSuchMonster = "!!!그런 몬스터는 없습니다"
const CanNotUseSuchItem = "!!!사용할 수 없는 종류의 아이템입니다"
const NoSpecificTargetForItem = "!!!아이템을 사용할 타겟이 없습니다"
const NoMatchItemAndDoor = "!!!아이템과 문 종류가 적합하지 않습니다"

const QuitGame = "종료했습니다."
const GameOver = "실패했습니다."
const GameClear = "탈출했습니다."

const PanicEquipmentListLengthAndBodyPartListLength = "Length of equipment list and body part list must be same"
