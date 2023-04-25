package main

func canMove(grid [6][7][4]int, pos *Position, dir int) bool {
	switch grid[pos.r][pos.c][dir] {
	case Clear:
		return true
	case OpenDoor:
		return true
	default:
		return false
	}
}
