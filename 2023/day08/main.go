package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type node struct {
	left, right string
}

type data struct {
	sync.RWMutex
	sync.WaitGroup
	cycles map[string]int
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(1)
	}

	s := bufio.NewScanner(f)

	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	instructions := strings.Split(lines[0], "")

	nodes := make(map[string]node)
	re := regexp.MustCompile(`(\w+)`)

	starts := []string{}
	for i := 2; i < len(lines); i++ {
		parsed := re.FindAllString(lines[i], -1)
		if parsed[0][2] == 'A' {
			starts = append(starts, parsed[0])
		}
		nodes[parsed[0]] = node{left: parsed[1], right: parsed[2]}
	}

	a := time.Now()
	d := data{cycles: map[string]int{}}
	for _, start := range starts {
		d.Add(1)
		go steps(start, instructions, nodes, &d)
	}
	d.Wait()

	params := []int{}
	for _, v := range d.cycles {
		params = append(params, v)
	}

	fmt.Println("part 1:", d.cycles["AAA"])
	fmt.Println("part 2:", LCM(params[0], params[1], params[1:]...))
	fmt.Printf("\nTotal time: %s\n", time.Since(a))
}

func steps(start string, instructions []string, nodes map[string]node, d *data) {
	counter := 0
	key := start
	ok := true

	for ok {
		i := instructions[counter%len(instructions)]

		if i == "L" {
			key = nodes[key].left
		} else if i == "R" {
			key = nodes[key].right
		}

		if strings.HasSuffix(key, "Z") {
			ok = false
		}

		counter++
	}

	d.Lock()
	d.cycles[start] = counter
	d.Unlock()

	d.Done()
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
