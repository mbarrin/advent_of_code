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

func ManhattenDistance(x1, x2, y1, y2 int) int {
	return Abs(x1-y1) + Abs(x2-y2)
}
