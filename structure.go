package main

type Position struct {
	r int
	c int
}

type User struct {
	pos   Position
	items []uint
}
