package main

func canMove(grid [6][7][4]int, pos *Position) bool {
	switch grid[pos.r][pos.c][pos.d] {
	case Clear:
		return true
	case WoodDoorOpen:
		return true
	case GlassDoorOpen:
		return true
	case LockedDoorOpen:
		return true
	default:
		return false
	}
}
