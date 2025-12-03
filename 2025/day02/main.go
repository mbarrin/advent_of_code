package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	pattern := regexp.MustCompile(`(\d+)-(\d+)`)

	matches := pattern.FindAllSubmatch(input, -1)

	defer util.TimeTaken(time.Now())

	totalOne := 0
	totalTwo := 0
	for _, match := range matches {
		totalOne += valid(match[1], match[2])
		totalTwo += newValid(match[1], match[2])
	}
	fmt.Println("part 1", totalOne)
	fmt.Println("part 2", totalTwo)
}

func valid(start, end []byte) int {
	startInt, err := strconv.Atoi(string(start))
	if err != nil {
		log.Fatal(err)
	}
	endInt, err := strconv.Atoi(string(end))
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for i := startInt; i <= endInt; i++ {
		x := []byte(strconv.Itoa(i))
		if len(x)%2 != 0 {
			continue
		}
		firstHalf := x[0 : len(x)/2]
		secondHalf := x[len(x)/2:]
		if slices.Compare(firstHalf, secondHalf) == 0 {
			count += i
		}
	}
	return count
}

func newValid(start, end []byte) int {
	startInt, err := strconv.Atoi(string(start))
	if err != nil {
		log.Fatal(err)
	}
	endInt, err := strconv.Atoi(string(end))
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for i := startInt; i <= endInt; i++ {
		x := []byte(strconv.Itoa(i))
		if len(x)%2 == 0 {
			one := x[0 : len(x)/2]
			two := x[len(x)/2:]
			if slices.Compare(one, two) == 0 {
				count += i
				continue
			}
		}
		if len(x)%3 == 0 {
			one := x[0 : len(x)/3]
			two := x[len(x)/3 : len(x)/3*2]
			three := x[len(x)/3*2:]
			if slices.Compare(one, two) == 0 && slices.Compare(two, three) == 0 {
				count += i
				continue
			}
		}
		if len(x)%4 == 0 {
			one := x[0 : len(x)/4]
			two := x[len(x)/4 : len(x)/2]
			three := x[len(x)/2 : len(x)/4*3]
			four := x[len(x)/4*3:]
			if slices.Compare(one, two) == 0 && slices.Compare(two, three) == 0 && slices.Compare(three, four) == 0 {
				count += i
				continue
			}
		}
		if len(x)%5 == 0 {
			one := x[0 : len(x)/5]
			two := x[len(x)/5 : len(x)/5*2]
			three := x[len(x)/5*2 : len(x)/5*3]
			four := x[len(x)/5*3 : len(x)/5*4]
			five := x[len(x)/5*4:]
			if slices.Compare(one, two) == 0 && slices.Compare(two, three) == 0 && slices.Compare(three, four) == 0 && slices.Compare(four, five) == 0 {
				count += i
				continue
			}
		}
		if len(slices.Compact(x)) == 1 && len(x) > 1 {
			count += i
		}
	}
	return count
}
