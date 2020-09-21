package app

import (
	"fmt"
)

var initialSudoku = [9][9]int{

{5,3,0,0,7,0,0,0,0},
{6,0,0,1,9,5,0,0,0},
{0,9,8,0,0,0,0,6,0},
{8,0,0,0,6,0,0,0,3},
{4,0,0,8,0,3,0,0,1},
{7,0,0,0,2,0,0,0,6},
{0,6,0,0,0,0,2,8,0},
{0,0,0,4,1,9,0,0,5},
{0,0,0,0,8,0,0,7,9},
}

var solvedSudoku =  [9][9]int{
{5,3,4,6,7,8,9,1,2},
{6,7,2,1,9,5,3,4,8},
{1,9,8,3,4,2,5,6,7},
{8,5,9,7,6,1,4,2,3},
{4,2,6,8,5,3,7,9,1},
{7,1,3,9,2,4,8,5,6},
{9,6,1,5,3,7,2,8,4},
{2,8,7,4,1,9,6,3,5},
{3,4,5,2,8,6,1,7,9},
}

func ReplacesValueInInitialSudokuWithPlayerinput(playerInput int) int {

	var rowIndex int = 0
	var columnIndex int = 2
	playerInput = 4

	if !ReturnsTrueIfPlayerInputIs1to9(playerInput) {	
		fmt.Printf ("Only numbers from 1 to 9 are allowed")
		return 0
	}
	if playerInput != (solvedSudoku [rowIndex][columnIndex]) {
		return 0
	}
	initialSudoku[rowIndex][columnIndex] = playerInput
	for _, v := range initialSudoku {
		fmt.Printf("%d\n", v)
		}
	return (initialSudoku[rowIndex][columnIndex])
}

func ReturnsTrueIfPlayerInputIs1to9(playerInput int) bool {
	var number = playerInput;
	
    if number < 1 || number > 9 {
        return false
    }
    return true
}

