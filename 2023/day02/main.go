package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grab map[string]int

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)

	total := 0
	powers := 0
	for s.Scan() {
		l := s.Text()
		total += parseGame(l)
		powers += minCubes(l)
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", powers)
}

func parseGame(s string) int {
	var id int
	var limits = grab{"red": 12, "green": 13, "blue": 14}

	split := strings.Split(s, ":")
	fmt.Sscanf(split[0], "Game %d", &id)

	grabs := strings.Split(split[1], ";")
	for _, grab := range grabs {
		for _, y := range strings.Split(grab, ",") {
			var count int
			var colour string
			fmt.Sscanf(y, " %d %s", &count, &colour)

			if count > limits[colour] {
				return 0
			}
		}
	}
	return id
}

func minCubes(s string) int {
	var limits = grab{}

	split := strings.Split(s, ":")
	grabs := strings.Split(split[1], ";")
	for _, grab := range grabs {
		for _, y := range strings.Split(grab, ",") {
			var count int
			var colour string
			fmt.Sscanf(y, " %d %s", &count, &colour)

			if val, ok := limits[colour]; ok {
				if count > val {
					limits[colour] = count
				}
			} else {
				limits[colour] = count
			}
		}
	}
	power := 1
	for _, v := range limits {
		power *= v
	}

	return power
}
