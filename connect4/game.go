package connect4

import "errors"

type Game struct {
	Board           Board
	movesCounter    int
	isPlayerTwoTurn bool
}

func isIllegalMove(move int) bool {
	return move < 1 || move > boardColumns
}

func (g *Game) MakeMove(column int) (err error) {
	if isIllegalMove(column) {
		err = errors.New("Illegal move.")
		return
	}

	column--
	for i := 0; i < 7; i++ {
		if g.Board.PlayerOne[i][column] == 0 && g.Board.PlayerTwo[i][column] == 0 {
			if !g.isPlayerTwoTurn {
				g.Board.PlayerTwo[i][column] = 1
			} else {
				g.Board.PlayerOne[i][column] = 1
			}

			g.isPlayerTwoTurn = !g.isPlayerTwoTurn
			break
		}
	}

	return
}

func (g *Game) IsGameOver() bool {
	// g.Board.PlayerOne
	return false
}
