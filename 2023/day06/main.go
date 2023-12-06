package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`(\d+)`)

func main() {

	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(1)
	}

	nums := r.FindAllString(string(f), -1)
	info := map[int]int{}
	for i := 0; i < len(nums)/2; i++ {
		key, _ := strconv.Atoi(nums[i])
		val, _ := strconv.Atoi(nums[i+len(nums)/2])
		info[key] = val
	}

	total := 1
	for k, v := range info {
		total *= succesful(k, v)
	}

	fmt.Println("part 1:", total)

	t, d := strings.Join(nums[:len(nums)/2], ""), strings.Join(nums[len(nums)/2:], "")

	time, _ := strconv.Atoi(t)
	dist, _ := strconv.Atoi(d)

	fmt.Println("part 2:", succesful(time, dist))
}

func succesful(t, d int) (count int) {
	for i := 1; i < t; i++ {
		if i*(t-i) > d {
			count++
		}
	}
	return count
}
