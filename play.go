package main

import "fmt"

const boardRows int = 6
const boardColumns int = 7
const playerOneSign = "X"
const playerTwoSign = "O"

type Board struct {
	playerOne [boardRows][boardColumns]int
	playerTwo [boardRows][boardColumns]int
}

func (b *Board) print() {
	for x := len(b.playerOne) - 1; x >= 0; x-- {
		for y, v := range b.playerOne[x] {
			player := " "
			if v == 1 {
				player = playerOneSign
			} else if b.playerTwo[x][y] == 1 {
				player = playerTwoSign
			} else {

			}
			fmt.Printf("[%v]", player)

			// println("ONE: [%v, %v]: %v, ", x, y, v)
			// println("TWO: [%v, %v]: %v, ", x, y, b.playerOne[x][y])
		}
		println()
	}
	println(" 1  2  3  4  5  6  7")
}

type Game struct {
	board           Board
	movesCounter    int
	isPlayerTwoTurn bool
}

func (g *Game) makeMove(column int) {
	column--
	for i := 0; i < 7; i++ {
		if g.board.playerOne[i][column] == 0 && g.board.playerTwo[i][column] == 0 {
			if !g.isPlayerTwoTurn {
				g.board.playerTwo[i][column] = 1
			} else {
				g.board.playerOne[i][column] = 1
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
	game.board.print()
	game.makeMove(2)
	game.makeMove(3)
	game.makeMove(4)
	game.makeMove(2)
	game.makeMove(3)
	game.makeMove(4)
	game.board.print()
	// for !game.isGameOver() {
	// fmt.Println("next move")
	// }
}

func (g *Game) isGameOver() bool {
	// g.board.playerOne
	return false
}
