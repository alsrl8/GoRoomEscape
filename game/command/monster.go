package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

var monsterMap = map[constants.MonsterType]structure.Monster{
	constants.Squirrel: {
		Attribute: structure.Attribute{
			Health:  50,
			Attack:  5,
			Defense: 0},
		DropItem: []structure.DropItem{
			{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.7, Num: 1}},
		},
	},
	constants.Rabbit: {
		Attribute: structure.Attribute{
			Health:  70,
			Attack:  7,
			Defense: 3,
		},
		DropItem: []structure.DropItem{
			{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.5, Num: 1}},
			{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.3, Num: 2}}},
	},
	constants.Deer: {
		Attribute: structure.Attribute{
			Health:  100,
			Attack:  10,
			Defense: 5,
		},
		DropItem: []structure.DropItem{
			{ItemType: constants.Key, DropPercentage: structure.DropPercentage{Percentage: 1.0, Num: 1}}},
	},
}

func isMonsterExistInRoom(room *structure.Room) bool {
	for _, monsterNum := range room.Monsters {
		if monsterNum >= 1 {
			return true
		}
	}
	return false
}

func printMonstersInRoom(room *structure.Room) {
	for monsterType, monsterNum := range room.Monsters {
		if monsterNum <= 0 {
			continue
		}
		fmt.Printf("%s: %d\n", constants.MonsterTypeStringMap[monsterType], monsterNum)
	}
}
