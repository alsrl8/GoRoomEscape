package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

func FindMonsterByName(room *structure.Room, monsterName string) *structure.Monster {
	if !isMonsterExistInRoom(room) || room.Monster.MonsterType != constants.StringMonsterTypeMap[monsterName] {
		return nil
	}
	return room.Monster
}

func isMonsterExistInRoom(room *structure.Room) bool {
	return room.Monster != nil
}

func PrintMonsterInRoom(room *structure.Room) {
	monster := room.Monster
	printMonster(monster)
}

func printMonster(monster *structure.Monster) {
	PrintLine()
	fmt.Printf(constants.MonsterStatus, constants.MonsterTypeStringMap[monster.MonsterType], monster.Attribute.Health, monster.Attribute.Attack, monster.Attribute.Defense)
}

func reduceMonsterHealth(monster *structure.Monster, attack int) int {
	return reduceHealth(&monster.Attribute, attack)
}

func DamageMonsterByPlayer(status *structure.Status, monster *structure.Monster) {
	damage := reduceMonsterHealth(monster, status.Attribute.Attack)
	fmt.Printf(constants.AttackMonster, constants.MonsterTypeStringMap[monster.MonsterType])
	fmt.Printf(constants.DamageMonster, constants.MonsterTypeStringMap[monster.MonsterType], damage)
}

func DamagePlayerByMonster(status *structure.Status, monster *structure.Monster) {
	damage := reduceHealth(&status.Attribute, monster.Attribute.Attack)
	fmt.Printf(constants.GetAttackedByMonster, constants.MonsterTypeStringMap[monster.MonsterType])
	fmt.Printf(constants.DamageByMonster, constants.MonsterTypeStringMap[monster.MonsterType], damage)
}

func RemoveMonsterInRoom(room *structure.Room) {
	room.Monster = nil
}

func CarveMonster(status *structure.Status, monster *structure.Monster) {
	itemType, itemNum := GetItemByPercentage(&monster.DropItems)
	if itemType == constants.Nothing {
		return
	}
	(*status.Inventory)[itemType] += itemNum
	fmt.Printf(constants.GetItem, constants.ItemTypeStringMap[itemType], itemNum)
}
