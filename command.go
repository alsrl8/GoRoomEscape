package main

import "fmt"

// 현재 위치에서 입력된 방향으로 움직일 수 있는지 boolean 반환
func canMove(grid [6][7][4][2]int, pos *Position, dir int) bool {
	switch grid[pos.r][pos.c][dir][0] {
	case Wall:
		return false
	case Clear:
		return true
	case WoodDoor:
		switch grid[pos.r][pos.c][dir][1] {
		case Open:
			return true
		case Closed:
			return false
		default:
			return false
		}
	case GlassDoor:
		return false
	case LockedDoor:
		switch grid[pos.r][pos.c][dir][1] {
		case Open:
			return true
		case Closed:
			return false
		default:
			return false
		}
	default:
		return false
	}
}

func isValidPoistion(grid [6][7][4][2]int, pos Position) bool {
	if pos.r >= len(grid) || pos.r < 0 {
		return false
	} else if pos.c >= len(grid[0]) || pos.c < 0 {
		return false
	} else {
		return true
	}
}

// 현재 위치에서 입력된 방향 기준으로 마주보는 반대쪽 방의 위치와 방향을 계산한다.
func getOppositePositionAndDir(pos Position, dir int) (oPosition Position, oDir int) {
	dRow, dCol := [4]int{0, -1, 0, 1}, [4]int{1, 0, -1, 0}
	oPosition.r = pos.r + dRow[dir]
	oPosition.c = pos.c + dCol[dir]
	oDir = (dir + 2) % 4
	return
}

// 방 위치, 방 안에 있는 아이템, 다음으로 이동 가능한 방향을 출력한다.
func showRoomInfo(doorGrid [6][7][4][2]int, itemGrid *[6][7]int, user *User) {
	pos := user.pos

	fmt.Println("=== 현재 방 정보 ===")
	fmt.Printf("r: %d\tc: %d\n", pos.r, pos.c)
	for d := 0; d < 4; d++ {
		fmt.Printf("%s(%s) - %s\n", DirStringMap[d], DirStringEngMap[d], WallStringMap[doorGrid[pos.r][pos.c][d][0]])
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

// 입력받은 문의 이름이 현재 방의 어느 방향에 있는지 반환한다.
// 해당하는 문이 없으면 -1을 반환한다.
func getDoorDirection(doorGrid [6][7][4][2]int, user User, door string) (d int) {
	wall := StringWallMap[door]
	for ; d < 4; d++ {
		switch doorGrid[user.pos.r][user.pos.c][d][0] {
		case wall:
			return d
		}
	}
	return -1
}

// 입력받은 문의 이름으로 현재 방에서 문을 연다.
func openDoor(doorGrid *[6][7][4][2]int, user User, door string) {
	doorDir := getDoorDirection(*doorGrid, user, door)
	if doorDir == -1 { // 해당하는 문이 현재 방에 없는 경우
		fmt.Println(NoSuchDoor, door)
		return
	}

	switch StringWallMap[door] {
	case WoodDoor:
		if doorGrid[user.pos.r][user.pos.c][doorDir][1] == Open {
			fmt.Println(AlreadyOpenDoor, door)
		} else {
			doorGrid[user.pos.r][user.pos.c][doorDir][1] = Open
			oppositePos, oppositeDir := getOppositePositionAndDir(user.pos, doorDir)
			if isValidPoistion(*doorGrid, oppositePos) {
				doorGrid[oppositePos.r][oppositePos.c][oppositeDir][1] = Open
			}
			fmt.Println(SucceedOpeningDoor, door)
		}
	case GlassDoor:
		if user.items[Hammer] == 0 { // 유리문을 부술 망치가 없는 경우
			fmt.Println(NotEnoughItemsToOpenDoor, Hammer, "현재 개수: ", user.items[Hammer])
		} else {
			doorGrid[user.pos.r][user.pos.c][doorDir][0] = Clear
			doorGrid[user.pos.r][user.pos.c][doorDir][1] = Open
			oppositePos, oppositeDir := getOppositePositionAndDir(user.pos, doorDir)
			if isValidPoistion(*doorGrid, oppositePos) {
				doorGrid[oppositePos.r][oppositePos.c][oppositeDir][0] = Clear
				doorGrid[oppositePos.r][oppositePos.c][oppositeDir][1] = Open
			}
			fmt.Println(SucceedOpeningDoor, door)
		}
	case LockedDoor:
		if user.items[Key] == 0 { // 잠긴 문을 열 열쇠가 없는 경우
			fmt.Println(NotEnoughItemsToOpenDoor, Key, "현재 개수: ", user.items[Key])
		} else {
			doorGrid[user.pos.r][user.pos.c][doorDir][0] = WoodDoor
			doorGrid[user.pos.r][user.pos.c][doorDir][1] = Open
			oppositePos, oppositeDir := getOppositePositionAndDir(user.pos, doorDir)
			if isValidPoistion(*doorGrid, oppositePos) {
				doorGrid[oppositePos.r][oppositePos.c][oppositeDir][0] = WoodDoor
				doorGrid[oppositePos.r][oppositePos.c][oppositeDir][1] = Open
			}
			fmt.Println(SucceedOpeningDoor, door)
		}
	default:
		fmt.Println(CanNotOpenSuchDoor, door)
	}
}
