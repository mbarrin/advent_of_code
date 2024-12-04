package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
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

	wg := sync.WaitGroup{}
	output := make(chan int)
	totals := []int{}
	for _, seed := range seeds {
		wg.Add(1)
		i, _ := strconv.Atoi(seed)
		go func() {
			defer wg.Done()
			search("seed", i, &wg, output)
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for i := range output {
		totals = append(totals, i)
	}

	fmt.Println("part 1:", slices.Min(totals))

	wg = sync.WaitGroup{}
	output = make(chan int)
	totals = []int{}
	for i := 0; i < len(seeds)-1; i += 2 {
		foo, _ := strconv.Atoi(seeds[i])
		bar, _ := strconv.Atoi(seeds[i+1])
		for j := foo; j < foo+bar; j++ {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				search("seed", x, &wg, output)
			}(j)
		}
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for i := range output {
		totals = append(totals, i)
	}
	fmt.Println("part 2:", slices.Min(totals))

}

func search(name string, i int, wg *sync.WaitGroup, c chan int) int {
	for _, x := range transformations[name].info {
		if i >= x.min && i <= x.max {
			i += x.offset
			break
		}
	}
	if transformations[name].to != "location" {
		i = search(transformations[name].to, i, wg, c)
	}

	c <- i
	return i
}
