package main

import (
	"fmt"

	. "github.com/dwaffe/connect4mcts/connect4"
	"github.com/eiannone/keyboard"
)

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

func parseToInt(x rune) int {
	return int(x - '0')
}

func play() {
	game := NewGame()
	game.Board.Print()

	for !game.IsGameOver() {
		fmt.Println("Your move:")
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		err = game.MakeMove(parseToInt(char))
		if err != nil {
			fmt.Print(err)
			fmt.Println(" Press ESC to exit")

		} else {
			game.Board.Print()
		}

	}
}
