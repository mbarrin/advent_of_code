package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	rowID int
	colID int
}

type points map[point]byte

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)

	grid := make(points)
	rowID := 0
	for s.Scan() {
		for colID, item := range s.Bytes() {
			if item != '.' {
				p := point{rowID, colID}
				grid[p] = item
			}
		}
		rowID++
	}

	total := []int{}
	power := 0
	for p, v := range grid {
		if !isNumber(v) {
			count := 0
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					edge := point{p.rowID + i, p.colID + j}
					if val, ok := grid[edge]; ok {
						temp := []byte{}
						if isNumber(val) {
							temp = append(temp, val)
							delete(grid, edge)
							x := point{edge.rowID, edge.colID - 1}
							left := grid[x]
							y := point{edge.rowID, edge.colID + 1}
							right := grid[y]
							if isNumber(left) {
								temp = append([]byte{left}, temp...)
								delete(grid, x)
								x := point{edge.rowID, edge.colID - 2}
								left = grid[x]
								if isNumber(left) {
									temp = append([]byte{left}, temp...)
									delete(grid, x)
								}
							}
							if isNumber(right) {
								temp = append(temp, right)
								delete(grid, y)
								y := point{edge.rowID, edge.colID + 2}
								right = grid[y]
								if isNumber(right) {
									temp = append(temp, right)
									delete(grid, y)
								}
							}
							total = append(total, toInt(temp))

							if v == '*' {
								count++
							}
						}
					}
				}
			}
			if count == 2 {
				power += total[len(total)-2] * total[len(total)-1]
			}
		}
	}

	fmt.Println("part 1:", sum(total))
	fmt.Println("part 2:", power)
}

func sum(i []int) int {
	total := 0
	for _, x := range i {
		total += x
	}
	return total
}

func toInt(nums []byte) int {
	b, _ := strconv.Atoi(string(nums))
	return b
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}
