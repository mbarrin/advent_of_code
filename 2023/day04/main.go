package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)

	var total float64
	var cards int
	multipliers := map[int]int{}
	counter := 1
	for s.Scan() {
		data := map[string]int{}
		multipliers[counter]++

		l := s.Text()

		t := strings.Split(l, ":")
		t = strings.Split(t[1], "|")
		x := strings.Join(t, "")

		wins := 0
		for _, a := range strings.Split(x, " ") {
			if a != "" {
				data[a]++
				if data[a] == 2 {
					wins += 1
				}
			}
		}

		if wins != 0 {
			total += math.Pow(2.0, float64(wins-1))
		}

		cards += multipliers[counter]
		for i := 1; i <= wins; i++ {
			multipliers[counter+i] += multipliers[counter]
		}
		counter++
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", cards)
}
