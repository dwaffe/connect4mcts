package connect4

import (
	"math/rand"
	"time"
)

type Node struct {
	depth                int
	board                Board
	actionCount          int     // N
	value                float32 // W
	meanValueOfNextState float32 // Q
	// probabilityOfSelecting // P
}

type Player interface {
	getMove(board Board, amIPlayerOne bool) int
}

func GetEasyPlayer() func(board Board, amIPlayerOne bool) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return func(board Board, amIPlayerOne bool) int {
		return rand.Intn(8)
	}
}

func GetMediumPlayer() func(board Board, amIPlayerOne bool) int {

	return func(board Board, amIPlayerOne bool) int {

		return rand.Intn(8)
	}
}
