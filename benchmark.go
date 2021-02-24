package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/dwaffe/connect4mcts/connect4"
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
	start := time.Now()
	seconds := 3
	playout_count := 0

	for {
		play()
		elapsed_time := time.Since(start)
		playout_count++
		if int(elapsed_time.Seconds()) >= seconds {
			fmt.Println("Finished iterating in: ", elapsed_time, " seconds")
			fmt.Println("Played games: ", playout_count)
			break
		}
	}
}

func parseToInt(x rune) int {
	return int(x - '0')
}

func play() {
	game := NewGame()

	for !game.IsGameOver() {
		game.MakeMove(rand.Intn(8))
		// game.Board.Print()
	}
}
