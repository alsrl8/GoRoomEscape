package structure

import "goproject/constants"

type Object struct {
	constants.ObjectType
}

type ObjectPositionAndType struct {
	Object
	Position
}
