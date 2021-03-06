package connect4

import "fmt"

const boardRows int = 6
const boardColumns int = 7
const playerOneSign = "X"
const playerTwoSign = "O"

// Board to play connect4
type Board struct {
	PlayerOne [boardRows][boardColumns]int
	PlayerTwo [boardRows][boardColumns]int
}

func (b *Board) Print() {
	println(b.String())
}

func (b *Board) IsEmpty(row, column int) bool {
	return b.PlayerOne[row][column] == 0 && b.PlayerTwo[row][column] == 0
}

func (b *Board) String() (result string) {
	for x := len(b.PlayerOne) - 1; x >= 0; x-- {
		for y, v := range b.PlayerOne[x] {
			player := " "
			if v == 1 {
				player = playerOneSign
			} else if b.PlayerTwo[x][y] == 1 {
				player = playerTwoSign
			} else {

			}
			result += fmt.Sprintf("[%v]", player)
		}
		result += "\n"
	}
	result += " 1  2  3  4  5  6  7"

	return
}

func (b *Board) IsIllegalMove(move int) bool {
	if move < 1 || move > boardColumns {
		return true
	}

	return !b.IsEmpty(boardRows-1, move-1)
}
