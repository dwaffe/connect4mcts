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

func GetHumanPlayer() func(game Game, amIPlayerOne bool) int {
	return func(game Game, amIPlayerOne bool) int {

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

func GetEasyPlayer() func(game Game, amIPlayerOne bool) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return func(game Game, amIPlayerOne bool) int {
		var me int
		if amIPlayerOne == true {
			me = 1
		} else {
			me = 2
		}
		for i := 1; i < boardColumns; i++ {
			game := Game{Board: game.Board, movesCounter: game.movesCounter, isPlayerOneTurn: amIPlayerOne}
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

func GetMediumPlayer() func(game Game, amIPlayerOne bool) int {

	return func(game Game, amIPlayerOne bool) int {
		var me int
		if amIPlayerOne {
			me = 1
		} else {
			me = 2
		}

		game.isPlayerOneTurn = amIPlayerOne
		node := Node{game: game, playoutCount: 1, winningCount: 1, me: me}
		for i := 0; i < 10000; i++ {
			node.selectAndExpand()
		}

		return node.getBestMove() + 1
	}
}
