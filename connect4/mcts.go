package connect4

import (
	"fmt"
	"math"
)

const exploreCoefficients = 1.414

type Node struct {
	// board        Board
	game         Game
	playoutCount int
	winningCount int
	legalMoves   []Node
	parentNode   *Node

	// actionCount          int     // N
	// value                float32 // W
	// meanValueOfNextState float32 // Q
	// probabilityOfSelecting // P

}

func (n *Node) selectAndExpand() {
	if len(n.legalMoves) == 0 {
		n.createLegalNodes()
	}
	var calculatedUCT [7]float64

	for i, childNode := range n.legalMoves {
		calculatedUCT[i] = childNode.calculateUpperConfidenceBound()
	}

	fmt.Println(calculatedUCT)
}

func (n *Node) createLegalNodes() {
	for i := 1; i <= 7; i++ {
		if n.game.Board.IsIllegalMove(i) {
			continue
		}
		copyOfGame := n.game
		copyOfGame.MakeMove(i)
		n.legalMoves = append(n.legalMoves, Node{game: copyOfGame})
	}
}

func (n *Node) calculateUpperConfidenceBound() float64 {
	var parentPlayoutCount float64
	if n.parentNode == nil {
		parentPlayoutCount = float64(n.playoutCount)
	} else {
		parentPlayoutCount = float64(n.parentNode.playoutCount)
	}

	return float64(n.winningCount)/float64(n.playoutCount) + (exploreCoefficients * float64(math.Sqrt(float64((math.Log(parentPlayoutCount))/float64(n.playoutCount)))))
}
