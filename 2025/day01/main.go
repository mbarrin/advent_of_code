package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

var (
	pattern = regexp.MustCompile(`(\w)(\d+)`)
)

func main() {
	defer util.TimeTaken(time.Now())

	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}

	location := 50
	totalOne := 0
	totalTwo := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		var direction string
		var count int
		loops := 0

		matches := pattern.FindAllStringSubmatch(s.Text(), -1)

		direction = matches[0][1]
		count, err := strconv.Atoi(matches[0][2])
		if err != nil {
			os.Exit(1)
		}

		if direction == "L" {
			location, loops = rotate(location, -count)
		} else {
			location, loops = rotate(location, count)
		}

		if location%100 == 0 {
			totalOne++
			totalTwo++
		}
		totalTwo += loops
	}

	fmt.Println("part 1: ", totalOne)
	fmt.Println("part 2: ", totalTwo)
}

func rotate(start int, count int) (int, int) {
	location := start + count
	loops := 0

	if location >= 100 {
		loops += util.Abs(location / 100)
		location %= 100
		if location%100 == 0 {
			loops--
		}
	} else if location < 0 {
		if start != 0 {
			loops++
		}
		loops += util.Abs(location / 100)
		location %= 100
		if location < 0 {
			location += 100
		}
		if location%100 == 0 {
			loops--
		}
	}

	return location, loops
}
