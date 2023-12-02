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
		id, power := parseGame(l)
		total += id
		powers += power
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", powers)
}

func parseGame(s string) (int, int) {
	var id int
	var limits = grab{"red": 12, "green": 13, "blue": 14}
	var lowest = grab{}

	split := strings.Split(s, ":")
	fmt.Sscanf(split[0], "Game %d", &id)

	grabs := strings.Split(split[1], ";")
	for _, grab := range grabs {
		for _, y := range strings.Split(grab, ",") {
			var count int
			var colour string
			fmt.Sscanf(y, " %d %s", &count, &colour)

			if count > limits[colour] {
				id = 0
			}

			if val, ok := lowest[colour]; ok {
				if count > val {
					lowest[colour] = count
				}
			} else {
				lowest[colour] = count
			}
		}
	}
	power := 1
	for _, v := range lowest {
		power *= v
	}
	return id, power
}
