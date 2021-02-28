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
	winningPlayer   int
}

type lastMove struct {
	column          int
	row             int
	isPlayerOneMove bool
}

func NewGame() *Game {
	return &Game{isPlayerOneTurn: true}
}

func (g *Game) Play(PlayerOne, PlayerTwo func(Board, bool) int) {
	var player func(Board, bool) int
	for !g.IsGameOver() {
		if g.isPlayerOneTurn {
			player = PlayerOne
		} else {
			player = PlayerTwo
		}
		g.MakeMove(player(g.Board, g.isPlayerOneTurn))
		g.Board.Print()
	}

	if g.winningPlayer == 0 {
		fmt.Println("It's a draw")
	} else {
		fmt.Printf("Player %v has won!\n", g.winningPlayer)
	}

}

func (g *Game) MakeMove(column int) (err error) {
	if g.Board.IsIllegalMove(column) {
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
	diagonalWinCounter := 0
	mirrorDiagonalWinCounter := 0

	for column, row := minColumn, minRow; column <= maxColumn; column++ {
		if playerBoard[lastMove.row][column] == 1 {
			horizontalWinCounter++
		} else {
			horizontalWinCounter = 0
		}

		if horizontalWinCounter == 4 {
			g.setWinningState(lastMove)
			break
		}

		// diagonal:
		tempRow := lastMove.row - (3 - row)
		tempColumn := lastMove.column + (3 - column)
		if tempRow <= maxRow && tempRow >= 0 && tempColumn >= 0 && tempColumn <= maxColumn && playerBoard[tempRow][tempColumn] == 1 {
			diagonalWinCounter++
		} else {
			diagonalWinCounter = 0
		}

		if diagonalWinCounter == 4 {
			g.setWinningState(lastMove)
			break
		}

		// mirror diagonal:
		tempRow = lastMove.row - 3 + row
		tempColumn = lastMove.column - 3 + row
		if tempRow <= maxRow && tempRow >= 0 && tempColumn >= 0 && tempColumn <= maxColumn && playerBoard[tempRow][tempColumn] == 1 {
			mirrorDiagonalWinCounter++
		} else {
			mirrorDiagonalWinCounter = 0
		}

		if mirrorDiagonalWinCounter == 4 {
			g.setWinningState(lastMove)
			break
		}

		// vertical:
		if row > maxRow {
			continue
		}

		if playerBoard[row][lastMove.column] == 1 {
			verticalWinCounter++
		} else {
			verticalWinCounter = 0
		}

		if verticalWinCounter == 4 {
			g.setWinningState(lastMove)
			break
		}

		row++
	}

	if g.movesCounter >= 42 {
		g.isGameOver = true
	}
}

func (g *Game) setWinningState(lastMove lastMove) {
	g.isGameOver = true
	if lastMove.isPlayerOneMove {
		g.winningPlayer = 1
	} else {
		g.winningPlayer = 2
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
