package data

import (
	"goproject/constants"
	"goproject/structure"
)

func GetMapRowLen() int {
	return 2
}

func GetMapColLen() int {
	return 2
}

func GetMapStartPosition() structure.Position {
	return structure.Position{Row: 0, Col: 0}
}

func GetMapObjectPositionAndType() *[]structure.ObjectPositionAndType {
	return &[]structure.ObjectPositionAndType{
		{Position: structure.Position{Row: GetMapStartPosition().Row, Col: GetMapStartPosition().Col}, Object: structure.Object{ObjectType: constants.DungeonEntrance}},
	}
}

func GetMapNpcPositionAndType() *[]structure.NpcPositionAndType {
	return &[]structure.NpcPositionAndType{
		{Position: structure.Position{Row: 0, Col: 0}, Npc: structure.Npc{NpcType: constants.Merchant}},
		{Position: structure.Position{Row: 0, Col: 1}, Npc: structure.Npc{NpcType: constants.GodOfDeath}},
		{Position: structure.Position{Row: 1, Col: 0}, Npc: structure.Npc{NpcType: constants.Blacksmith}},
	}
}
