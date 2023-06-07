package terminal

import (
	"fmt"
	"goproject/constants"
	"goproject/game/command"
	"goproject/game/data"
	"goproject/structure"
	"regexp"
)

func battleToMonster(status *structure.Status, monster *structure.Monster) (gameOverFlag bool) {
	clearTerminal()
	fmt.Printf(constants.DuringBattle, constants.MonsterTypeStringMap[monster.MonsterType])
	for {
		room := command.GetCurrentRoom(status)
		command.ShowUserNameAndStatus(status)
		command.PrintMonsterInRoom(room)
		input := getInput()
		clearTerminal()
		switch input {
		case "공격":
			command.DamageMonsterByPlayer(status, monster)
			if command.IsDead(monster.Attribute) {
				command.RemoveMonsterInRoom(room)
				command.CarveMonster(status, monster)
				fmt.Printf(constants.KillMonster, constants.MonsterTypeStringMap[monster.MonsterType])
				goto FinishTheBattle
			}
		case "도망":
			runSuccessFlag := command.Run(data.RunSuccessPercentage)
			if runSuccessFlag {
				fmt.Println(constants.RunningSucceeded)
				goto FinishTheBattle
			} else {
				fmt.Println(constants.RunningFailed)
			}
		default:
			reg, _ := regexp.Compile(" 사용$")
			if reg.MatchString(input) {
				itemName := reg.ReplaceAllString(input, "")
				itemType := constants.StringItemTypeMap[itemName]
				if err := command.ValidateItemUsability(status.Inventory, itemType, false); err != nil {
					fmt.Println(err.Error())
					continue
				}
				command.UseItemByName(status, itemName)
				fmt.Printf(constants.UseItem, itemName)
			} else {
				fmt.Println(constants.WrongInput, input)
				continue
			}
		}
		command.DamagePlayerByMonster(status, monster)
		if command.IsDead(status.Attribute) {
			gameOverFlag = true
			fmt.Printf(constants.GetKilled)
			goto FinishTheBattle
		}
	}
FinishTheBattle:
	return
}
