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

	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	safeOne := 0
	safeTwo := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		temp := []int{}
		hash := map[int]int{}
		dupe := 0
		for _, x := range line {
			num, _ := strconv.Atoi(x)
			hash[num]++
			if hash[num] > 1 {
				dupe++
			}
			if hash[num] > 2 || dupe > 1 {
				continue
			}
			temp = append(temp, num)
		}

		foo := make([]int, len(temp))
		copy(foo, temp)

		if isSafeOne(temp) {
			safeOne++
		}
		if isSafeTwo(foo) {
			safeTwo++
		}

	}
	fmt.Println("part 1:", safeOne)
	fmt.Println("part 2:", safeTwo)
}

func isSafeOne(s []int) bool {
	if safeGrowthOne(s) {
		return true
	}
	slices.Reverse(s)
	return safeGrowthOne(s)
}

func safeGrowthOne(s []int) bool {
	for i := 0; i < len(s)-1; i += 1 {
		diff := s[i+1] - s[i]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeTwo(s []int) bool {
	fmt.Println()
	foo := make([]int, len(s))
	copy(foo, s)

	if safeGrowthTwo(s) {
		return true
	}
	slices.Reverse(foo)
	return safeGrowthTwo(foo)
}

func safeGrowthTwo(s []int) bool {
	foo := make([]int, len(s))
	copy(foo, s)

	failure := 0
	for i := 1; i < len(s)-1; i += 1 {
		if failure > 1 {
			fmt.Println("failed:", foo)
			return false
		}
		// preDiff := s[i] - s[i-1]
		// postDiff := s[i+1] - s[i]

		// if gap between i and i-1 > 3 and gap between i and i+1 > 3: remove
		// if gap between i and i+1 > 3: remove i

		//if s[i] > s[i-1] && s[i] > s[i+1] {
		//	failure++
		//	s = slices.Delete(s, i, i+1)
		//	fmt.Println("hmm after:", s)
		//	i--
		//} else if s[i] < s[i-1] && s[i] < s[i+1] {
		//	failure++
		//	s = slices.Delete(s, i, i+1)
		//	fmt.Println("hah after:", s)
		//	i--
		//} else if preDiff < 1 || preDiff > 3 {
		//	failure++
		//	fmt.Println("pre before:", s)
		//	s = slices.Delete(s, i-1, i)
		//	i--
		//	fmt.Println("pre after:", s)
		//} else if postDiff < 1 || postDiff > 3 {
		//	failure++
		//	fmt.Println("post before:", s)
		//	s = slices.Delete(s, i+1, i+2)
		//	i--
		//	fmt.Println("post after:", s)
		//}
	}
	return failure < 2
}
