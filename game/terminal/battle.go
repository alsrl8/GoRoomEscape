package terminal

import (
	"fmt"
	"goproject/constants"
	"goproject/game/command"
	"goproject/structure"
)

func battleToMonster(status *structure.Status, monster *structure.Monster) (gameOverFlag bool) {
	clearTerminal()
	fmt.Printf(constants.DuringBattle, constants.MonsterTypeStringMap[monster.MonsterType])
	for {
		command.PrintUserNameAndStatus(status)
		command.PrintMonsterInRoom(status.Room)
		input := getInput()
		clearTerminal()
		switch input {
		case "공격":
			command.DamageMonsterByPlayer(status, monster)
			if command.IsDead(monster.Attribute) {
				command.RemoveMonsterInRoom(status.Room)
				command.CarveMonster(status, monster)
				fmt.Printf(constants.KillMonster, constants.MonsterTypeStringMap[monster.MonsterType])
				goto FinishTheBattle
			}
			command.DamagePlayerByMonster(status, monster)
			if command.IsDead(status.Attribute) {
				gameOverFlag = true
				fmt.Printf(constants.GetKilled)
				goto FinishTheBattle
			}
		case "방어":
		case "도망":
		default:
			fmt.Println(constants.WrongInput, input)
		}

	}
FinishTheBattle:
	return
}
