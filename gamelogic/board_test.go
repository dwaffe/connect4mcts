package gamelogic

import "testing"

func TestEmptyBoardToString(t *testing.T) {
	b := Board{
		PlayerOne: [6][7]int{},
		PlayerTwo: [6][7]int{},
	}
	want := `[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
 1  2  3  4  5  6  7`
	if got := b.String(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestFilledBoardToString(t *testing.T) {
	b := Board{
		PlayerOne: [6][7]int{
			{1, 0, 0, 0, 0, 0, 1},
		},
		PlayerTwo: [6][7]int{
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
		},
	}
	want := `[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ]
[ ][ ][O][ ][ ][ ][ ]
[X][ ][O][ ][ ][ ][X]
 1  2  3  4  5  6  7`
	if got := b.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
