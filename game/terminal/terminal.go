package terminal

import (
	"bufio"
	"fmt"
	"goproject/constants"
	"goproject/game/command"
	"goproject/structure"
	"os"
	"os/exec"
	"regexp"
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
		switch input {
		case "q", "quit", "Q", "그만":
			goto QuitLoop
		case "e", "east", "E", "동", "동 가", "동쪽으로 가":
			status.Room = command.Move(status.Room, constants.East)
		case "n", "north", "N", "북", "북 가", "북쪽으로 가":
			status.Room = command.Move(status.Room, constants.North)
		case "w", "west", "W", "서", "서 가", "서쪽으로 가":
			status.Room = command.Move(status.Room, constants.West)
		case "s", "south", "S", "남", "남 가", "남쪽으로 가":
			status.Room = command.Move(status.Room, constants.South)
		case "eq", "장비":
			command.ShowEquipment(*status)
		// Multiple Words
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
				continue
			}
			reg, _ = regexp.Compile(" 사용 ")
			if reg.MatchString(input) {
				s := strings.Split(input, " 사용 ")
				itemName, doorName := s[0], s[1]
				itemType := constants.StringItemTypeMap[itemName]
				doorType := constants.StringDoorTypeMap[doorName]
				if err := command.ValidateItemUsability(status.Inventory, itemType, true); err != nil {
					fmt.Println(err.Error())
					continue
				} else if err = command.ValidateDoorExist(status.Room, doorType); err != nil {
					fmt.Println(err.Error())
					continue
				} else if err = command.ValidateItemDoorMatch(itemType, doorType); err != nil {
					fmt.Println(err.Error())
				}
				command.UseItemToDoorByName(status.Room, status.Inventory, itemName, doorName)
				fmt.Printf(constants.UseItem, itemName)
				continue
			}
			reg, _ = regexp.Compile("( 열기| 열어| 열)$")
			if reg.MatchString(input) {
				doorName := reg.ReplaceAllString(input, "")
				doorType := constants.StringDoorTypeMap[doorName]
				if err := command.ValidateDoorExist(status.Room, doorType); err != nil {
					fmt.Println(err.Error())
					continue
				}
				command.OpenDoorByName(status.Room, doorType)
				continue
			}
			reg, _ = regexp.Compile("( 닫기| 닫아| 닫)$")
			if reg.MatchString(input) {
				doorName := reg.ReplaceAllString(input, "")
				doorType := constants.StringDoorTypeMap[doorName]
				if err := command.ValidateDoorExist(status.Room, doorType); err != nil {
					fmt.Println(err.Error())
					continue
				}
				command.CloseDoorByName(status.Room, doorType)
				continue
			}
			reg, _ = regexp.Compile("( 착용| 장비)$")
			if reg.MatchString(input) {
				itemName := reg.ReplaceAllString(input, "")
				command.Equip(status, itemName)
				continue
			}
			reg, _ = regexp.Compile(" 해제$")
			if reg.MatchString(input) {
				bodyPartName := reg.ReplaceAllString(input, "")
				command.Disarm(status, bodyPartName)
				continue
			}
			reg, _ = regexp.Compile(" 공격$")
			if reg.MatchString(input) {
				monsterName := reg.ReplaceAllString(input, "")
				monster := command.FindMonsterByName(status.Room, monsterName)
				if monster == nil {
					fmt.Println(constants.NoSuchMonster, monsterName)
					continue
				}
				gameOverFlag := battleToMonster(status, monster)
				if gameOverFlag {
					goto GameOver
				}
				continue
			}
			fmt.Println(constants.WrongInput, input)
		}
		if status.Room.GoalFlag {
			goto ExitLoop
		}
		command.PickUpItems(status.Room, status.Inventory)
	}
QuitLoop:
	clearTerminal()
	fmt.Println(constants.QuitGame)
	return
GameOver:
	clearTerminal()
	fmt.Println(constants.GameOver)
	return
ExitLoop:
	clearTerminal()
	fmt.Println(constants.GameClear)
}
