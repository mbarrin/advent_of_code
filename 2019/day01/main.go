package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()

	s := bufio.NewScanner(input)

	totalOne, totalTwo := 0, 0
	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		totalOne += fuel(num)
		totalTwo += maxFuel(num)
	}
	fmt.Println("part 1:", totalOne)
	fmt.Println("part 2:", totalTwo)
}

func fuel(i int) int {
	return (i / 3) - 2
}

func maxFuel(i int) int {
	f := fuel(i)
	if f >= 0 {
		return f + maxFuel(f)
	}
	return 0
}
