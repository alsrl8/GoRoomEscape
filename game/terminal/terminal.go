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
	err := cmd.Run()
	if err != nil {
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
	userName := getInput()
	status.Name = userName
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
		case "도망":
			continue // TODO 도망 명령어
		// Multiple Words
		default:
			// 아이템
			reg, _ := regexp.Compile(" 사용$")
			if reg.MatchString(input) {
				itemName := reg.ReplaceAllString(input, "")
				command.UseItem(status.Inventory, itemName)
				continue
			}
			// 아이템 With Target
			reg, _ = regexp.Compile(" 사용 ")
			if reg.MatchString(input) {
				s := strings.Split(input, " 사용 ")
				itemName, doorName := s[0], s[1]
				command.UseItemToDoor(status.Room, status.Inventory, itemName, doorName)
				continue
			}
			// 열기
			reg, _ = regexp.Compile("( 열기| 열어| 열)$")
			if reg.MatchString(input) {
				doorName := reg.ReplaceAllString(input, "")
				command.OpenDoorByName(status.Room, status.Inventory, doorName)
				continue
			}
			// 닫기
			reg, _ = regexp.Compile("( 닫기| 닫아| 닫)$")
			if reg.MatchString(input) {
				doorName := reg.ReplaceAllString(input, "")
				command.CloseDoorByName(status.Room, doorName)
				continue
			}
			// 착용
			reg, _ = regexp.Compile("( 착용| 장비)$")
			if reg.MatchString(input) {
				itemName := reg.ReplaceAllString(input, "")
				command.Equip(status, itemName)
				continue
			}
			// 해제
			reg, _ = regexp.Compile(" 해제$")
			if reg.MatchString(input) {
				bodyPartName := reg.ReplaceAllString(input, "")
				command.Disarm(status, bodyPartName)
				continue
			}
			// 공격
			reg, _ = regexp.Compile(" 공격$")
			if reg.MatchString(input) {
				monsterName := reg.ReplaceAllString(input, "")
				monster := command.GetMonsterInRoomByName(status.Room, monsterName)
				if monster == nil {
					fmt.Println(constants.NoSuchMonster, monsterName)
					continue
				}
				command.AttackMonster(status, monster)
				if command.IsDead(status.Attribute) {
					goto GameOver
				} else if command.IsDead(monster.Attribute) {
					command.RemoveMonsterInRoom(status.Room)
					command.CarveMonster(status, monster)
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
	fmt.Println("종료했습니다.")
	return
GameOver:
	clearTerminal()
	fmt.Println("실패했습니다.")
	return
ExitLoop:
	clearTerminal()
	fmt.Println("탈출했습니다.")
}
