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
		isClosed := room.Doors[direction].Closed
		return fmt.Sprintf("%s(%s)", constants.DoorTypeStringMap[doorType], constants.DoorCloseStateStringMap[isClosed])
	}

	if room.Directions[direction] == nil {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.Wall])
	} else {
		return fmt.Sprintf(constants.SpaceTypeStringMap[constants.EmptyRoom])
	}
}

func showRoomInfo(room *structure.Room, bag *structure.Inventory) {
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

func openDoorByName(room *structure.Room, bag *structure.Inventory, doorName string) {
	door := command.FindDoorByName(room, doorName)
	if door == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
		return
	} else if !door.Closed {
		fmt.Println(constants.AlreadyOpenDoor, doorName)
		return
	}

	command.OpenDoor(door, bag)
}

func closeDoorByName(room *structure.Room, doorName string) {
	door := command.FindDoorByName(room, doorName)
	if door == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
		return
	} else if door.Closed {
		fmt.Println(constants.AlreadyClosedDoor, doorName)
		return
	}

	door.Closed = true
}

func newBag() *structure.Inventory {
	bag := make(structure.Inventory)
	return &bag
}

func pickUpItems(room *structure.Room, bag *structure.Inventory) {
	for itemType, itemNum := range room.Items {
		if itemNum == 0 {
			continue
		}
		fmt.Printf("%s을 (%d)개 주웠습니다.\n", constants.ItemTypeStringMap[itemType], itemNum)
		room.Items[itemType] -= itemNum
		(*bag)[itemType] += itemNum
	}
}

func useItem(room *structure.Room, bag *structure.Inventory, itemName string, doorName string) {
	if !command.IsItemInBag(bag, itemName) {
		fmt.Println(constants.NoSuchItem, itemName)
	} else if command.FindDoorByName(room, doorName) == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
	}

	if constants.Hammer == constants.StringItemTypeMap[itemName] {
		if constants.GlassDoor == constants.StringDoorTypeMap[doorName] {
			command.BreakGlassDoorAndReduceHammer(room, bag)
			return
		}
	}

	if constants.Key == constants.StringItemTypeMap[itemName] {
		if constants.LockedDoor == constants.StringDoorTypeMap[doorName] {
			command.UnlockLockedDoorAndReduceKey(room, bag)
			return
		}
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
		case "장비":

		default:
			// 아이템
			reg, _ := regexp.Compile(" 사용 ")
			if reg.MatchString(input) {
				s := strings.Split(input, " 사용 ")
				itemName, doorName := s[0], s[1]
				useItem(room, bag, itemName, doorName)
				continue
			}
			// 열기
			reg, _ = regexp.Compile("( 열기| 열어| 열)$")
			if reg.MatchString(input) {
				doorName := reg.ReplaceAllString(input, "")
				openDoorByName(room, bag, doorName)
				continue
			}
			// 닫기
			reg, _ = regexp.Compile("( 닫기| 닫아| 닫)$")
			if reg.MatchString(input) {
				doorName := reg.ReplaceAllString(input, "")
				closeDoorByName(room, doorName)
				continue
			}
			fmt.Println(constants.WrongInput, input)
		}
		if room.GoalFlag {
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
