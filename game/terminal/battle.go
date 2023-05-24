package terminal

import (
	"fmt"
	"goproject/constants"
	"goproject/game/command"
	"goproject/structure"
)

func battleToMonster(status *structure.Status, monster *structure.Monster) (gameOverFlag bool) {
	clearTerminal()
	fmt.Printf("<<< 몬스터(%s)와 전투 중입니다 >>>\n", constants.MonsterTypeStringMap[monster.MonsterType])
	for {
		command.PrintUserNameAndStatus(status)
		command.PrintMonsterInRoom(status.Room)
		input := getInput()
		clearTerminal()
		switch input {
		case "공격":
			command.AttackMonster(status, monster)
		case "방어":
		case "도망":
		default:
			fmt.Println(constants.WrongInput, input)
		}
		if command.IsDead(status.Attribute) {
			gameOverFlag = true
			fmt.Println("당신은 죽었습니다. ㅠㅠ")
			goto FinishTheBattle
		} else if command.IsDead(monster.Attribute) {
			command.RemoveMonsterInRoom(status.Room)
			command.CarveMonster(status, monster)
			fmt.Printf("당신은 %s를 죽였습니다. ㅠㅠ", constants.MonsterTypeStringMap[monster.MonsterType])
			goto FinishTheBattle
		}
	}
FinishTheBattle:
	return
}
