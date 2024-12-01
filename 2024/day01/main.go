package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

func main() {
	defer util.TimeTaken(time.Now())

	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}

	var (
		left, right          []int
		countMap             = make(map[int]int)
		distance, similarity int
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		var leftNum, rightNum int
		fmt.Sscanf(s.Text(), "%d   %d", &leftNum, &rightNum)

		left = append(left, leftNum)
		right = append(right, rightNum)

		countMap[rightNum]++
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		distance += util.Abs(right[i] - left[i])

		if val, ok := countMap[left[i]]; ok {
			similarity += left[i] * val
		}
	}

	fmt.Println("part 1: ", distance)

	fmt.Println("part 2: ", similarity)
}
