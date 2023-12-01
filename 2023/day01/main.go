package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	lookupOne = map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}

	lookupTwo = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}

	totals = map[int]int{1: 0, 2: 0}
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		s := s.Text()
		totals[1] += find(s, lookupOne)
		totals[2] += find(s, lookupTwo)
	}
	fmt.Println("part 1:", totals[1])
	fmt.Println("part 2:", totals[2])
}

func find(s string, l map[string]int) int {
	return findFirst(s, l)*10 + findLast(s, l)
}

func findFirst(s string, lookup map[string]int) int {
	for {
		for k, v := range lookup {
			if strings.HasPrefix(s, k) {
				return v
			}
		}
		s = s[1:]
	}
}

func findLast(s string, lookup map[string]int) int {
	for {
		for k, v := range lookup {
			if strings.HasSuffix(s, k) {
				return v
			}
		}
		s = s[:len(s)-1]
	}
}
