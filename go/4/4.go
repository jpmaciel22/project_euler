package main

import (
	"fmt"
	"project_euler/libs"
)

func main() {
	max := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			n, status := libs.CheckPalindrome(j * i)
			if status && n > max {
				max = n
			}
		}
	}
	fmt.Println(max)
}
