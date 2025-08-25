package main

import (
	"fmt"
	"project_euler/libs"
)

func main() {
	var sum int64 = 0

	for i := range 2000000 {
		num, status := libs.CrivoEratostenes(int64(i))
		if status {
			sum += num
		}
	}
	fmt.Println(sum)
}
