package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/mbarrin/advent_of_code/util"
)

type game struct {
	ax       int
	ay       int
	bx       int
	by       int
	aEnd     int
	bEnd     int
	aPresses int
	bPresses int
	winning  bool
}

var pattern = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

func main() {
	defer util.TimeTaken(time.Now())

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	games := []game{}
	matches := pattern.FindAllSubmatch(input, -1)
	for _, match := range matches {
		ax, _ := strconv.Atoi(string(match[1]))
		ay, _ := strconv.Atoi(string(match[2]))
		bx, _ := strconv.Atoi(string(match[3]))
		by, _ := strconv.Atoi(string(match[4]))
		aEnd, _ := strconv.Atoi(string(match[5]))
		bEnd, _ := strconv.Atoi(string(match[6]))

		g := game{
			ax: ax, ay: ay,
			bx: bx, by: by,
			aEnd: aEnd, bEnd: bEnd,
		}

		a := g.ax - g.ay
		b := g.bx - g.by
		end := aEnd - bEnd

		if a == 0 {
			//fmt.Println(g)
			//fmt.Println(b)
			//fmt.Println(bEnd)
			//fmt.Println(end)
			//fmt.Println(by)

			fmt.Println("foo:", end/b)
			fmt.Println("foo:", b*end/b)

		}

		aPress, bPress := 1, 0

		for {
			bPress++
			if a == 0 {
				break
			}

			if (end-(b*bPress))%a == 0 {
				aPress = (end - (b * bPress)) / a
			} else {
				continue
			}

			if aPress > 100 || bPress > 100 {
				g.winning = false
				break
			}
			if aEnd == (aPress*g.ax)+(bPress*g.bx) && bEnd == (aPress*g.ay)+(bPress*g.by) {
				g.aPresses = aPress
				g.bPresses = bPress
				g.winning = true
				break
			}
		}

		games = append(games, g)

	}
	//fmt.Println(games)
	total := 0
	for _, game := range games {
		if game.winning {
			total += (game.aPresses * 3) + (game.bPresses * 1)
		}
	}

	// too low 33678
	fmt.Println("part 1:", total)
}
