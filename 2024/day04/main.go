package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(input)
}
