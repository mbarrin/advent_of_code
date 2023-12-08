package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type node struct {
	left, right string
}

type data struct {
	sync.RWMutex
	cycles map[string]int
}

var c = make(chan int)
var foo = make(chan int)
var nums = []int{}

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

	starts, ends := []string{}, []string{}
	for i := 2; i < len(lines); i++ {
		parsed := re.FindAllString(lines[i], -1)
		if parsed[0][2] == 'A' {
			starts = append(starts, parsed[0])
		} else if parsed[0][2] == 'Z' {
			ends = append(ends, parsed[0])
		}

		nodes[parsed[0]] = node{left: parsed[1], right: parsed[2]}
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	go steps("AAA", "ZZZ", instructions, nodes, &wg)
	go getVals(&wg)
	wg.Wait()

	fmt.Println("part 1:", nums[0])

	blah := data{cycles: map[string]int{}}
	for _, start := range starts {
		wg.Add(1)
		go stepsNew(start, instructions, nodes, &wg, &blah)
	}
	wg.Wait()

	params := []int{}
	for _, v := range blah.cycles {
		params = append(params, v)
	}
	fmt.Println("part 2:", LCM(params[0], params[1], params[1:]...))
}

func stepsNew(start string, instructions []string, nodes map[string]node, wg *sync.WaitGroup, b *data) {
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

		counter++

		if strings.HasSuffix(key, "Z") {
			ok = false
		}
	}

	b.Lock()
	b.cycles[start] = counter
	b.Unlock()

	wg.Done()
}

func steps(start, end string, instructions []string, nodes map[string]node, wg *sync.WaitGroup) {
	counter := 0

	key := start
	for key != end {
		i := instructions[counter%len(instructions)]
		if i == "L" {
			key = nodes[key].left
		} else if i == "R" {
			key = nodes[key].right
		}
		counter++
	}
	c <- counter
	close(c)
	wg.Done()
}

func getVals(wg *sync.WaitGroup) {
	for y := range c {
		nums = append(nums, y)
	}
	wg.Done()
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
