package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Terminal 내용을 Clear 한다.(Windows)
func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Terminal에서 사용자 입력을 받는다.(Windows)
func getInput() string {
	fmt.Printf("입력 >>> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\r')
	input = strings.TrimSuffix(input, "\r")
	return input
}

// 방 위치, 방 안에 있는 아이템, 다음으로 이동 가능한 방향을 출력한다.
func showRoomInfo(doorGrid [6][7][4]int, itemGrid *[6][7]int, user *User) {
	pos := user.pos

	fmt.Println("=== 현재 방 정보 ===")
	fmt.Printf("r: %d\tc: %d\n", pos.r, pos.c)
	for d := 0; d < 4; d++ {
		fmt.Printf("%s(%s) - %s\n", DirStringMap[d], DirStringEngMap[d], WallStringMap[doorGrid[pos.r][pos.c][d]])
	}

	// 방 안에 있는 아이템 출력
	item := itemGrid[pos.r][pos.c]
	fmt.Println()
	fmt.Println("***", ItemStringMap[item], "***")

	// 아이템이 있었다면 아이템 목록에 추가하고 방에 있던 아이템 삭제
	user.items[item]++
	itemGrid[pos.r][pos.c] = Empty

	// 다음으로 이동 가능한 방향 출력
	fmt.Println()
	fmt.Printf("이동 가능한 장소 >>> ")
	for d := 0; d < 4; d++ {
		if canMove(doorGrid, &user.pos, d) {
			fmt.Printf("%s(%s) ", DirStringMap[d], DirStringEngMap[d])
		}
	}
	fmt.Println()

	fmt.Println("====================")
}

// 소지한 아이템 목록을 출력한다.
func showOwnedItems(user User) {
	fmt.Println("소지한 아이템 목록")
	for i := 1; i <= ItemTypeNum; i++ {
		fmt.Printf("%s: %d\n", ItemStringMap[i], user.items[i])
	}
	fmt.Println()
}

// 아이템이 존재하는 경우에만 아이템 설명을 출력한다.
func showItemInfo(user User, item string) {
	i, ok := StringItemMap[item]

	if !ok || user.items[i] == 0 {
		fmt.Println(NoSuchItem)
	} else {
		fmt.Println(ItemDescMap[i])
	}
}

func getDoorDirection(doorGrid [6][7][4]int, user User, door string) (d int) {
	wall := StringWallMap[door]
	for ; d < 4; d++ {
		switch doorGrid[user.pos.r][user.pos.c][d] {
		case wall:
			return d
		}
	}
	return -1
}

func gameStart() {

	doorGrid := initDoorGrid()
	itemGrid := initItemGrid()
	user := initUser()

	clearTerminal()
	for {
		showRoomInfo(doorGrid, &itemGrid, &user)
		input := getInput()
		clearTerminal()
		switch input {
		case "q", "quit", "Q", "그만":
			goto ExitLoop
		// 이동
		case "e", "east", "E", "동", "동 가", "동쪽으로 가":
			if canMove(doorGrid, &user.pos, East) {
				user.pos.c++
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "n", "north", "N", "북", "북 가", "북쪽으로 가":
			if canMove(doorGrid, &user.pos, North) {
				user.pos.r--
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "w", "west", "W", "서", "서 가", "서쪽으로 가":
			if canMove(doorGrid, &user.pos, West) {
				user.pos.c--
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "s", "south", "S", "남", "남 가", "남쪽으로 가":
			if canMove(doorGrid, &user.pos, South) {
				user.pos.r++
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "l", "look", "L", "봐":
			showRoomInfo(doorGrid, &itemGrid, &user)
		case "i", "item", "I", "소지품":
			showOwnedItems(user)
		default:
			sInput := strings.Split(input, " ")
			if len(sInput) == 2 {
				if sInput[1] == "봐" { // 입력된 아이템 정보를 본다.
					showItemInfo(user, sInput[0])
				} else if sInput[1] == "열어" { // 입력된 문을 연다.
					door := sInput[0]
					doorDir := getDoorDirection(doorGrid, user, door)
					if doorDir == -1 { // 입력된 문이 현재 방에 없는 경우
						fmt.Println(NoSuchDoor, input)
					} else {

					}
				} else {
					fmt.Print(WrongInput, input)
				}
			} else {
				fmt.Println(WrongInput, input)
			}
		}
	}

ExitLoop:
	clearTerminal()
	fmt.Println("Finished")
}
