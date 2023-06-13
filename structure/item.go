package structure

import (
	"goproject/constants"
	"math/rand"
	"time"
)

type Item struct {
	ItemType constants.ItemType
}

type ItemPositionAndType struct {
	RoomPosition Position
	ItemType     constants.ItemType
}

type DropPercentage struct {
	Percentage float64
	Num        int
}

type DropItem struct {
	ItemType       constants.ItemType
	DropPercentage DropPercentage
}

type DropItemSlice []DropItem

func (dropItems DropItemSlice) GetItemByPercentage() (constants.ItemType, int) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Float64()
	totalProbability := 0.0

	for _, dropItem := range dropItems {
		totalProbability += dropItem.DropPercentage.Percentage
		if totalProbability >= randomNum {
			return dropItem.ItemType, dropItem.DropPercentage.Num
		}
	}
	return constants.Nothing, 0
}
