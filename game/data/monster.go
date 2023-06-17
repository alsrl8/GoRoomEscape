package data

import (
	"goproject/constants"
	"goproject/structure"
)

func GetMonsterWithPositionData(stageNum constants.StageNum) []structure.MonsterWithPosition {
	switch stageNum {
	case constants.Stage0:
		return []structure.MonsterWithPosition{}
	case constants.Stage2:
		return []structure.MonsterWithPosition{
			{
				RoomPosition: structure.Position{Row: 3, Col: 4},
				Monster: structure.Monster{
					MonsterType: constants.Squirrel,
					Attribute: &structure.Attribute{
						Health:  50,
						Attack:  5,
						Defense: 0,
					},
					DropItems: []structure.DropItem{
						{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.7, Num: 1}},
					},
				},
			},
			{
				RoomPosition: structure.Position{Row: 4, Col: 6},
				Monster: structure.Monster{
					MonsterType: constants.Squirrel,
					Attribute: &structure.Attribute{
						Health:  50,
						Attack:  5,
						Defense: 0,
					},
					DropItems: []structure.DropItem{
						{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.7, Num: 1}},
					},
				},
			},
			{
				RoomPosition: structure.Position{Row: 6, Col: 1},
				Monster: structure.Monster{
					MonsterType: constants.Deer,
					Attribute: &structure.Attribute{
						Health:  100,
						Attack:  10,
						Defense: 5,
					},
					DropItems: []structure.DropItem{
						{ItemType: constants.Key, DropPercentage: structure.DropPercentage{Percentage: 1.0, Num: 1}},
					},
				},
			},
			{
				RoomPosition: structure.Position{Row: 6, Col: 4},
				Monster: structure.Monster{
					MonsterType: constants.Rabbit,
					Attribute: &structure.Attribute{
						Health:  70,
						Attack:  7,
						Defense: 3,
					},
					DropItems: []structure.DropItem{
						{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.5, Num: 1}},
						{ItemType: constants.HealPotion, DropPercentage: structure.DropPercentage{Percentage: 0.3, Num: 2}},
					},
				},
			},
		}
	default:
		panic("Invalid Dungeon Level")
	}
}
