package connect4

type Game struct {
	Board           Board
	movesCounter    int
	isPlayerTwoTurn bool
}

func (g *Game) MakeMove(column int) {
	column--
	for i := 0; i < 7; i++ {
		if g.Board.PlayerOne[i][column] == 0 && g.Board.PlayerTwo[i][column] == 0 {
			if !g.isPlayerTwoTurn {
				g.Board.PlayerTwo[i][column] = 1
			} else {
				g.Board.PlayerOne[i][column] = 1
			}
			break
		}
	}

	g.isPlayerTwoTurn = !g.isPlayerTwoTurn
}

func (g *Game) isGameOver() bool {
	// g.Board.PlayerOne
	return false
}
