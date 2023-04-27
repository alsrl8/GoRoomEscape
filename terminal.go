package main

import (
	"bufio"
	"fmt"
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

// Terminal에서 사용자 입력을 받는다.(Windows)
func getInput() string {
	fmt.Printf("입력 >>> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\r')
	input = strings.TrimSuffix(input, "\r")
	return input
}

func initTerminal(doorGrid [6][7][4][2]int, itemGrid [6][7]int, user User) {
	clearTerminal()
	for {
		showRoomInfo(doorGrid, &itemGrid, &user)
		input := getInput()
		clearTerminal()
		switch input {
		case "q", "quit", "Q", "그만":
			goto QuitLoop
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
			// 아이템
			reg, _ := regexp.Compile(" 봐$")
			if reg.MatchString(input) {
				item := strings.TrimSuffix(input, " 봐")
				showItemInfo(user, item)
			} else {
				// 열기
				reg, _ = regexp.Compile("( 열기| 열어| 열)$")
				if reg.MatchString(input) {
					door := reg.ReplaceAllString(input, "")
					openDoor(&doorGrid, user, door)
				} else {
					// 닫기
					reg, _ = regexp.Compile("( 닫기| 닫어| 닫)$")
					if reg.MatchString(input) {
						door := reg.ReplaceAllString(input, "")
						closeDoor(&doorGrid, user, door)
					} else {
						fmt.Println(WrongInput, input)
					}
				}
			}
		}

		// 탈출 조건: 현재 위치가 grid의 바깥인 경우
		if !isValidPoistion(doorGrid, user.pos) {
			goto ExitLoop
		}
	}

ExitLoop:
	clearTerminal()
	fmt.Println("탈출했습니다.")
	return

QuitLoop:
	clearTerminal()
	fmt.Println("종료했습니다.")
}
