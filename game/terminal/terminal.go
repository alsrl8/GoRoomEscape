package terminal

import (
	"bufio"
	"fmt"
	"goproject/constants"
	"goproject/game/command"
	"goproject/structure"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func clearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return
		}
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	default:
		fmt.Println(constants.UnsupportedOperatingSystem)
	}
}

func getInput() (input string) {
	fmt.Printf("입력 >>> ")
	reader := bufio.NewReader(os.Stdin)
	switch runtime.GOOS {
	case "windows":
		input, _ = reader.ReadString('\r')
		input = strings.TrimSuffix(input, "\r")
	case "linux", "darwin":
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
	default:
		fmt.Println(constants.UnsupportedOperatingSystem)
	}
	return
}

func printTime() {
	currentTime := time.Now()
	fmt.Println("현재 시각:", currentTime)
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
	var inputChan chan string
	inputChan = make(chan string, 1)

	for {
		command.PrintLine()
		go func() {
			inputChan <- getInput()
		}()

		input := <-inputChan
		clearTerminal()
		fmt.Println("당신의 입력 >>>", input)

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
		} else if ret.GameOverFlag {
			goto GameOver
		}
	}
QuitLoop:
	clearTerminal()
	fmt.Println(constants.QuitGame)
	return
GameOver:
	clearTerminal()
	fmt.Println(constants.GameOver)
	return
}

func isMaxNumCommand(cmd string) bool {
	switch cmd {
	case "모두", "다", "전부":
		return true
	default:
		return false
	}
}

func handlePickUpItemCommand(status *structure.Status, tokens []string) {
	itemName := tokens[0]
	itemType, has := constants.StringItemTypeMap[itemName]
	if !has {
		fmt.Println(constants.NoSuchItem, itemName)
		return
	}

	var itemNum int
	if len(tokens) == 2 {
		itemNum = 1
	} else if isMaxNumCommand(tokens[1]) {
		room := status.Location.(*structure.Room)
		itemNum = room.GetItemNumber(itemType)
	} else {
		var err error
		itemNum, err = strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Println(constants.WrongInput, strings.Join(tokens, ""))
			return
		}
	}
	command.PickupItem(status, itemType, itemNum)
}

func handleThrowItemCommand(status *structure.Status, tokens []string) {
	itemName := tokens[0]
	itemType, has := constants.StringItemTypeMap[itemName]
	if !has {
		fmt.Println(constants.NoSuchItem, itemName)
		return
	}

	var itemNum int
	if len(tokens) == 2 {
		itemNum = 1
	} else if isMaxNumCommand(tokens[1]) {
		itemNum = status.Inventory.GetItemNumber(itemType)
	} else {
		var err error
		itemNum, err = strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Println(constants.WrongInput, strings.Join(tokens, ""))
			return
		}
		err = command.ValidateItemExist(status.Inventory, itemType, itemNum)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	command.ThrowItem(status, itemType, itemNum)
}

