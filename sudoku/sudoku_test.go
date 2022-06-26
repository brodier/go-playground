package sudoku

import "testing"
import "reflect"

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
	if !reflect.DeepEqual(got,want) {
		t.Errorf("Hello() = %v, want %v", got, want)
	}
}


func TestIsValid(t *testing.T) {

	cells := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9,
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
