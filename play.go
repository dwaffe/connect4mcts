package main

import "fmt"

const boardRows int = 6
const boardColumns int = 7

type Board struct {
	playerOne [boardRows][boardColumns]int
	playerTwo [boardRows][boardColumns]int
}

type Game struct {
	board         Board
	playerTwoTurn bool
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
	game := Game{Board{}, true}
	for !game.isGameOver() {
		fmt.Println("next move")
	}
}

func (g *Game) isGameOver() bool {
	// g.board.playerOne
	return false
}
