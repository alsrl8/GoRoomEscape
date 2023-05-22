package command

import (
	"goproject/constants"
	"goproject/structure"
	"math/rand"
)

func GetItemByPercentage(dropItems *[]structure.DropItem) (constants.ItemType, int) {
	randomNum := rand.Float64()
	totalProbability := 0.0

	for _, dropItem := range *dropItems {
		totalProbability += dropItem.DropPercentage.Percentage
		if totalProbability >= randomNum {
			return dropItem.ItemType, dropItem.DropPercentage.Num
		}
	}
	return constants.Nothing, 0
}
