package connect4

import (
	"errors"
	"fmt"
)

type Game struct {
	Board           Board
	movesCounter    int
	isPlayerOneTurn bool
	isGameOver      bool
}

type lastMove struct {
	column          int
	row             int
	isPlayerOneMove bool
}

func (g *Game) isIllegalMove(move int) bool {
	if move < 1 || move > boardColumns {
		return true
	}

	return !g.Board.IsEmpty(boardRows-1, move-1)
}

func (g *Game) MakeMove(column int) (err error) {
	if g.isIllegalMove(column) {
		err = errors.New("Illegal move.")
		return
	}

	var lastPlayerMove lastMove

	column--
	for i := 0; i < 7; i++ {
		if g.Board.IsEmpty(i, column) {
			if g.isPlayerOneTurn {
				g.Board.PlayerOne[i][column] = 1
			} else {
				g.Board.PlayerTwo[i][column] = 1
			}

			lastPlayerMove = lastMove{column, i, g.isPlayerOneTurn}
			g.isPlayerOneTurn = !g.isPlayerOneTurn
			break
		}
	}

	g.updateState(lastPlayerMove)

	return
}

func (g *Game) IsGameOver() bool {
	return g.isGameOver
}

func (g *Game) updateState(lastMove lastMove) {
	g.movesCounter++
	fmt.Println(lastMove)

	minColumn := max(lastMove.column-3, 0)
	maxColumn := min(lastMove.column+3, boardColumns-1)

	minRow := max(lastMove.row-3, 0)
	maxRow := min(lastMove.row+3, boardRows-1)

	var playerBoard [boardRows][boardColumns]int

	if lastMove.isPlayerOneMove {
		playerBoard = g.Board.PlayerOne
	} else {
		playerBoard = g.Board.PlayerTwo
	}

	horizontalWinCounter := 0
	verticalWinCounter := 0
	row := minRow
	for column := minColumn; column <= maxColumn; column++ {
		// fmt.Println("check:", column, lastMove.row, " is:", playerBoard[lastMove.row][column])
		if playerBoard[lastMove.row][column] == 1 {
			horizontalWinCounter++
		} else {
			horizontalWinCounter = 0
		}

		if horizontalWinCounter == 4 {
			g.isGameOver = true
			break
		}

		if row > maxRow {
			continue
		}

		if playerBoard[row][lastMove.column] == 1 {
			verticalWinCounter++
		} else {
			verticalWinCounter = 0
		}

		if verticalWinCounter == 4 {
			g.isGameOver = true
			break
		}
		row++
	}

	if g.movesCounter >= 42 {
		g.isGameOver = true
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
