package connect4

import (
	"math/rand"
	"time"
)

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
		node := Node{game: Game{isPlayerOneTurn: amIPlayerOne}, playoutCount: 1, winningCount: 1}
		node.selectAndExpand()
		return rand.Intn(8)
	}
}
