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
		t.Errorf("Game should be over:\n%v", g.Board.String())
	}

	if g.winningPlayer != 1 {
		t.Errorf("PlayerOne should won:\n%v", g.Board.String())
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
		t.Errorf("Game should be over:\n%v", g.Board.String())
	}

	if g.winningPlayer != 1 {
		t.Errorf("PlayerOne should won:\n%v", g.Board.String())
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
	g := NewGame()
	g.MakeMove(1)
	g.MakeMove(3)
	g.MakeMove(2)
	g.MakeMove(5)
	g.MakeMove(2)
	g.MakeMove(1)

	if g.IsGameOver() == true {
		t.Errorf("Game should NOT be over:\n%v", g.Board.String())
	}

	if g.winningPlayer != 0 {
		t.Errorf("There is no winner:\n%v", g.Board.String())
	}
}

func TestNotWinningState2(t *testing.T) {
	g := NewGame()
	g.MakeMove(5)
	g.MakeMove(6)
	g.MakeMove(5)
	g.MakeMove(5)
	g.MakeMove(4)

	if g.IsGameOver() == true {
		t.Errorf("Game should NOT be over:\n%v", g.Board.String())
	}

	if g.winningPlayer != 0 {
		t.Errorf("There is no winner:\n%v", g.Board.String())
	}
}
