package main

func initDoorGrid() [6][7][4]int {
	var grid [6][7][4]int
	grid[0][4][East] = Clear
	grid[0][4][South] = WoodDoorClosed
	grid[0][5][East] = Clear
	grid[0][5][West] = Clear
	grid[0][6][East] = LockedDoorClosed
	grid[0][6][West] = Clear
	grid[1][3][East] = Clear
	grid[1][3][South] = Clear
	grid[1][4][North] = WoodDoorClosed
	grid[1][4][West] = Clear
	grid[2][0][East] = Clear
	grid[2][1][East] = GlassDoorClosed
	grid[2][1][West] = Clear
	grid[2][1][South] = Clear
	grid[2][2][East] = Clear
	grid[2][2][West] = GlassDoorClosed
	grid[2][3][North] = Clear
	grid[2][3][West] = Clear
	grid[2][3][South] = Clear
	grid[3][1][North] = Clear
	grid[3][1][South] = Clear
	grid[3][3][North] = Clear
	grid[3][3][South] = Clear
	grid[4][1][North] = Clear
	grid[4][1][South] = Clear
	grid[4][3][East] = WoodDoorClosed
	grid[4][3][North] = Clear
	grid[4][4][West] = WoodDoorClosed
	grid[4][4][South] = Clear
	grid[5][1][North] = Clear
	grid[5][4][North] = Clear

	return grid
}

func initItemGrid() [6][7]int {
	var grid [6][7]int
	grid[2][0] = Hammer
	grid[5][4] = Key
	return grid
}

func initPosition() Position {
	position := Position{
		r: 5,
		c: 1,
		d: North,
	}
	return position
}
