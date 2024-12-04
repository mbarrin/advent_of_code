package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type point struct {
	rowID, colID int
}

type points map[point][]point

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(1)
	}

	scanner := bufio.NewScanner(f)

	rowID := 0
	p := make(points)
	start := point{}
	for scanner.Scan() {
		for colID, x := range scanner.Text() {
			connections := []point{}
			switch x {
			case '|':
				connections = []point{{rowID - 1, colID}, {rowID + 1, colID}}
			case '-':
				connections = []point{{rowID, colID - 1}, {rowID, colID + 1}}
			case 'L':
				connections = []point{{rowID - 1, colID}, {rowID, colID + 1}}
			case 'J':
				connections = []point{{rowID - 1, colID}, {rowID, colID - 1}}
			case '7':
				connections = []point{{rowID + 1, colID}, {rowID, colID - 1}}
			case 'F':
				connections = []point{{rowID + 1, colID}, {rowID, colID + 1}}
			case 'S':
				start = point{rowID: rowID, colID: colID}
			default:
				continue
			}
			p[point{rowID: rowID, colID: colID}] = connections

		}
		rowID++
	}

	n := point{start.rowID, start.colID - 1}
	if val, ok := p[n]; ok {
		if slices.Contains(val, start) {
			p[start] = append(p[start], n)
		}
	}

	s := point{start.rowID, start.colID + 1}
	if val, ok := p[s]; ok {
		if slices.Contains(val, start) {
			p[start] = append(p[start], s)
		}
	}

	e := point{start.rowID + 1, start.colID}
	if val, ok := p[e]; ok {
		if slices.Contains(val, start) {
			p[start] = append(p[start], e)
		}
	}

	w := point{start.rowID - 1, start.colID}
	if val, ok := p[w]; ok {
		if slices.Contains(val, start) {
			p[start] = append(p[start], w)
		}
	}

	fmt.Println(travel(p, &start, &start, nil, 0))
}

func travel(points map[point][]point, start, current, previous *point, counter int) int {
	next := points[*current][0]

	if previous != nil && next == *previous {
		next = points[*current][1]
	}

	if next == *start && counter != 0 {
		return counter/2 + 1
	}

	return travel(points, start, &next, current, counter+1)
}
