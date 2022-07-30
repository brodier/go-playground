package sudoku

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestSqId(t *testing.T) {
	want := []int{1, 1, 1, 2, 2, 2, 3, 3, 3,
		1, 1, 1, 2, 2, 2, 3, 3, 3,
		1, 1, 1, 2, 2, 2, 3, 3, 3,
		4, 4, 4, 5, 5, 5, 6, 6, 6,
		4, 4, 4, 5, 5, 5, 6, 6, 6,
		4, 4, 4, 5, 5, 5, 6, 6, 6,
		7, 7, 7, 8, 8, 8, 9, 9, 9,
		7, 7, 7, 8, 8, 8, 9, 9, 9,
		7, 7, 7, 8, 8, 8, 9, 9, 9}
	got := make([]int, 81)

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			got[x%9+y*9] = SqId(x, y) + 1
		}
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Hello() = %v, want %v", got, want)
	}
}

func TestIsValid(t *testing.T) {

	cells := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		9, 1, 2, 3, 4, 5, 6, 7, 8,
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		6, 7, 8, 9, 1, 2, 3, 4, 5,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		2, 3, 4, 5, 6, 7, 8, 9, 1,
		5, 6, 7, 8, 9, 1, 2, 3, 4}

	if !isValid(cells) {
		t.Errorf("IsValid should be true for this test")
	}
}

func TestSolver(t *testing.T) {
	cells := []int{1, 2, 0, 0, 7, 0, 5, 6, 0,
		5, 0, 7, 9, 3, 2, 0, 8, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 1, 0, 2, 4, 0, 0, 5, 0,
		3, 0, 8, 0, 0, 0, 4, 0, 2,
		0, 7, 0, 0, 8, 5, 0, 1, 0,
		0, 0, 0, 7, 0, 0, 0, 0, 0,
		0, 8, 0, 4, 2, 3, 7, 0, 1,
		0, 3, 4, 0, 1, 0, 0, 2, 8}
	want := []int{1, 2, 3, 8, 7, 4, 5, 6, 9,
		5, 6, 7, 9, 3, 2, 1, 8, 4,
		8, 4, 9, 6, 5, 1, 2, 3, 7,
		9, 1, 6, 2, 4, 7, 8, 5, 3,
		3, 5, 8, 1, 9, 6, 4, 7, 2,
		4, 7, 2, 3, 8, 5, 9, 1, 6,
		2, 9, 1, 7, 6, 8, 3, 4, 5,
		6, 8, 5, 4, 2, 3, 7, 9, 1,
		7, 3, 4, 5, 1, 9, 6, 2, 8}
	solution := NewSudoku(want)
	log.Printf("Init Sudoku : %v", cells)
	toSolve := NewSudoku(cells)
	if !toSolve.Propagate() {
		t.Errorf("Invalid starting position %v", cells)
	}
	got, found := toSolve.Solve()
	if !found {
		t.Errorf("Solver not found solution")
	}
	log.Printf("Solution : %v", solution)
	if fmt.Sprintf("%v", got) == fmt.Sprintf("%v", solution) {
		t.Errorf("Hello() = %v, want %v", got, solution)
	}

}