func handleDiscardItemCommand(status *structure.Status, tokens []string) {
	itemName := tokens[0]
	itemType, has := constants.StringItemTypeMap[itemName]
	if !has {
		fmt.Println(constants.NoSuchItem, itemName)
		return
	}

	var itemNum int
	if len(tokens) == 2 {
		itemNum = 1
	} else if isMaxNumCommand(tokens[1]) {
		itemNum = status.Inventory.GetItemNumber(itemType)
	} else {
		if err := command.ValidateItemExist(status.Inventory, itemType, itemNum); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	command.DiscardItem(status.Inventory, itemType, itemNum)
}

func handleSingleTokenCommand(input string, status *structure.Status) (ret structure.CommandResult) {
	switch input {
	case "Q", "q", "quit":
		ret.QuitLoopFlag = true
	case "E", "e", "동", "동쪽":
		command.Move(status, constants.East)
	case "N", "n", "북", "북쪽":
		command.Move(status, constants.North)
	case "W", "w", "서", "서쪽":
		command.Move(status, constants.West)
	case "S", "s", "남", "남쪽":
		command.Move(status, constants.South)
	case "정보":
		command.ShowUserNameAndStatus(status)
	case "소지":
		command.ShowInventory(status.Inventory)
	case "장비", "EQ", "eq":
		command.ShowBodyParts(*status)
	case "보다", "봐":
		status.Location.ShowInfo()
	case "시간":
		printTime()
	case "입장":
		command.EnterDungeon(status)
		command.PrintLine()
		fmt.Println("던전에 입장했습니다.")
	case "퇴장":
		command.ExitDungeon(status)
		command.PrintLine()
		fmt.Println("던전에서 퇴장했습니다.")
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
		room := command.GetCurrentRoom(status)
		if err := command.ValidateDoorExist(room, doorType); err != nil {
			fmt.Println(err.Error())
			return
		}
		command.OpenDoorByName(room, doorType)
	case "닫아", "닫":
		doorName := tokens[0]
		doorType := constants.StringDoorTypeMap[doorName]
		room := command.GetCurrentRoom(status)
		if err := command.ValidateDoorExist(room, doorType); err != nil {
			fmt.Println(err.Error())
			return
		}
		command.CloseDoorByName(room, doorType)
	case "보다", "봐":
		tokenLen := len(tokens)
		switch tokenLen {
		case 2:
			// TODO 물건/몬스터 정보를 보는 기능 추가
		}
	case "주워":
		handlePickUpItemCommand(status, tokens)
	case "버려":
		handleThrowItemCommand(status, tokens)
	case "입어", "장비":
		itemName := tokens[0]
		itemType := constants.StringItemTypeMap[itemName]
		command.Equip(status, itemType)
	case "벗어":
		itemName := tokens[0]
		itemType := constants.StringItemTypeMap[itemName]
		command.Disarm(status, itemType)
	case "태워":
		handleDiscardItemCommand(status, tokens)
	case "풀어":
		tokenLen := len(tokens)
		switch tokenLen {
		case 3:
			itemName, doorName := tokens[0], tokens[1]
			itemType := constants.StringItemTypeMap[itemName]
			doorType := constants.StringDoorTypeMap[doorName]
			room := command.GetCurrentRoom(status)
			if err := command.ValidateItemUsability(status.Inventory, itemType, true); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateDoorExist(room, doorType); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateItemDoorMatch(itemType, doorType); err != nil {
				fmt.Println(err.Error())
			}
			command.UseItemToDoorByName(room, status.Inventory, itemName, doorName)
			fmt.Printf(constants.UseItem, itemName)
		case 5:
			itemName, doorName := tokens[0], tokens[2]
			//itemNum, doorNum := tokens[1], tokens[3]	TODO 문과 아이템 번호가 주어졌을 때 `풀어` 명령어 처리
			itemType := constants.StringItemTypeMap[itemName]
			doorType := constants.StringDoorTypeMap[doorName]
			room := command.GetCurrentRoom(status)
			if err := command.ValidateItemUsability(status.Inventory, itemType, true); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateDoorExist(room, doorType); err != nil {
				fmt.Println(err.Error())
				return
			} else if err = command.ValidateItemDoorMatch(itemType, doorType); err != nil {
				fmt.Println(err.Error())
			}
			command.UseItemToDoorByName(room, status.Inventory, itemName, doorName)
			fmt.Printf(constants.UseItem, itemName)
		default:
			fmt.Println(constants.WrongInput, input)
		}
	case "공격":
		tokenLen := len(tokens)
		switch tokenLen {
		case 2:
			monsterName := tokens[0]
			room := command.GetCurrentRoom(status)
			monster := command.FindMonsterByName(room, monsterName)
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
			room := command.GetCurrentRoom(status)
			monster := command.FindMonsterByName(room, monsterName)
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
