package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/dwaffe/connect4mcts/connect4"
)

func main() {
	start := time.Now()
	seconds := 3
	playout_count := 0

	for {
		playOneRandomGame()
		elapsed_time := time.Since(start)
		playout_count++
		if int(elapsed_time.Seconds()) >= seconds {
			fmt.Println("Finished iterating in: ", elapsed_time, " seconds")
			fmt.Println("Played games: ", playout_count)
			break
		}
	}
}

func playOneRandomGame() {
	game := NewGame()

	for !game.IsGameOver() {
		game.MakeMove(rand.Intn(8))
		// game.Board.Print()
	}
}
