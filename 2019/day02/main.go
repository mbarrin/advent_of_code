package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	inputSplit := strings.Split(strings.TrimSpace(string(input)), ",")
	codes := []int{}

	for _, x := range inputSplit {
		op, _ := strconv.Atoi(x)
		codes = append(codes, op)
	}

	codes[1], codes[2] = 12, 2
	fmt.Println("part 1:", compute(codes)[0])

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			codes[1], codes[2] = noun, verb
			result := compute(codes)
			if result[0] == 19690720 {
				fmt.Println("part 2:", 100*noun+verb)
				break
			}
		}
	}
}

func compute(codes []int) []int {
	local := make([]int, len(codes))
	copy(local, codes)

	for i := 0; i < len(local)-4; i++ {
		dest := local[i+3]
		one := local[i+1]
		two := local[i+2]

		switch local[i] {
		case 1:
			local[dest] = local[one] + local[two]
			i += 3
		case 2:
			local[dest] = local[one] * local[two]
			i += 3
		case 99:
			return local
		}
	}
	return local
}
