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
		command.PrintUserNameAndStatus(status)
		command.PrintMonsterInRoom(status.Room)
		input := getInput()
		clearTerminal()
		switch input {
		case "공격":
			if !command.IsAttackAble(status) {
				fmt.Println(constants.CanNotAttack)
				continue
			}
			command.DamageMonsterByPlayer(status, monster)
			if command.IsDead(monster.Attribute) {
				command.RemoveMonsterInRoom(status.Room)
				command.CarveMonster(status, monster)
				fmt.Printf(constants.KillMonster, constants.MonsterTypeStringMap[monster.MonsterType])
				goto FinishTheBattle
			}
		case "방어":
			if !command.IsGuardAble(status) {
				fmt.Println(constants.CanNotGuard)
				continue
			}
			command.Guard(status)
			fmt.Println(constants.RaiseShield)
			continue
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
				err := command.ValidateItemUsability(status.Inventory, itemType)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				command.UseItemByName(status, itemName)
				command.DropGuard(status)
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
