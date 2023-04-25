package main

func initDoorGrid() [6][7][4]int {
	var grid [6][7][4]int
	grid[0][4][East] = Clear
	grid[0][4][South] = ClosedDoor
	grid[0][5][East] = Clear
	grid[0][5][West] = Clear
	grid[0][6][East] = LockedDoor
	grid[0][6][West] = Clear
	grid[1][3][East] = Clear
	grid[1][3][South] = Clear
	grid[1][4][North] = ClosedDoor
	grid[1][4][West] = Clear
	grid[2][0][East] = Clear
	grid[2][1][East] = GlassDoor
	grid[2][1][West] = Clear
	grid[2][1][South] = Clear
	grid[2][2][East] = Clear
	grid[2][2][West] = GlassDoor
	grid[2][3][North] = Clear
	grid[2][3][West] = Clear
	grid[2][3][South] = Clear
	grid[3][1][North] = Clear
	grid[3][1][South] = Clear
	grid[3][3][North] = Clear
	grid[3][3][South] = Clear
	grid[4][1][North] = Clear
	grid[4][1][South] = Clear
	grid[4][3][East] = ClosedDoor
	grid[4][3][North] = Clear
	grid[4][4][West] = ClosedDoor
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

func initUser() User {
	user := User{
		pos: Position{
			r: 5,
			c: 1,
		},
		items: []uint{0, 0, 0},
	}
	return user
}
