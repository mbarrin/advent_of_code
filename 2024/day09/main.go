package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

// even is file
// odd is space

type node struct {
	before *node
	after  *node
	value  int
}

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("part 1:", basic(input))
	fmt.Println("part 2:", advanced(input))
}

func basic(input []byte) int {
	output := expand(input)

	keys := []int{}
	for k := range output {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	defrag(output, keys)

	return checkSum(output, keys)
}

func advanced(input []byte) int {
	start := expandToNode(input)

	for start.after != nil {
		fmt.Println(start)
		start = *start.after
	}
	fmt.Println(start)
	return 0
}

func checkSum(output map[int]int, keys []int) int {
	checksum := 0
	for _, x := range keys {
		if output[x] == -1 {
			break
		}
		checksum += x * output[x]
	}

	return checksum
}

func expandToNode(input []byte) node {
	id, loc := 0, 0

	var current node
	for i, x := range input {
		for j := 0; j < int(x)-48; j++ {
			if i%2 == 0 {
				prev := current
				current = node{value: id, before: &prev}
			} else {
				prev := current
				current = node{value: -1, before: &prev}
			}
		}
		if i%2 == 0 {
			id++
		}
		loc++
	}

	end := current
	current = *current.before
	current.after = &end

	for current.before != nil {
		temp := current
		current.before.after = &temp
		current = *current.before
	}

	return current
}

func expand(input []byte) map[int]int {
	output := map[int]int{}
	id, loc := 0, 0
	for i, x := range input {
		for j := 0; j < int(x)-48; j++ {
			if i%2 == 0 {
				output[loc] = id
			} else {
				output[loc] = -1
			}
			loc++
		}
		if i%2 == 0 {
			id++
		}
	}
	return output
}

func defrag(output map[int]int, keys []int) {
	lastFree := len(keys) - 1
	for _, x := range keys {
		if output[x] == -1 {
			output[x] = output[lastFree]
			output[lastFree] = -1

			for output[lastFree] == -1 {
				lastFree--
			}
		}
		if x >= lastFree {
			break
		}
	}
}

func disk(d map[int]int, k []int) string {
	output := ""
	for _, x := range k {
		if d[x] == -1 {
			output += "."
		} else {
			output += strconv.Itoa(d[x])
		}
	}
	return output
}
