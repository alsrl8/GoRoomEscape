package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
	"math"
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

func attackFromMonsterInRoom(status *structure.Status) {
	room := status.Room
	if !isMonsterExistInRoom(room) {
		return
	}
	monster := room.Monster
	status.Attribute.Health -= int(math.Min(0, float64(status.Attribute.Defense-monster.Attribute.Attack)))
}

func attackMonster(status *structure.Status, monster *structure.Monster) {

}
