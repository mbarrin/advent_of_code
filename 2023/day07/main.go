package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards []int
	bid   int
	rank  int
}

type hands []hand

var mapping = map[rune]int{
	'1': 1, '2': 2, '3': 3,
	'4': 4, '5': 5, '6': 6,
	'7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 11, 'Q': 12,
	'K': 13, 'A': 14,
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(1)
	}
	s := bufio.NewScanner(f)

	game := map[int]hands{}

	for s.Scan() {
		x := s.Text()

		split := strings.Split(x, " ")
		bid, _ := strconv.Atoi(split[1])

		initialRank := rank(cardToNum(split[0]))
		hand := hand{cards: cardToNum(split[0]), bid: bid, rank: initialRank}

		game[initialRank] = append(game[initialRank], hand)
	}

	orderedKeys := []int{}
	for k := range game {
		orderedKeys = append(orderedKeys, k)
	}
	slices.Sort(orderedKeys)

	ranks := hands{}

	for _, rank := range orderedKeys {
		sort.Sort(game[rank])
		for _, v := range game[rank] {
			ranks = append(ranks, v)
		}
	}

	total := 0
	for i := range ranks {
		fmt.Println(ranks[i])
		total += ranks[i].bid * (i + 1)
	}

	fmt.Println("part 1:", total)

}

func cardToNum(s string) (i []int) {
	for _, x := range s {
		i = append(i, mapping[x])
	}
	return
}

func rank(b []int) int {
	m := make(map[int]int)

	for _, x := range b {
		m[x]++
	}

	orderedKeys := []int{}
	for k := range m {
		orderedKeys = append(orderedKeys, k)
	}
	slices.Sort(orderedKeys)

	maxKey, maxVal := 0, 0
	for _, x := range orderedKeys {
		if m[x] >= maxVal {
			maxKey = x
			maxVal = m[x]
		}
	}

	//fmt.Println("key:", maxKey)
	//fmt.Println("val:", maxVal)

	m[maxKey] += m[11]

	delete(m, 11)

	var rank int

	switch len(m) {
	case 0:
		rank = 7
	case 1:
		rank = 7
	case 2:
		rank = 5
		for _, v := range m {
			if v == 4 {
				rank = 6
			}
		}
	case 3:
		rank = 3
		for _, v := range m {
			if v == 3 {
				rank = 4
			}
		}
	case 4:
		rank = 2
	default:
		rank = 1
	}

	return rank
}

func (h hands) Len() int {
	return len(h)
}

func (h hands) Less(i, j int) bool {
	for x := range h[i].cards {
		foo, bar := h[i].cards[x], h[j].cards[x]

		if foo == 11 {
			foo = 1
		}

		if bar == 11 {
			bar = 1
		}

		if foo < bar {
			return true
		}

		if foo > bar {
			return false
		}
	}
	return false
}

func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
