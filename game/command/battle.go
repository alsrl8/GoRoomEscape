package command

import (
	"goproject/structure"
	"math"
	"math/rand"
	"time"
)

func reduceHealth(attribute *structure.Attribute, attack int) int {
	damage := int(math.Max(0, float64(attack-attribute.Defense)))
	attribute.Health -= damage
	return damage
}

func IsDead(attribute structure.Attribute) bool {
	return attribute.Health <= 0
}

func Run(successPercentage float64) bool {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Float64()
	return randomNum <= successPercentage
}
