package sudoku

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

type Set map[int]bool

const sudokuSize = 9

func SqId(x, y int) int {
	return (x/3 + (y/3)*3)
}

func isValid(cells []int) bool {
	var list = make([]Set, 27)
	for i := 0; i < 27; i++ {
		list[i] = make(Set, 9)
	}
	for x := 0; x < sudokuSize; x++ {
		for y := 0; y < sudokuSize; y++ {
			sqId := SqId(x, y)
			pos := x + 9*y
			list[x][cells[pos]] = true
			list[y+sudokuSize][cells[pos]] = true
			list[sqId+2*sudokuSize][cells[pos]] = true
		}
	}
	nb := 0
	for _, set := range list {
		if val := len(set); val == 9 {
			nb++
		}
	}
	return nb == 27
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string
	var cells []int
	for i := 0; i < 9; i++ {
		scanner.Scan()
		inputs = strings.Split(scanner.Text(), " ")
		for j := 0; j < 9; j++ {
			n, _ := strconv.ParseInt(inputs[j], 10, 32)
			cells = append(cells, int(n))
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(isValid(cells)) // Write answer to stdout
}
