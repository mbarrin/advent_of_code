package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}

	s := bufio.NewScanner(f)

	var (
		left, right []int
		countMap    = make(map[int]int)
	)

	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, " ")

		leftNum, err := strconv.Atoi(split[0])
		if err != nil {
			os.Exit(1)
		}
		left = append(left, leftNum)

		rightNum, err := strconv.Atoi(split[len(split)-1])
		if err != nil {
			os.Exit(1)
		}
		right = append(right, rightNum)

		countMap[rightNum]++
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	similarity := 0

	for i := range left {
		dist := (right[i] - left[i])
		if dist < 0 {
			dist = 0 - dist
		}
		distance += dist

		if val, ok := countMap[left[i]]; ok {
			similarity += left[i] * val
		}
	}

	fmt.Println("part 1: ", distance)

	fmt.Println("part 2: ", similarity)
}
