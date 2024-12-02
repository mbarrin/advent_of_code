package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

func main() {
	defer util.TimeTaken(time.Now())

	f, err := os.Open("sample.txt")
	if err != nil {
		os.Exit(1)
	}

	safe := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		temp := []int{}
		for _, x := range line {
			num, _ := strconv.Atoi(x)
			temp = append(temp, num)
		}

		if isSafe(temp) {
			safe++
		}

	}
	fmt.Println("part 1:", safe)
}

func isSafe(s []int) bool {
	if safeGrowth(s) {
		return true
	}
	slices.Reverse(s)
	if safeGrowth(s) {
		return true
	}
	return false
}

func safeGrowth(s []int) bool {
	failures := 0
	for i := 0; i < len(s)-1; i += 1 {
		diff := s[i+1] - s[i]
		if diff < 1 || diff > 3 {
			failures += 1
			return false
		}
	}
	return true
}
