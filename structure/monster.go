package structure

import (
	"fmt"
	"goproject/constants"
)

type Monster struct {
	constants.MonsterType
	Attribute
	DropItems []DropItem
}

type MonsterWithPosition struct {
	RoomPosition Position
	Monster
}

func (monster *Monster) ShowInfo() {
	fmt.Println(constants.LineDivider)
	fmt.Printf(constants.MonsterStatus, constants.MonsterTypeStringMap[monster.MonsterType], monster.Attribute.Health, monster.Attribute.Attack, monster.Attribute.Defense)
}

func (monster *Monster) Attack(target Attribute) (damage int) {
	damage = monster.Attribute.Attack - target.Defense
	target.Health -= damage
	return damage
}

func (monster *Monster) GetDamage(enemy Attribute) (damage int) {
	damage = enemy.Attack - monster.Attribute.Defense
	monster.Attribute.Health -= damage
	return
}

func (monster *Monster) Carve() (constants.ItemType, int) {
	var dropItems DropItemSlice = monster.DropItems
	itemType, itemNum := dropItems.GetItemByPercentage()
	return itemType, itemNum
}
