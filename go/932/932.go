package main

import (
	"fmt"
	"strconv"
)

func summedSquaredEqualsNumber(a, b int) (int, bool) {
	n := (a + b) * (a + b)
	concatenate := strconv.Itoa(a) + strconv.Itoa(b)
	if strconv.Itoa(int(n)) == concatenate {
		return int(n), true
	}
	return int(n), false
}

func main() {

	sumAll2025Numbers := 0
	for i := 1; i < 10000000000000000; i++ {
		for j := 1; j < 10000000000000000; j++ {
			n, result := summedSquaredEqualsNumber(i, j)
			if result {
				sumAll2025Numbers += n
			}
		}
	}
	fmt.Println(sumAll2025Numbers)
}
