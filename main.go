package main

import (
	"goproject/game/initialize"
	"goproject/game/terminal"
	"goproject/structure"
)

func gameStart(status *structure.Status) {
	terminal.SetUserInputAsUserName(status)
	terminal.RunTerminal(status)
}

func main() {
	status := initialize.InitGame()
	//status := initialize.InitGameAndReturnStatus()
	gameStart(status)
}
