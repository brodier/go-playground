package sudoku

import (
	"fmt"
	"log"
)

type Sudoku []int
type Choices map[int]Set

var (
	solvedPos      = initSolvedPos()
	solvedPosDigit = initSolvedPosDigit()
	linkedCellMap  = initLinkedCells()
	sqOffset       = []int{0, 1, 2, 9, 10, 11, 18, 19, 20}
	undefiedPos    = 1<<sudokuSize - 1
)

func pos(x, y int) int { return x + y*9 }
func y(pos int) int    { return pos / 9 }
func x(pos int) int    { return pos % 9 }
func s0(pos int) int {
	var sqId = SqId(x(pos), y(pos))
	return (sqId%3)*3 + (sqId/3)*27
}

func initLinkedCells() map[int]Set {
	var linkedCellsMap = map[int]Set{}
	for pos := 0; pos < sudokuSize*sudokuSize; pos++ {
		var linkedCells = Set{}
		x0 := y(pos) * 9
		y0 := x(pos)
		s0 := s0(pos)
		for r := x0; r < x0+sudokuSize; r++ {
			linkedCells[r] = true
		}
		for c := y0; c < sudokuSize*sudokuSize; c += sudokuSize {
			linkedCells[c] = true
		}
		for _, soff := range sqOffset {
			linkedCells[s0+soff] = true
		}
		delete(linkedCells, pos)
		linkedCellsMap[pos] = linkedCells
	}
	return linkedCellsMap
}

func initSolvedPos() Set {
	var sp = Set{}
	for i := 0; i < sudokuSize; i++ {
		sp[1<<i] = true
	}
	log.Printf("Solved Pos value : %v", sp)
	return sp
}

func initSolvedPosDigit() map[int]string {
	var sp = make(map[int]string, sudokuSize)
	for i := 0; i < sudokuSize; i++ {
		sp[1<<i] = fmt.Sprintf("%d", i+1)
	}
	return sp
}

func (s *Sudoku) propagate(posToPropagate Set) bool {
	for len(posToPropagate) > 0 {
		var pos int
		for posItt := range posToPropagate {
			pos = posItt
			delete(posToPropagate, posItt)
			break
		}
		var value = (*s)[pos]
		for lp := range linkedCellMap[pos] {
			if (*s)[lp] == value {
				return false
			}
			if !solvedPos[(*s)[lp]] {
				(*s)[lp] &= value ^ undefiedPos
				if solvedPos[(*s)[lp]] {
					posToPropagate[lp] = true
				}
			}
		}
	}
	return true
}

func (s *Sudoku) Propagate() bool {
	posToPropagate := Set{}
	for pos := 0; pos < sudokuSize*sudokuSize; pos++ {
		if (*s)[pos] == 0 {
			(*s)[pos] = undefiedPos
		}
		if solvedPos[(*s)[pos]] {
			posToPropagate[pos] = true
		}
	}
	return (*s).propagate(posToPropagate)
}

func (s *Sudoku) getChoicesMap() Choices {
	var allChoices = Choices{}
	for pos := 0; pos < sudokuSize*sudokuSize; pos++ {
		if !solvedPos[(*s)[pos]] {
			var posChoices = Set{}
			for choice := range solvedPos {
				var chMask = choice & (*s)[pos]
				if chMask > 0 {
					posChoices[choice] = true
				}
			}
			allChoices[pos] = posChoices
		}
	}
	return allChoices
}

func (s *Sudoku) Solve() (Sudoku, bool) {
	var selectedPos int
	var nbChoices = 10
	var choicesMap = s.getChoicesMap()
	if len(choicesMap) == 0 {
		return *s, true
	}
	for pos, choices := range choicesMap {
		if nbChoices > len(choices) {
			nbChoices = len(choices)
			selectedPos = pos
		}
	}
	for choice := range choicesMap[selectedPos] {
		var result, found = s.explore(selectedPos, choice)
		if found {
			return result, found
		}
	}
	return nil, false
}

func (s *Sudoku) explore(pos int, choice int) (Sudoku, bool) {
	exploringSol := make(Sudoku, len(*s))
	copy(exploringSol, *s)
	exploringSol[pos] = choice
	var posToPropagate = Set{}
	posToPropagate[pos] = true
	if exploringSol.propagate(posToPropagate) {
		return exploringSol.Solve()
	}
	return nil, false
}

func NewSudoku(cells []int) *Sudoku {
	var s = make(Sudoku, sudokuSize*sudokuSize)
	for pos := 0; pos < sudokuSize*sudokuSize; pos++ {
		if cells[pos] == 0 {
			s[pos] = undefiedPos
		} else {
			s[pos] = 1 << (cells[pos] - 1)
		}
	}
	return &s
}

func (s *Sudoku) String() string {
	var out string
	out = ""
	for y := 0; y < sudokuSize; y++ {
		for x := 0; x < sudokuSize; x++ {
			out += fmt.Sprintf("%s", solvedPosDigit[(*s)[pos(x, y)]])
		}
		out += fmt.Sprintf("\n")
	}
	return out
}
