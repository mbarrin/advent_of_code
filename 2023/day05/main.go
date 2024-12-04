package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type mapping struct {
	to   string
	info []info
}

type info struct {
	min, max int
	offset   int
}

type shortcut struct {
	level   string
	current int
}

var transformations = make(map[string]*mapping)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	seeds := strings.Split(lines[0], " ")[1:]

	var name, dest string

	for i := 1; i < len(lines); i++ {
		if strings.HasSuffix(lines[i], "map:") {
			mappingLine := strings.Split(lines[i][:len(lines[i])-5], "-")
			name, dest = mappingLine[0], mappingLine[2]

			transformations[name] = &mapping{to: dest}
		} else if lines[i] == "" {
		} else {
			var to, from, count int
			fmt.Sscanf(lines[i], "%d %d %d", &to, &from, &count)
			mappingInfo := info{min: from, max: from + count - 1, offset: to - from}
			transformations[name].info = append(transformations[name].info, mappingInfo)
		}
	}

	totals := []int{}
	for _, seed := range seeds {
		i, _ := strconv.Atoi(seed)
		totals = append(totals, search("seed", i, i))
	}

	fmt.Println("part 1:", slices.Min(totals))

	totals = []int{}
	for i := 0; i < len(seeds)-1; i += 2 {
		foo, _ := strconv.Atoi(seeds[i])
		bar, _ := strconv.Atoi(seeds[i+1])
		for j := foo; j < foo+bar; j++ {
			totals = append(totals, search("seed", j, j))
		}
	}

	fmt.Println("part 2:", slices.Min(totals))
}

func search(name string, i, current int) int {
	for _, x := range transformations[name].info {
		if i >= x.min && i <= x.max {
			i += x.offset
			break
		}
	}
	if transformations[name].to != "location" {
		i = search(transformations[name].to, i, current)
	}

	return i
}
