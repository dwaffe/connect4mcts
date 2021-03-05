package connect4

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

type Player interface {
	getMove(board Board, amIPlayerOne bool) int
}

func GetHumanPlayer() func(board Board, amIPlayerOne bool) int {
	return func(board Board, amIPlayerOne bool) int {

		fmt.Println("Your move:")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		return parseToInt(char)
	}
}

func parseToInt(x rune) int {
	return int(x - '0')
}

func GetEasyPlayer() func(board Board, amIPlayerOne bool) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return func(board Board, amIPlayerOne bool) int {
		var me int
		if amIPlayerOne == true {
			me = 1
		} else {
			me = 2
		}
		for i := 1; i < boardColumns; i++ {
			game := Game{Board: board, movesCounter: 8, isPlayerOneTurn: amIPlayerOne}
			if game.Board.IsIllegalMove(i) {
				continue
			}

			game.MakeMove(i)
			if game.winningPlayer == me {
				return i
			}
		}
		return rand.Intn(8)
	}
}

func GetMediumPlayer() func(board Board, amIPlayerOne bool) int {

	return func(board Board, amIPlayerOne bool) int {
		var me int
		if amIPlayerOne {
			me = 1
		} else {
			me = 2
		}
		node := Node{game: Game{isPlayerOneTurn: amIPlayerOne}, playoutCount: 1, winningCount: 1, me: me}
		for i := 0; i < 5000; i++ {
			node.selectAndExpand()
		}

		return node.getBestMove() + 1
	}
}
