package structure

import (
	"goproject/constants"
)

type Npc struct {
	constants.NpcType
}

type NpcPositionAndType struct {
	Npc
	Position
}
