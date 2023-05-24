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
	fmt.Printf("%s >> Health: %d, Attack: %d, Defense: %d\n", constants.MonsterTypeStringMap[monster.MonsterType], monster.Attribute.Health, monster.Attribute.Attack, monster.Attribute.Defense)
}

func reduceMonsterHealth(monster *structure.Monster, attack int) {
	reduceHealth(&monster.Attribute, attack)
}

func AttackMonster(status *structure.Status, monster *structure.Monster) {
	reduceMonsterHealth(monster, status.Attribute.Attack)
	reduceHealth(&status.Attribute, monster.Attribute.Attack)
	fmt.Printf("몬스터(%s)와 공격을 주고 받았습니다.\n", constants.MonsterTypeStringMap[monster.MonsterType])
}

func RemoveMonsterInRoom(room *structure.Room) {
	room.Monster = nil
}

func CarveMonster(status *structure.Status, monster *structure.Monster) {
	itemType, itemNum := GetItemByPercentage(&monster.DropItem)
	if itemType == constants.Nothing {
		return
	}
	(*status.Inventory)[itemType] += itemNum
	fmt.Printf("다음을 얻었습니다 >> %s: %d개\n", constants.ItemTypeStringMap[itemType], itemNum)
}
