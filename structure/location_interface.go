package structure

import (
	"goproject/constants"
)

type Location interface {
	Move(direction constants.Direction) *Location
	ShowInfo()
}
