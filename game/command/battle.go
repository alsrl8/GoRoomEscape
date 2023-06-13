package command

import (
	"goproject/structure"
	"math/rand"
	"time"
)

func IsDead(attribute structure.Attribute) bool {
	return attribute.Health <= 0
}

func Run(successPercentage float64) bool {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Float64()
	return randomNum <= successPercentage
}
