package main

import (
	"fmt"
	"time"

	. "github.com/dwaffe/connect4mcts/connect4"
	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println("Play connect4!")
	// benchmark()
	botPlay()
	// play()
}

func parseToInt(x rune) int {
	return int(x - '0')
}

func botPlay() {
	game := NewGame()
	game.Play(GetEasyPlayer(), GetEasyPlayer())
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

func benchmark() {
	start := time.Now()
	seconds := 3
	playout_count := 0

	for {
		botPlay()
		elapsed_time := time.Since(start)
		playout_count++
		if int(elapsed_time.Seconds()) >= seconds {
			fmt.Println("Finished iterating in: ", elapsed_time, " seconds")
			fmt.Println("Played games: ", playout_count)
			break
		}
	}
}
