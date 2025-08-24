package main

import (
	"fmt"
	"strconv"
)

func checkPalindrome(n int) (int, bool) {
	str := strconv.Itoa(n)
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	if reverse == str {
		return n, true
	}
	return n, false
}

func main() {
	max := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			n, status := checkPalindrome(j * i)
			if status && n > max {
				max = n
			}
		}
	}
	fmt.Println(max)
}
