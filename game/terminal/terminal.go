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
	cmd.Run()
}

func printLine() {
	fmt.Println("==================================================")
}

// Terminal에서 사용자 입력을 받는다.(Windows)
func getInput() string {
	fmt.Printf("입력 >>> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\r')
	input = strings.TrimSuffix(input, "\r")
	return input
}

func getNextRoomInfo(room *structure.Room, direction constants.Direction) string {
	if room.Doors[direction] != nil {
		doorType := room.Doors[direction].DoorType
		isClosed := room.Doors[direction].IsClosed
		return fmt.Sprintf("%s(%s)", constants.DoorTypeStringMap[doorType], constants.DoorCloseStateStringMap[isClosed])
	}

	if room.Directions[direction] == nil {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.EmptyRoom])
	}
}

func showRoomInfo(room *structure.Room, bag *map[constants.ItemType]int) {
	printLine()
	for _, d := range constants.DirectionList {
		fmt.Printf("%s(%s) - %s\n", constants.DirStringMap[d], constants.DirStringEngMap[d], getNextRoomInfo(room, d))
	}

	printLine()
	fmt.Printf("아이템 정보 >>> ")
	for itemType, itemNum := range *bag {
		if itemNum == 0 {
			continue
		}
		fmt.Printf("%s(%d) ", constants.ItemTypeStringMap[itemType], itemNum)
	}
	fmt.Println()

	printLine()
	fmt.Printf("이동 가능한 방향 >>> ")
	for _, d := range constants.DirectionList {
		if !command.CanMove(room, d) {
			continue
		}
		fmt.Printf("%s(%s) ", constants.DirStringMap[d], constants.DirStringEngMap[d])
	}
	fmt.Println()

	printLine()
}

func move(room *structure.Room, direction constants.Direction) *structure.Room {
	if !command.CanMove(room, direction) {
		fmt.Println(constants.CanNotMoveWarning)
		return room
	}
	return command.Move(room, direction)
}

func openDoorByName(room *structure.Room, doorName string) {
	door := command.FindDoorByName(room, doorName)
	if door == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
		return
	} else if !command.IsItemsEnoughToOpenDoor(door) {
		fmt.Println(constants.NotEnoughItemsToOpenDoor, doorName)
		return
	}
	door.IsClosed = false
}

func closeDoorByName(room *structure.Room, doorName string) {
	door := command.FindDoorByName(room, doorName)
	if door == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
		return
	} else if door.IsClosed {
		fmt.Println(constants.AlreadyClosedDoor, doorName)
		return
	}

	door.IsClosed = true
}

func newBag() *map[constants.ItemType]int {
	bag := make(map[constants.ItemType]int)
	return &bag
}

func pickUpItems(room *structure.Room, bag *map[constants.ItemType]int) {
	for itemType, itemNum := range room.Items {
		if itemNum == 0 {
			continue
		}
		fmt.Printf("%s을 (%d)개 주웠습니다.\n", constants.ItemTypeStringMap[itemType], itemNum)
		room.Items[itemType] -= itemNum
		(*bag)[itemType] += itemNum
	}
}

func RunTerminal(startRoom *structure.Room) {
	room := startRoom
	bag := newBag()
	clearTerminal()

	for {
		showRoomInfo(room, bag)
		input := getInput()
		clearTerminal()
		switch input {
		case "q", "quit", "Q", "그만":
			goto QuitLoop
		case "e", "east", "E", "동", "동 가", "동쪽으로 가":
			room = move(room, constants.East)
		case "n", "north", "N", "북", "북 가", "북쪽으로 가":
			room = move(room, constants.North)
		case "w", "west", "W", "서", "서 가", "서쪽으로 가":
			room = move(room, constants.West)
		case "s", "south", "S", "남", "남 가", "남쪽으로 가":
			room = move(room, constants.South)
		case "l", "look", "L", "봐":
			// showRoomInfo(doorGrid, &itemGrid, &user)
		case "i", "item", "I", "소지품":
			// showOwnedItems(user)
		default:
			// 아이템
			reg, _ := regexp.Compile(" 봐$")
			if reg.MatchString(input) {
				item := strings.TrimSuffix(input, " 봐")
				fmt.Println(item)
			} else {
				// 열기
				reg, _ = regexp.Compile("( 열기| 열어| 열)$")
				if reg.MatchString(input) {
					doorName := reg.ReplaceAllString(input, "")
					openDoorByName(room, doorName)
				} else {
					// 닫기
					reg, _ = regexp.Compile("( 닫기| 닫어| 닫)$")
					if reg.MatchString(input) {
						doorName := reg.ReplaceAllString(input, "")
						closeDoorByName(room, doorName)
					} else {
						fmt.Println(constants.WrongInput, input)
					}
				}
			}
		}
		if room.IsGoal {
			goto ExitLoop
		}
		pickUpItems(room, bag)
	}
QuitLoop:
	clearTerminal()
	fmt.Println("종료했습니다.")
	return
ExitLoop:
	clearTerminal()
	fmt.Println("탈출했습니다.")
}
