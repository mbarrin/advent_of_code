package util

import (
	"fmt"
	"time"
)

func Abs(i int) int {
	if i < 0 {
		return 0 - i
	}
	return i
}

func TimeTaken(s time.Time) {
	fmt.Println("Took: ", time.Since(s))
}
