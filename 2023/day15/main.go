package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	line := []byte{}
	for s.Scan() {
		line = s.Bytes()
	}

	total := 0
	stream := []byte{}

	for _, x := range line {
		if x == ',' {
			total += hash(stream)
			stream = []byte{}
			continue
		}
		stream = append(stream, x)
	}
	total += hash(stream)

	fmt.Println("part 1:", total)

	instructions := strings.Split(string(line), ",")

	for _, x := range instructions {
		//var label string
		var boxID int
		var focalLength int
		y := []byte(x)
		if ok, _ := regexp.Match(`\w+=\d+`, y); ok {
			a := strings.Split(x, "=")
			boxID = hash([]byte(a[0]))
			focalLength, _ = strconv.Atoi(a[1])
			fmt.Println(boxID)
			fmt.Println(focalLength)
		} else {
			a := strings.Split(x, "-")
			fmt.Println(a)
		}
	}

	fmt.Println("part 2:", total)
}

func hash(b []byte) (h int) {
	for _, x := range b {
		h += int(x)
		h *= 17
		h %= 256
	}
	return
}
