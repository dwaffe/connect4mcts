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

func parseToLegalMoveOrReturn0(x rune) int {
	i := int(x - '0')
	if i < 1 || i > 7 {
		return 0
	}

	return i
}

func play() {
	game := Game{Board{}}
	game.Board.Print()
	game.MakeMove(2)
	game.MakeMove(3)
	game.MakeMove(4)
	game.MakeMove(2)
	game.MakeMove(3)
	game.MakeMove(4)
	game.MakeMove(7)
	game.Board.Print()

	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("You pressed: %d\r\n", parseToLegalMoveOrReturn0(char))

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

	// for !game.isGameOver() {
	// fmt.Println("next move")
	// }
}
