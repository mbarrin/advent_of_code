package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`(\d+)`)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)

	x := []string{}
	y := []string{}
	for s.Scan() {
		a := strings.Join(r.FindAllString(s.Text(), -1), " ")
		b := strings.Join(r.FindAllString(s.Text(), -1), "")
		x = append(x, a)
		y = append(y, b)
	}

	var one, two, three, four, five, six, seven, eight int
	time := x[0]
	fmt.Sscanf(time, "%d %d %d %d", &one, &two, &three, &four)
	distance := x[1]
	fmt.Sscanf(distance, "%d %d %d %d", &five, &six, &seven, &eight)

	foo := map[int]int{
		one:   five,
		two:   six,
		three: seven,
		four:  eight,
	}

	total := 1
	for k, v := range foo {
		total *= succesful(k, v)
	}

	fmt.Println("part 1:", total)

	time = y[0]
	fmt.Sscanf(time, "%d", &one)
	distance = y[1]
	fmt.Sscanf(distance, "%d", &two)
	fmt.Println("part 2:", succesful(one, two))
}

func succesful(t, d int) int {
	if t == 0 {
		return 1
	}
	count := 0
	for i := 1; i < t; i++ {
		if i*(t-i) > d {
			count++
		}
	}
	return count
}
