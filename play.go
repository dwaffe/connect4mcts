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
	game := Game{}
	game.Board.Print()

	for !game.IsGameOver() {

		fmt.Println("Your move:")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		game.MakeMove(parseToInt(char))
		game.Board.Print()
	}

	// reader := bufio.NewReader(os.Stdin)
	// char, _, err := reader.ReadRune()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// switch char {
	// case '1':
	// 	fmt.Println("A Key '1'")
	// 	break
	// case '2':
	// 	fmt.Println("a Key 2")
	// 	break
	// }

}
