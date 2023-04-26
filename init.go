package main

func initDoorGrid() [6][7][4][2]int {
	var grid [6][7][4][2]int
	grid[0][4][East][0] = Clear
	grid[0][4][South][0] = WoodDoor
	grid[0][4][South][1] = Closed
	grid[0][5][East][0] = Clear
	grid[0][5][West][0] = Clear
	grid[0][6][East][0] = LockedDoor
	grid[0][6][East][1] = Closed
	grid[0][6][West][0] = Clear
	grid[1][3][East][0] = Clear
	grid[1][3][South][0] = Clear
	grid[1][4][North][0] = WoodDoor
	grid[1][4][North][1] = Closed
	grid[1][4][West][0] = Clear
	grid[2][0][East][0] = Clear
	grid[2][1][East][0] = GlassDoor
	grid[2][1][West][0] = Clear
	grid[2][1][South][0] = Clear
	grid[2][2][East][0] = Clear
	grid[2][2][West][0] = GlassDoor
	grid[2][3][North][0] = Clear
	grid[2][3][West][0] = Clear
	grid[2][3][South][0] = Clear
	grid[3][1][North][0] = Clear
	grid[3][1][South][0] = Clear
	grid[3][3][North][0] = Clear
	grid[3][3][South][0] = Clear
	grid[4][1][North][0] = Clear
	grid[4][1][South][0] = Clear
	grid[4][3][East][0] = WoodDoor
	grid[4][3][East][1] = Closed
	grid[4][3][North][0] = Clear
	grid[4][4][West][0] = WoodDoor
	grid[4][4][West][1] = Closed
	grid[4][4][South][0] = Clear
	grid[5][1][North][0] = Clear
	grid[5][4][North][0] = Clear

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
