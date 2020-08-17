package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	rowLength    = 9
	minArgsCount = 2
)

type Board [9][9]int

func NewBoard(b [9][9]int) *Board {
	board := Board(b)
	return &board
}

func (b *Board) Solve() bool {
	if !b.hasEmptyCell() {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for num := 9; num >= 1; num-- {
					b[i][j] = num
					if b.isValid() {
						if b.Solve() {
							return true
						}
						b[i][j] = 0
					} else {
						b[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func (b *Board) hasEmptyCell() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func (b *Board) isValid() bool {
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[b[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[b[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[b[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func (b *Board) String() string {
	var builder strings.Builder
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			builder.WriteString(fmt.Sprintf("%d", b[row][col]))
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < minArgsCount {
		fmt.Printf("Usage: %s <file containing sudoku puzzle>\n", os.Args[0])
		os.Exit(1)
	}
	fileInput := os.Args[1]
	f, err := os.Open(fileInput)
	checkErr(err)

	lineRegxp := regexp.MustCompile("^[0-9]{9}$")

	scanner := bufio.NewScanner(f)
	var grid [9][9]int
	var l int
	for scanner.Scan() {
		line := scanner.Text()
		if lineRegxp.MatchString(line) {
			i := l % 9
			if i == 0 {
				grid = [9][9]int{}
			}
			for j, d := range line {
				grid[i][j], _ = strconv.Atoi(string(d))
			}
			if i == rowLength-1 {
				board := NewBoard(grid)
				if board.Solve() {
					fmt.Print(board)
					fmt.Printf("First sum of first 3 numbers: %d\n\n", board[0][0]+board[0][1]+board[0][2])
				} else {
					fmt.Print(board)
					fmt.Println("Unsolvable")
				}
			}
			l++
		} else {
			fmt.Println(line)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
