package structure

import "goproject/constants"

type Monster struct {
	MonsterType constants.MonsterType
	Attribute   Attribute
	DropItem    []DropItem
}

type MonsterWithPosition struct {
	RoomPosition Position
	Monster      Monster
}
