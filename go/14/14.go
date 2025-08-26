package main

import (
	"fmt"
	"project_euler/libs"
)

func main() {
	max := 0
	maxNumber := 0
	for i := 1; i < 1000000; i++ {
		till1 := libs.CollatzLength(i)
		if till1 > max {
			max = till1
			maxNumber = i
		}
	}
	fmt.Println(max, maxNumber)
}
