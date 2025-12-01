package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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
		os.Exit(1)
	}

	grid := map[point]int{}
	intersections := []int{}
	//travelled := map[point]int{}

	s := bufio.NewScanner(input)

	dist := 0
	for s.Scan() {
		location := point{0, 0}

		steps := strings.Split(s.Text(), ",")

		for i, step := range steps {
			num, _ := strconv.Atoi(string(step[1:]))

			switch step[0] {
			case 'R':
				for i := 0; i < num; i++ {
					location.y++
					grid[location]++
					if grid[location] > 1 {
						intersections = append(intersections, util.ManhattenDistance(0, 0, location.x, location.y))
					}
				}
				dist += num
			case 'D':
				for i := 0; i < num; i++ {
					location.x++
					grid[location]++
					if grid[location] > 1 {
						intersections = append(intersections, util.ManhattenDistance(0, 0, location.x, location.y))
					}
				}
				dist += num
			case 'L':
				for i := 0; i < num; i++ {
					location.y--
					grid[location]++
					if grid[location] > 1 {
						intersections = append(intersections, util.ManhattenDistance(0, 0, location.x, location.y))
					}
				}
				dist += num
			case 'U':
				for i := 0; i < num; i++ {
					location.x--
					grid[location]++
					if grid[location] > 1 {
						intersections = append(intersections, util.ManhattenDistance(0, 0, location.x, location.y))
					}
				}
				dist += num
			}
			fmt.Println("i, dist", i, dist)
		}
	}

	fmt.Println("part 1:", slices.Min(intersections))
}
