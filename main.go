package main

func gameStart() {
	doorGrid := initDoorGrid()
	itemGrid := initItemGrid()
	user := initUser()
	initTerminal(doorGrid, itemGrid, user)
}

func main() {
	gameStart()
}
