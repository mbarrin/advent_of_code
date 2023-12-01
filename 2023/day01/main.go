package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var lookup = map[string]int{
	"one": 1, "two": 2, "three": 3,
	"four": 4, "five": 5, "six": 6,
	"seven": 7, "eight": 8, "nine": 9,
}

var foo = regexp.MustCompile(`twone|eightwo|oneight|eighthree|sevenine|one|two|three|four|five|six|seven|eight|nine|[0-9]`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var totals = map[int]int{1: 0, 2: 0}
	for s.Scan() {
		nums := map[int][]int{1: {}, 2: {}}
		bar := foo.FindAll(s.Bytes(), -1)
		for _, x := range bar {
			if len(x) > 1 {
				if string(x) == "twone" {
					nums[2] = append(nums[2], 2, 1)
				} else if string(x) == "oneight" {
					nums[2] = append(nums[2], 1, 8)
				} else if string(x) == "eighthree" {
					nums[2] = append(nums[2], 8, 3)
				} else if string(x) == "sevenine" {
					nums[2] = append(nums[2], 7, 9)
				} else if string(x) == "eightwo" {
					nums[2] = append(nums[2], 8, 2)
				} else {
					nums[2] = append(nums[2], lookup[string(x)])
				}
			} else {
				for k := range nums {
					nums[k] = append(nums[k], (int(x[0] - 48)))
				}
			}
		}
		for k, v := range nums {
			if len(v) == 1 {
				totals[k] = totals[k] + v[0]*10 + v[0]
			} else if len(v) > 1 {
				totals[k] = totals[k] + v[0]*10 + v[len(v)-1]
			}
		}
	}
	fmt.Println("part 1:", totals[1])
	fmt.Println("part 2:", totals[2])
}
