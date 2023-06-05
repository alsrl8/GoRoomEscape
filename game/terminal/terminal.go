package terminal

import (
	"bufio"
	"fmt"
	"goproject/constants"
	"goproject/game/command"
	"goproject/structure"
	"os"
	"os/exec"
	"strings"
)

// Terminal 내용을 Clear 한다.(Windows)
func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return
	}
}

// Terminal 사용자 입력을 받는다.(Windows)
func getInput() string {
	fmt.Printf("입력 >>> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\r')
	input = strings.TrimSuffix(input, "\r")
	return input
}

func SetUserInputAsUserName(status *structure.Status) {
	clearTerminal()
	fmt.Println("사용자 이름")
	for {
		userName := getInput()
		if len(userName) > 0 && len(userName) <= 16 {
			status.Name = userName
			return
		} else {
			clearTerminal()
			fmt.Println(constants.InvalidUserName, userName)
		}
	}
}

func RunTerminal(status *structure.Status) {
	clearTerminal()
	for {
		command.ShowRoomAndInventoryInfo(status)
		input := getInput()
		clearTerminal()

		var ret structure.CommandResult
		tokenLen := len(strings.Split(input, " "))
		switch tokenLen {
		case 1:
			ret = handleSingleTokenCommand(input, status)
		default:
			ret = HandleMultiTokenCommand(input, status)
		}

		if ret.QuitLoopFlag {
			goto QuitLoop
		} else if ret.ExitLoopFlag {
			goto ExitLoop
		} else if ret.GameOverFlag {
			goto GameOver
		}

		if status.Room.GoalFlag {
			goto ExitLoop
		}
	}
QuitLoop:
	clearTerminal()
	fmt.Println(constants.QuitGame)
GameOver:
	clearTerminal()
	fmt.Println(constants.GameOver)
	return
ExitLoop:
	clearTerminal()
	fmt.Println(constants.GameClear)
}

func handleSingleTokenCommand(input string, status *structure.Status) (ret structure.CommandResult) {
	switch input {
	case "Q", "q":
		ret.QuitLoopFlag = true
	case "E", "e":
		status.Room = command.Move(status.Room, constants.East)
	case "N", "n":
		status.Room = command.Move(status.Room, constants.North)
	case "W", "w":
		status.Room = command.Move(status.Room, constants.West)
	case "S", "s":
		status.Room = command.Move(status.Room, constants.South)
	case "EQ", "eq":
		command.ShowBodyParts(*status)
	}
	return
}

func HandleMultiTokenCommand(input string, status *structure.Status) (ret structure.CommandResult) {
	tokens := strings.Split(input, " ")
	lastToken := tokens[len(tokens)-1]
	switch lastToken {
	case "사용":
		itemName := tokens[0]
		itemType := constants.StringItemTypeMap[itemName]
		if err := command.ValidateItemUsability(status.Inventory, itemType, false); err != nil {
			fmt.Println(err.Error())
			return
		}
		command.UseItemByName(status, itemName)
		fmt.Printf(constants.UseItem, itemName)
	case "열어", "열":
		doorName := tokens[0]
		doorType := constants.StringDoorTypeMap[doorName]
		if err := command.ValidateDoorExist(status.Room, doorType); err != nil {
			fmt.Println(err.Error())
			return
		}
		command.OpenDoorByName(status.Room, doorType)
	case "닫아", "닫":
		doorName := tokens[0]
		doorType := constants.StringDoorTypeMap[doorName]
		if err := command.ValidateDoorExist(status.Room, doorType); err != nil {
			fmt.Println(err.Error())
			return
		}
		command.CloseDoorByName(status.Room, doorType)
	case "보다", "봐":
	case "입어", "장비":
		itemName := tokens[0]
		itemType := constants.StringItemTypeMap[itemName]
		command.Equip(status, itemType)
	case "벗어":
		itemName := tokens[0]
		itemType := constants.StringItemTypeMap[itemName]
		command.Disarm(status, itemType)
	case "풀어":
		tokenLen := len(tokens)
		switch tokenLen {
		case 3:
			itemName, doorName := tokens[0], tokens[1]
			itemType := constants.StringItemTypeMap[itemName]
			doorType := constants.StringDoorTypeMap[doorName]
			if err := command.ValidateItemUsability(status.Inventory, itemType, true); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateDoorExist(status.Room, doorType); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateItemDoorMatch(itemType, doorType); err != nil {
				fmt.Println(err.Error())
			}
			command.UseItemToDoorByName(status.Room, status.Inventory, itemName, doorName)
			fmt.Printf(constants.UseItem, itemName)
		case 5:
			itemName, doorName := tokens[0], tokens[2]
			//itemNum, doorNum := tokens[1], tokens[3]	TODO 문과 아이템 번호가 주어졌을 때 `풀어` 명령어 처리
			itemType := constants.StringItemTypeMap[itemName]
			doorType := constants.StringDoorTypeMap[doorName]
			if err := command.ValidateItemUsability(status.Inventory, itemType, true); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateDoorExist(status.Room, doorType); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateItemDoorMatch(itemType, doorType); err != nil {
				fmt.Println(err.Error())
			}
			command.UseItemToDoorByName(status.Room, status.Inventory, itemName, doorName)
			fmt.Printf(constants.UseItem, itemName)
		default:
			fmt.Println(constants.WrongInput, input)
		}
	case "공격":
		tokenLen := len(tokens)
		switch tokenLen {
		case 2:
			monsterName := tokens[0]
			monster := command.FindMonsterByName(status.Room, monsterName)
			if monster == nil {
				fmt.Println(constants.NoSuchMonster, monsterName)
				return
			}
			gameOverFlag := battleToMonster(status, monster)
			if gameOverFlag {
				ret.GameOverFlag = true
			}
		case 3:
			monsterName := tokens[0]
			//monsterNum := tokens[1] TODO 몬스터 번호가 주어졌을 때 `공격` 명령어 처리
			monster := command.FindMonsterByName(status.Room, monsterName)
			if monster == nil {
				fmt.Println(constants.NoSuchMonster, monsterName)
				return
			}
			gameOverFlag := battleToMonster(status, monster)
			if gameOverFlag {
				ret.GameOverFlag = true
			}
		}
	default:
		fmt.Println(constants.WrongInput, input)
	}
	return
}
