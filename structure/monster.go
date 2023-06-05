package structure

import "goproject/constants"

type Monster struct {
	constants.MonsterType
	Attribute
	DropItems []DropItem
}

type MonsterWithPosition struct {
	RoomPosition Position
	Monster
}
