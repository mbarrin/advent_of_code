package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	start, end := 0, 0
	for s.Scan() {
		data := map[int][]int{}
		line := strings.Split(s.Text(), " ")

		for _, x := range line {
			a, _ := strconv.Atoi(x)
			data[0] = append(data[0], a)
		}

		reduce(data, 0)

		end += data[0][len(data[0])-1]
		start += data[0][0]
	}

	fmt.Println("part 1:", end)
	fmt.Println("part 2:", start)
}

func reduce(data map[int][]int, level int) {
	allZero := true
	for i := 0; i < len(data[level])-1; i++ {
		tmp := data[level][i+1] - data[level][i]
		data[level+1] = append(data[level+1], tmp)
		if tmp != 0 {
			allZero = false
		}
	}

	if !allZero {
		reduce(data, level+1)
	} else {
		currentLevel := level + 1
		for currentLevel != 0 {
			a := data[currentLevel][len(data[currentLevel])-1]
			b := data[currentLevel-1][len(data[currentLevel-1])-1]
			data[currentLevel-1] = append(data[currentLevel-1], a+b)

			x := data[currentLevel][0]
			y := data[currentLevel-1][0]
			data[currentLevel-1] = append([]int{y - x}, data[currentLevel-1]...)

			currentLevel--
		}
	}
}
