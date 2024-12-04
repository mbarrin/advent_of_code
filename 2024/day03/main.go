package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

var (
	pattern       = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	negatePattern = regexp.MustCompile(`(?s)don't\(\).*?do\(\)`)
)

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("part 1:", sumMults(input))
	fmt.Println("part 2:", sumMults(removeMults(input)))
}

func sumMults(input []byte) (total int) {
	matches := pattern.FindAll(input, -1)

	for _, x := range matches {
		var one, two int
		fmt.Sscanf(string(x), "mul(%d,%d)", &one, &two)
		total += (one * two)
	}
	return total
}

func removeMults(input []byte) []byte {
	input = negatePattern.ReplaceAll(input, []byte{})
	return bytes.Split(input, []byte("don't()"))[0]
}
