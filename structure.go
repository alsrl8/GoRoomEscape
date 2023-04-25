package main

type Position struct {
	r uint
	c uint
}

type User struct {
	pos   Position
	items []uint
}
