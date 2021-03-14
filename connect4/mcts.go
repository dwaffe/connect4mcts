package connect4

import (
	"fmt"
	"math"
)

const exploreCoefficients = 1.414
const maxDepth = 4

type Node struct {
	// board        Board
	game         Game
	playoutCount int
	winningCount int
	legalMoves   [7]*Node
	parentNode   *Node
	me           int

	// actionCount          int     // N
	// value                float32 // W
	// meanValueOfNextState float32 // Q
	// probabilityOfSelecting // P

}

func (n *Node) selectAndExpand(currentDepth int) {
	n.createLegalNodes()
	var calculatedUCT [7]float64

	for i, childNode := range n.legalMoves {
		if childNode == nil {
			calculatedUCT[i] = -100.0
			continue
		}
		calculatedUCT[i] = childNode.calculateUpperConfidenceBound()
	}
	// fmt.Println(calculatedUCT)
	// fmt.Println(n.legalMoves)
	if currentDepth == maxDepth {
		n.legalMoves[getHighestValue(calculatedUCT)].playout()
	} else {
		n.legalMoves[getHighestValue(calculatedUCT)].selectAndExpand(currentDepth + 1)
	}
}

func getHighestValue(array [7]float64) int {
	highestValue := -99999999.9
	var highestIndex int
	for i := 0; i < len(array); i++ {
		if array[i] > highestValue {
			highestIndex = i
			highestValue = array[i]
		}
	}

	return highestIndex
}

func (n *Node) getBestMove() int {

	for i := 0; i < 10000; i++ {
		n.selectAndExpand(0)
	}

	var playoutCount [7]float64

	for i, childNode := range n.legalMoves {
		if childNode == nil {
			playoutCount[i] = 0
			continue
		}
		playoutCount[i] = float64(childNode.playoutCount)
	}

	fmt.Println("last:", playoutCount)
	return getHighestValue(playoutCount) + 1
}

func (n *Node) createLegalNodes() {
	for i := 0; i < boardColumns; i++ {
		if n.legalMoves[i] != nil {
			continue
		}
		if n.game.Board.IsIllegalMove(i + 1) {
			continue
		}
		copyOfGame := n.game
		copyOfGame.MakeMove(i)
		n.legalMoves[i] = &Node{game: copyOfGame, playoutCount: 1, winningCount: 0, me: n.me}
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

func (n *Node) playout() {
	tempGame := n.game
	tempGame.Play(GetEasyPlayer(), GetEasyPlayer(), true)
	switch tempGame.winningPlayer {
	case 0:
		n.update(0.5)
	case n.me:
		n.update(1)
	default:
		n.update(0)
	}
}

func (n *Node) update(value float64) {
	n.playoutCount++
	if value == 1 {
		n.winningCount++
	}

	if value == 0 {
		n.winningCount--
	}

	if n.parentNode != nil {
		n.parentNode.update(value)
	}
}
