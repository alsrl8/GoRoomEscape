package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

func isMonsterExistInRoom(room *structure.Room) bool {
	return room.Monster != nil
}

func printMonsterInRoom(room *structure.Room) {
	monster := room.Monster
	printMonster(monster)
}

func printMonster(monster *structure.Monster) {
	fmt.Printf("%s >> Health: %d, Attack: %d, Defense: %d\n", constants.MonsterTypeStringMap[monster.MonsterType], monster.Attribute.Health, monster.Attribute.Attack, monster.Attribute.Defense)
}

func reduceMonsterHealth(monster *structure.Monster, attack int) {
	reduceHealth(&monster.Attribute, attack)
}

func AttackMonsterInRoom(status *structure.Status, monsterName string) {
	room := status.Room
	if !isMonsterExistInRoom(room) {
		fmt.Println(constants.NoSuchMonster)
		return
	} else if room.Monster.MonsterType != constants.StringMonsterTypeMap[monsterName] {
		fmt.Println(constants.NoSuchMonster)
		return
	}
	monster := room.Monster
	reduceMonsterHealth(monster, status.Attribute.Attack)
	reduceHealth(&status.Attribute, monster.Attribute.Attack)
}
