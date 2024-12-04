package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

type point struct {
	x, y int
}

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer input.Close()

	grid := map[point]byte{}

	s := bufio.NewScanner(input)
	x := 0
	for s.Scan() {
		for y, char := range s.Bytes() {
			grid[point{x: x, y: y}] = char
		}
		x++
	}

	totalOne, totalTwo := 0, 0
	for p, b := range grid {
		if b == 'X' {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if checkDirection(grid, p, 'M', i, j) {
						totalOne++
					}
				}
			}
		}
		if b == 'A' {
			if checkCross(grid, p) {
				totalTwo++
			}
		}
	}
	fmt.Println("part 1:", totalOne)
	fmt.Println("part 2:", totalTwo)
}

func checkDirection(grid map[point]byte, p point, b byte, x int, y int) bool {
	if grid[point{p.x + x, p.y + y}] != b {
		return false
	}

	var next byte
	if b == 'M' {
		next = 'A'
	} else if b == 'A' {
		next = 'S'
	} else if b == 'S' {
		return true
	} else {
		return false
	}

	return checkDirection(grid, point{p.x + x, p.y + y}, next, x, y)
}

func checkCross(grid map[point]byte, p point) bool {
	topLeft := grid[point{p.x - 1, p.y - 1}]
	topRight := grid[point{p.x - 1, p.y + 1}]
	bottomLeft := grid[point{p.x + 1, p.y - 1}]
	bottomRight := grid[point{p.x + 1, p.y + 1}]

	temp := []byte{topLeft, topRight, bottomLeft, bottomRight}
	valid := []string{"MMSS", "MSMS", "SSMM", "SMSM"}

	return slices.Contains(valid, string(temp))
}
