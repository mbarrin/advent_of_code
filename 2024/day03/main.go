package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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
	stringedInput := string(input)

	fmt.Println("part 1:", getMults(stringedInput))

	cleaned := removeMults(stringedInput)
	fmt.Println("part 2:", getMults(cleaned))
}

func getMults(input string) int {
	matches := pattern.FindAllString(input, -1)

	total := 0
	for _, x := range matches {
		one, two := 0, 0
		fmt.Sscanf(x, "mul(%d,%d)", &one, &two)

		total += (one * two)
	}
	return total
}

func removeMults(input string) string {
	donts := negatePattern.FindAllString(input, -1)
	for _, x := range donts {
		input = strings.Replace(input, x, "", 1)
	}
	blah := strings.Split(input, "don't()")
	return blah[0]
}
