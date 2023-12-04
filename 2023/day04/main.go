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

		for _, a := range strings.Split(x, " ") {
			if a != "" {
				data[a]++
			}
		}

		foo := 0
		for _, v := range data {
			if v > 1 {
				foo += 1
			}
		}

		if foo != 0 {
			total += math.Pow(2.0, float64(foo-1))
		}

		cards += multipliers[counter]
		for i := 1; i <= foo; i++ {
			multipliers[counter+i] += multipliers[counter]
		}
		counter++
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", cards)
}
