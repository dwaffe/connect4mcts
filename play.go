package main

import (
	"fmt"

	. "github.com/dwaffe/connect4mcts/gamelogic"
)

type Game struct {
	board           Board
	movesCounter    int
	isPlayerTwoTurn bool
}

func (g *Game) makeMove(column int) {
	column--
	for i := 0; i < 7; i++ {
		if g.board.PlayerOne[i][column] == 0 && g.board.PlayerTwo[i][column] == 0 {
			if !g.isPlayerTwoTurn {
				g.board.PlayerTwo[i][column] = 1
			} else {
				g.board.PlayerOne[i][column] = 1
			}
			break
		}
	}

	g.isPlayerTwoTurn = !g.isPlayerTwoTurn
}

type Node struct {
	depth                int
	board                Board
	actionCount          int     // N
	value                float32 // W
	meanValueOfNextState float32 // Q
	// probabilityOfSelecting // P
}

func main() {
	fmt.Println("Play connect4!")
	play()
}

func play() {
	game := Game{Board{}, 0, false}
	game.board.Print()
	game.makeMove(2)
	game.makeMove(3)
	game.makeMove(4)
	game.makeMove(2)
	game.makeMove(3)
	game.makeMove(4)
	game.makeMove(7)
	game.board.Print()
	// for !game.isGameOver() {
	// fmt.Println("next move")
	// }
}

func (g *Game) isGameOver() bool {
	// g.board.PlayerOne
	return false
}
