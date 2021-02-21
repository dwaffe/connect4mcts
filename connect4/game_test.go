package connect4

import "testing"

func TestHorizontalWinningState(t *testing.T) {
	board := Board{
		PlayerOne: [6][7]int{
			{1, 0, 1, 1, 0, 1, 0},
		},
		PlayerTwo: [6][7]int{},
	}

	g := NewGame()
	g.Board = board
	g.MakeMove(2)
	g.Board.Print()

	if g.IsGameOver() != true {
		t.Errorf("Game should be over")
	}

	if g.winningPlayer != 1 {
		t.Errorf("PlayerOne should won")
	}
}

func TestVerticalWinningState(t *testing.T) {
	g := NewGame()
	g.MakeMove(1)
	g.MakeMove(3)
	g.MakeMove(2)
	g.MakeMove(3)
	g.MakeMove(2)
	g.MakeMove(3)
	g.MakeMove(2)
	g.MakeMove(3)

	g.Board.Print()

	if g.IsGameOver() != true {
		t.Errorf("Game should be over")
	}

	if g.winningPlayer != 2 {
		t.Errorf("PlayerTwo should won")
	}
}

func TestDiagonalWinningState(t *testing.T) {
	board := Board{
		PlayerOne: [6][7]int{
			{1, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		PlayerTwo: [6][7]int{
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
		},
	}

	g := NewGame()
	g.Board = board

	g.MakeMove(4)

	if g.IsGameOver() != true {
		t.Errorf("Game should be over: %v", g.Board)
	}

	if g.winningPlayer != 1 {
		t.Errorf("PlayerOne should won: %v", g.Board)
	}
}

func TestMirrorDiagonalWinningState(t *testing.T) {
	board := Board{
		PlayerOne: [6][7]int{
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		PlayerTwo: [6][7]int{
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
		},
	}

	g := NewGame()
	g.Board = board

	g.MakeMove(4)

	if g.IsGameOver() != true {
		t.Errorf("Game should be over: %v", g.Board.String())
	}

	if g.winningPlayer != 1 {
		t.Errorf("PlayerOne should won: %v", g.Board.String())
	}
}

func TestMirrorDiagonalWinningState2(t *testing.T) {
	board := Board{
		PlayerOne: [6][7]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 0, 1},
		},
		PlayerTwo: [6][7]int{
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}

	g := NewGame()
	g.Board = board

	g.MakeMove(4)

	if g.IsGameOver() != true {
		t.Errorf("Game should be over:\n%v", g.Board.String())
	}

	if g.winningPlayer != 1 {
		t.Errorf("PlayerOne should won:\n%v", g.Board.String())
	}
}

func TestNotWinningState(t *testing.T) {

}