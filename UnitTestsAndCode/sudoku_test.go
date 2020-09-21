package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplacesValueInInitialSudokuWithPlayerinput(t *testing.T) {
	type arguments struct {
		playerInput int
	}
	tests := []struct {
		name      string
		arguments arguments
		want      int
	}{
			{"Replaces value in original Sudoku if player's input matches appropriate position in solved sudoku", arguments{4}, 4},

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ReplacesValueInInitialSudokuWithPlayerinput(test.arguments.playerInput)
			assert.Equal(t, test.want, got)
		})
	}

}



func TestNoZeroValueLeftInOriginalSudoku(t *testing.T)  {
	type arguments struct {
		bool
	}
		tests := []struct {
			name      string
			arguments arguments
			want      bool
		}{
			{"Check no zero are left in original sudoku", arguments{true}, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := NoZeroValueLeftInOriginalSudoku(test.arguments.bool)
				assert.Equal(t, test.want, got)
			})
		}
}