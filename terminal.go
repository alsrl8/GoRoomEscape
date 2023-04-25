package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Position struct {
	r int
	c int
	d int
}

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showRoomInfo(pos *Position) {
	fmt.Println("=== 현재 방 정보 ===")
	fmt.Printf("r: %d\tc: %d\td: %d\n", pos.r, pos.c, pos.d)
	for d := 0; d < 4; d++ {
		fmt.Printf("%s - \n", DirStringMap[d])

	}
	fmt.Println("====================")
}

func getInput(input *string) {
	fmt.Printf("입력 >>> ")
	fmt.Scanf("%s", input)
}

func gameStart() {
	var input string
	stdin := bufio.NewReader(os.Stdin)

	doorGrid := initDoorGrid()
	fmt.Println(doorGrid)
	// itemGrid := initItemGrid()
	position := initPosition()
	// clearTerminal()

	for {
		showRoomInfo(&position)
		getInput(&input)
		clearTerminal()
		switch input {
		case "q":
			goto ExitLoop
		// 이동
		case "동", "동 가", "동쪽으로 가":
			position.d = East
			if canMove(doorGrid, &position) {
				position.c += 1
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "북", "북 가", "북쪽으로 가":
			position.d = North
			if canMove(doorGrid, &position) {
				position.r -= 1
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "서", "서 가", "서쪽으로 가":
			position.d = West
			if canMove(doorGrid, &position) {
				position.c -= 1
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		case "남", "남 가", "남쪽으로 가":
			position.d = South
			if canMove(doorGrid, &position) {
				position.r += 1
			} else {
				fmt.Println(CanNotMoveWarning)
			}
		default:
			continue
		}
		stdin.ReadString('\n')
	}

ExitLoop:
	clearTerminal()
	fmt.Println("Finished")
}
