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

	f, err := os.Open("badsample.txt")
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
	if safeGrowth(s, 0) < 2 {
		return true
	}
	//slices.Reverse(s)
	//return safeGrowth(s, 0) < 2
	//if safeGrowth(s, 0) < 2 {
	//	return true
	//}
	fmt.Println(s)
	return false
}

func safeGrowth(s []int, failures int) int {
	if failures > 1 {
		return failures
	}
	fmt.Println(s)

	for i := 0; i < len(s)-1; i += 1 {
		if s[i+1]-s[i] < 1 || s[i+1]-s[i] > 3 {
			if i == len(s)-2 && failures < 1 {
				fmt.Println(i)
				return failures
			}

			sCopy := make([]int, len(s))
			copy(sCopy, s)
			temp := slices.Delete(sCopy, i, i+1)
			failures += safeGrowth(temp, failures+1)
		}
	}
	return failures
}
